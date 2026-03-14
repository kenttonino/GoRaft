package server

import (
	"GoRaft/src/store"
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Server listens for income TCP connections and handles
// commands from clients (SET, GET, DEL).
type Server struct {
	// addr is the address we listen on (e.g. :7001).
	addr string
	// store is our KV database.
	// Shared across all connections.
	store *store.Store
}

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

// New creates a new Server with the given address and store.
// Example: New(":7001", myStore)
func New(addr string, store *store.Store) *Server {
	newServer := Server{addr: addr, store: store}
	return &newServer
}

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
