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
	s.ln = ln

	fmt.Println("GoRaft listening on", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			if s.ln == nil {
				return nil
			}
			fmt.Println("connection error:", err)
			continue
		}
		go s.handleConn(conn)
	}
}

// Stop closes the listener, causing Start's accept loop to return.
func (s *Server) Stop() {
	if s.ln != nil {
		s.ln.Close()
		s.ln = nil
	}
}
