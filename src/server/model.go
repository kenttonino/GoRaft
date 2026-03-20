package server

import "GoRaft/src/store"

// Server listens for income TCP connections and handles
// commands from clients (SET, GET, DEL).
type Server struct {
	// addr is the address we listen on (e.g. :7001).
	addr string
	// store is our KV database.
	// Shared across all connections.
	store *store.Store
}
