package store

import (
	"GoRaft/src/wal"
	"fmt"
)

// New creates a new Store and replays the WAL to restore previous state.
// walPath is the path to the WAL file "data/wal.log".
func New(walPath string) (*Store, error) {
	// Open the WAL file (creates it if it doesn't exist).
	w, err := wal.New(walPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open WAL: %w", err)
	}

	s := &Store{
		data: make(map[string]string),
		wal:  w,
	}

	// Replay the WAL to restore data from the last one.
	// This is what makes data survive crashes.
	if err := s.replay(walPath); err != nil {
		return nil, fmt.Errorf("failed to replay WAL: %w", err)
	}

	return s, nil
}
