package wal

import "fmt"

// Write appends a command to the WAL file on disk.
// Each command is written as one line.
// Example: Write("SET", "name", "goraft") -> "SET name goraft\n"
// IMPORTANT: This must be called BEFORE applying the command to
// the in-memory store. That's the whole point of write-ahead.
func (w *WAL) Write(cmd, key, value string) error {
	// Format the command as a single line.
	// For DEL commands, value will be empty.
	var line string
	if value != "" {
		line = fmt.Sprintf("%s %s %s\n", cmd, key, value)
	} else {
		line = fmt.Sprintf("%s %s\n".cmd, key)
	}

	// Write the line to disk.
	_, err := fmt.Fprint(w.file, line)
	if err != nil {
		return fmt.Errorf("failed to write to WAL: %w", err)
	}

	// Sync forces the OS to flush its internal buffer to disk
	// immediately. Without this, the OS might hold the data in
	// memory briefly, and a crash at the exact moment would still
	// lose the entry.
	return w.file.Sync()
}
