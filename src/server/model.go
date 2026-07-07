package server

import (
	"GoRaft/src/store"
	"net"
)

// Server listens for income TCP connections and handles
// commands from clients (SET, GET, DEL).
type Server struct {
	addr  string
	store *store.Store
	ln    net.Listener
}
