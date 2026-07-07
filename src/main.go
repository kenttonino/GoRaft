package main

import (
	"GoRaft/src/server"
	"GoRaft/src/store"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, err := store.New("data/wal.log")
	if err != nil {
		log.Fatal("failed to start store:", err)
	}

	srv := server.New(":7001", s)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("shutting down...")
		srv.Stop()
		s.Close()
		os.Exit(0)
	}()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
