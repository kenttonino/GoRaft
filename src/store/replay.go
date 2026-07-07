package store

import (
	"fmt"
)

func (s *Store) replay() error {
	entries, err := s.wal.Replay()
	if err != nil {
		return err
	}

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
