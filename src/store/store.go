package store

import "sync"

// Store is our in-memory key-value database.
// Think of it like a dictionary - you store a value under a key,
// and you can look it up or delete it later.
type Store struct {
	// mu is a lock that prevents two goroutines from reading/writing,
	// the data at the same time, which would cause bugs (called "race condition").
	mu sync.RWMutex

	// data is the actual map that holds our key-value pairs.
	// Example: data["name"] = "goraft".
	data map[string]string
}

// New creates and returns a fresh, empty Store.
// Always use this to create a Store - never create one directly.
func New() *Store {
	storeData := Store{
		data: make(map[string]string),
	}

	return &storeData
}

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
