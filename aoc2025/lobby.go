package aoc2025

import (
	"bufio"
	"log"
	"os"
)

// Lobby1 computes the total output joltage for part one of the lobby puzzle.
// Each string in banks represents a line of digit-labeled batteries. Exactly two
// batteries must be turned on per bank and the output is the two-digit number
// formed by their labels in order. The function returns the sum of each bank's
// maximum possible joltage.
func Lobby1(banks []string) int {
	// brute-force two-pointer search per bank, as suggested by the user.
	// for each line of digits we scan all ordered pairs and keep the largest
	// two-digit value encountered.
	total := 0
	for _, bank := range banks {
		maxVal := 0
		// nested loops over indices
		for i := 0; i < len(bank)-1; i++ {
			d1 := int(bank[i] - '0')
			for j := i + 1; j < len(bank); j++ {
				d2 := int(bank[j] - '0')
				val := d1*10 + d2
				if val > maxVal {
					maxVal = val
				}
			}
		}
		total += maxVal
	}
	return total
}

// Lobby2 is reserved for part two of the puzzle. Behavior is currently undefined
// and will be implemented after tests are written.
func Lobby2(banks []string) int {
	// implementation TBD
	return 0
}

// LobbyInput reads the battery banks from a file; each line corresponds to one
// bank (a contiguous string of digit characters). It returns the slice of lines.
func LobbyInput(filename string) []string {
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
