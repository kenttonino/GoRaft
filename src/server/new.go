package server

import "GoRaft/src/store"

func New(addr string, store *store.Store) *Server {
	newServer := Server{addr: addr, store: store}
	return &newServer
}
