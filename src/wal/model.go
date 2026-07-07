package wal

import (
	"os"
	"sync"
)

// WAL (Write-Ahead Log) is a file on disk that records every
// command before it is applied to the in-memory store. If the
// server crashes, we can replay this file on startup to restore
// all the data that was previously stored.
type WAL struct {
	mu   sync.Mutex
	file *os.File
}
