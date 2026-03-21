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

	// Read the file.
	var entries [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines.
		if line == "" {
			continue
		}

		// Split the line back into parts.
		// "SET name goraft" -> ["SET", "name", "goraft"]
		parts := strings.Fields(line)
		entries = append(entries, parts)
	}

	return entries, scanner.Err()
}
