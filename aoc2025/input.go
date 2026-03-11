package aoc2025

import (
	"bufio"
	"log"
	"os"
)

// ReadLines opens the named file and returns its contents as a slice of
// strings, one for each line. It is the common helper that replaces the
// various per-puzzle input functions previously scattered across the package.
// Any I/O errors are fatal, which matches the behaviour of the earlier
// helpers seen in other files.
func ReadLines(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	return lines
}
