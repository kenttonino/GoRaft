package server

import (
	"GoRaft/src/store"
	"net"
)

type Server struct {
	addr  string
	store *store.Store
	ln    net.Listener
}
