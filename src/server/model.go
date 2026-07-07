package server

import (
	"GoRaft/src/store"
	"net"
	"sync"
)

type Server struct {
	addr  string
	store *store.Store
	mu    sync.Mutex
	ln    net.Listener
}
