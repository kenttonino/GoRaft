package wal

import "fmt"

func (w *WAL) Write(cmd, key, value string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	var line string
	if value != "" {
		line = fmt.Sprintf("%s %s %s\n", cmd, key, value)
	} else {
		line = fmt.Sprintf("%s %s\n", cmd, key)
	}

	_, err := fmt.Fprint(w.file, line)
	if err != nil {
		return fmt.Errorf("failed to write to WAL: %w", err)
	}

	return w.file.Sync()
}
