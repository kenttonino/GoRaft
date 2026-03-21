package main

import (
	"GoRaft/src/server"
	"GoRaft/src/store"
	"log"
)

func main() {
	// Create the store with a WAL file at "data/wal.log".
	// On first run this creates a fresh file.
	// On restart it replays the file and restores all previous data.
	s, err := store.New("data/wal.log")
	if err != nil {
		log.Fatal("failed to start store:", err)
	}
	// Always close the WAL cleanly when the server stops.
	defer s.Close()

	// Create and start the TCP server on port 7001.
	srv := server.New(":7001", s)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
