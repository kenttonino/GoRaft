package store

// Set stores a value under the given key.
// If the key already exists, it gets overwritten.
// Example: Set("name", "goraft")
func (s *Store) Set(key, value string) {
	// Lock for writing - no other goroutine can read or write, until
	// we call Unlock via defer.
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

// Get retrieves the value for a given key.
// Returns the value and a boolean, true if found, false if not.
// Example: Get("name") -> "goraft", true
// Example: Get("name-2") -> "", false
func (s *Store) Get(key string) (string, bool) {
	// RLock for reading, multiple goroutines can read at the same time,
	// but writing is blocked until we RUnlock.
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	return val, ok
}

// Delete removes a key-value pair from the store.
// If the key doesn't exist, nothing happens - no error.
// Example: Delete("name")
func (s *Store) Delete(key string) {
	// Lock for writing.
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}
