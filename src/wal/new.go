package wal

import (
	"fmt"
	"os"
)

// New opens (or creates) the WAL file at the given path. If the
// file already exists, new entries are appended to the end.
// Example: New("data/wal.log")
func New(path string) (*WAL, error) {
	// os.MkdirAll creates the folder if it doesn't exist yet.
	// Example: "data/wal.log" -> creates the "data/" folder first.
	// 07 -> Octal.
	// 5 -> Owner.
	// 5 -> Group.
	err := os.MkdirAll(Filepath(path), 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create WAL directory: %w", err)
	}

	// Open the file for appending and writing.
	// os.O_CREATE -> create the file if it doesn't exist.
	// os.O_APPEND -> always write at the end of the file.
	// os.O_WRONLY -> we only write to this file handle.
	// 0644 -> File permission (owner can read/write others can read).
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open WAL file: %w", err)
	}

	return &WAL{file: file}, nil
}
