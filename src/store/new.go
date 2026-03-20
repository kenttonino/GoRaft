package store

// New creates and returns a fresh, empty Store.
// Always use this to create a Store - never create one directly.
func New() *Store {
	storeData := Store{
		data: make(map[string]string),
	}

	return &storeData
}
