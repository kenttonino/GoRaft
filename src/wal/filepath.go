package wal

// filepath extracts the directory part of a file path.
// Example: "data/wal.log" → "data"
// This is used to create the parent directory before
// opening the file.
func Filepath(path string) string {
	// Find the last slash in the path.
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[:i]
		}
	}
	// No slash found, file is in the current directory, no folder needed.
	return "."
}
