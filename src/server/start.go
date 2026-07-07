package server

import (
	"fmt"
	"net"
)

func (s *Server) Start() error {
	return s.listen(s.addr)
}

func (s *Server) StartOnRandom(ready chan<- string) (string, error) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	s.mu.Lock()
	s.ln = ln
	s.mu.Unlock()
	addr := ln.Addr().String()
	ready <- addr
	return addr, s.accept(ln)
}

func (s *Server) listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.ln = ln
	s.mu.Unlock()
	fmt.Println("GoRaft listening on", ln.Addr())
	return s.accept(ln)
}

func (s *Server) accept(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			s.mu.Lock()
			stopped := s.ln == nil
			s.mu.Unlock()
			if stopped {
				return nil
			}
			fmt.Println("connection error:", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ln != nil {
		s.ln.Close()
		s.ln = nil
	}
}
