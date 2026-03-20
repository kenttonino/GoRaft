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
