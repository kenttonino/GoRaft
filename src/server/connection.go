package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// handleConn reads and handles commands from a single client.
func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	// Read commands line by line from the client.
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		// Uppercase the command so SET, set, Set all work.
		cmd := strings.ToUpper(parts[0])

		switch cmd {

		// SET key value — write to WAL then store in memory
		case "SET":
			if len(parts) != 3 {
				fmt.Fprintln(conn, "ERR usage: SET key value")
				continue
			}
			// Set now returns an error — check it.
			if err := s.store.Set(parts[1], parts[2]); err != nil {
				fmt.Fprintln(conn, "ERR", err)
				continue
			}
			fmt.Fprintln(conn, "OK")

		// GET key — read from memory only, no WAL needed
		case "GET":
			if len(parts) != 2 {
				fmt.Fprintln(conn, "ERR usage: GET key")
				continue
			}
			val, ok := s.store.Get(parts[1])
			if !ok {
				fmt.Fprintln(conn, "NULL")
			} else {
				fmt.Fprintln(conn, val)
			}

		// DEL key — write to WAL then delete from memory
		case "DEL":
			if len(parts) != 2 {
				fmt.Fprintln(conn, "ERR usage: DEL key")
				continue
			}
			// Delete now returns an error — check it.
			if err := s.store.Delete(parts[1]); err != nil {
				fmt.Fprintln(conn, "ERR", err)
				continue
			}
			fmt.Fprintln(conn, "OK")

		default:
			fmt.Fprintln(conn, "ERR unknown command:", cmd)
		}
	}
}
