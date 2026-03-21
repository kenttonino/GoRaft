package wal

// Closes the WAL file when the server shuts down.
func (w *WAL) Close() error {
	return w.file.Close()
}
