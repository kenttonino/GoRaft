package store

import "fmt"

// Set writes the command to the WAL first, then stores it in memory.
// If the server crashes after the WAL write but before the memory
// update, the WAL replay on restart will recover it.
func (s *Store) Set(key, value string) error {
	// Write to WAL first, this is the "write-ahead" part.
	if err := s.wal.Write("SET", key, value); err != nil {
		return fmt.Errorf("WAL write failed: %w", err)
	}

	// Now safe to update memory.
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

// Get retrieves a value from memory, no WAL needed for reads.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// Delete writes the command to the WAL first, then removes from memory.
func (s *Store) Delete(key string) error {
	// Write to WAL first.
	if err := s.wal.Write("DEL", key, ""); err != nil {
		return fmt.Errorf("WAL write failed: %w", err)
	}

	// Now safe to delete from memory.
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}
