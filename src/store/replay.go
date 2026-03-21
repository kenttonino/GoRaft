package store

import (
	"GoRaft/src/wal"
	"fmt"
)

// replay reads the WAL file and re-applies every command to
// the in-memory map, restoring state from before the crash.
func (s *Store) replay(walPath string) error {
	entries, err := wal.Replay(walPath)
	if err != nil {
		return err
	}

	// Re-apply each command directly to the map (no WAL write this time
	// the entries are already on disk, we're just rebuilding memory).
	for _, parts := range entries {
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "SET":
			if len(parts) == 3 {
				s.data[parts[1]] = parts[2]
			}
		case "DEL":
			if len(parts) == 2 {
				delete(s.data, parts[1])
			}
		}
	}

	fmt.Printf("WAL replayed: %d entries restored\n", len(entries))
	return nil
}
