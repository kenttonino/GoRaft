package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// handleConn reads commands from a single client connection and sends back responses.
// It runs in its own goroutine.
func (s *Server) handleConn(conn net.Conn) {
	// Always close the connection when this function exits.
	defer conn.Close()

	// bufio.Scanner reads the connection line by line.
	// Each line from the client is one command (e.g. "SET name goraft").
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		// Read one line and clean up extra whitespace.
		line := strings.TrimSpace(scanner.Text())

		// Split the line into parts by spaces.
		// "SET name goraft" -> ["SET", "name", "goraft"]
		parts := strings.Fields(line)

		// Ignore empty lines.
		if len(parts) == 0 {
			continue
		}

		// The first word is always the command.
		// We uppercase it so "set", "SET", "Set" all work.
		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "SET":
			// SET requires exactly 3 parts: SET, key, value
			if len(parts) != 3 {
				fmt.Fprintln(conn, "ERR usage: SET key value")
				continue
			}
			s.store.Set(parts[1], parts[2])
			fmt.Fprintln(conn, "OK")

		case "GET":
			// GET requires exactly 2 parts: GET, key
			if len(parts) != 2 {
				fmt.Fprintln(conn, "Err usage: GET key")
				continue
			}
			val, ok := s.store.Get(parts[1])
			if !ok {
				// Key doesn't exist - return NULL.
				fmt.Fprintln(conn, "NULL")
			} else {
				fmt.Fprintln(conn, val)
			}

		case "DEL":
			// DEL requires exactly 2 parts: DEL, key
			if len(parts) != 2 {
				fmt.Fprintln(conn, "ERR usage: DEL key")
				continue
			}

			s.store.Delete(parts[1])
			fmt.Fprintln(conn, "OK")

		default:
			fmt.Fprintln(conn, "ERR unknown command:", cmd)
		}
	}
}
