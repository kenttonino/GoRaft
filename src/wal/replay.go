package wal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Replay(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to open WAL for replay: %w", err)
	}
	defer file.Close()

	return replayReader(file)
}

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
