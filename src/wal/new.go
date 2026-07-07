package wal

import (
	"fmt"
	"os"
)

func New(path string) (*WAL, error) {
	err := os.MkdirAll(Filepath(path), 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create WAL directory: %w", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open WAL file: %w", err)
	}

	return &WAL{file: file}, nil
}
