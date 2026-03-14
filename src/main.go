package main

import (
	"GoRaft/src/server"
	"GoRaft/src/store"
	"log"
)

func main() {
	// Create a fresh in-memory KV store.
	// This is where all our data lives while the server is running.
	s := store.New()

	// Create a TCP server that listens on port 7001,
	// and uses our KV store to handle commands.
	sServer := server.New(":7001", s)

	// Start the server.
	// This blocks forever, the server keeps running until you stop it.
	// If something goes wrong (e.g. port already in use), log.Fatal prints
	// the error and exist the program.
	if err := sServer.Start(); err != nil {
		log.Fatal(err)
	}
}
