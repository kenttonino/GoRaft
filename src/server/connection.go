package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "SET":
			if len(parts) != 3 {
				fmt.Fprintln(conn, "ERR usage: SET key value")
				continue
			}
			if err := s.store.Set(parts[1], parts[2]); err != nil {
				fmt.Fprintln(conn, "ERR", err)
				continue
			}
			fmt.Fprintln(conn, "OK")
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
		case "DEL":
			if len(parts) != 2 {
				fmt.Fprintln(conn, "ERR usage: DEL key")
				continue
			}
			if err := s.store.Delete(parts[1]); err != nil {
				fmt.Fprintln(conn, "ERR", err)
				continue
			}
			fmt.Fprintln(conn, "OK")
		default:
			fmt.Fprintln(conn, "ERR unknown command:", cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(conn, "ERR", err)
	}
}
