package store

// Close cleanly shuts down the WAL file.
// Always call this when the server stops.
func (s *Store) Close() error {
	return s.wal.Close()
}
