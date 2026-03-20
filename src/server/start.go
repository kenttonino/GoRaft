package server

import (
	"fmt"
	"net"
)

// Starts opens a TCP socket and begins accepting client connections.
// Each connection gets its own goroutine so multiple clients can
// connect at the same time without blocking each other.
func (s *Server) Start() error {
	// net.Listen opens a TCP socket on the given address.
	// Think of it like opening a door - clients can now knock.
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	// Close the door when the server stops.
	defer ln.Close()

	fmt.Println("GoRaft listening on", s.addr)

	// Keep waiting for new client connections forever.
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go s.handleConn(conn)
	}
}
