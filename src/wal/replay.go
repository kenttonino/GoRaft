package wal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Replay reads the WAL file from the beginning and returns all
// the commands in order, so they can be replayed on startup. Each
// entry is returned as a slice of strings, same as parts from the
// TCP server: ["SET", "name", "goraft"] or ["DEL", "name"]
func Replay(path string) ([][]string, error) {
	// Open the file for reading only.
	file, err := os.Open(path)
	if err != nil {
		// If the file doesn't exist yet, there's nothing to replay.
		// This s normal on first startup, return empty, no error.
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to open WAL for replay: %w", err)
	}
	defer file.Close()

	return replayReader(file)
}

// Replay reads all entries from the existing file handle, seeking
// to the beginning first. Use this instead of the package-level
// Replay to avoid opening a second file descriptor.
func (w *WAL) Replay() ([][]string, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, err := w.file.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("failed to seek WAL to start: %w", err)
	}

	return replayReader(w.file)
}

func replayReader(f *os.File) ([][]string, error) {
	var entries [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		entries = append(entries, parts)
	}
	return entries, scanner.Err()
}
