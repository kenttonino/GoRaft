package server

import "GoRaft/src/store"

// New creates a new Server with the given address and store.
// Example: New(":7001", myStore)
func New(addr string, store *store.Store) *Server {
	newServer := Server{addr: addr, store: store}
	return &newServer
}
