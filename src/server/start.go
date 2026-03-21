package server

import (
	"fmt"
	"net"
)

// Start opens a TCP socket and begins accepting client connections.
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("GoRaft listening on", s.addr)

	for {
		// Wait for a client to connect.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("connection error:", err)
			continue
		}

		// Handle this client in its own goroutine.
		go s.handleConn(conn)
	}
}
