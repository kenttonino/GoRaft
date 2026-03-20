package wal

import "os"

// WAL (Write-Ahead Log) is a file on disk that records every
// command before it is applied to the in-memory store. If the
// server crashes, we can replay this file on startup to restore
// all the data that was previously stored.
type WAL struct {
	// file is the actual file on disk where commands are written.
	file *os.File
}
