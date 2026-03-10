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
	// brute-force two-pointer search per bank
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

// PrefixSum_Lobby1 is an O(n) alternative to Lobby1. It precalculates,
// for each index i, the largest digit appearing *after* i, then uses that to
// compute the best two‑digit value for each possible first battery.  By
// restricting the search to i < len(bank)-1 we avoid invalid pairs (there is
// no second battery after the last position).
func PrefixSum_Lobby1(banks []string) int {
	total := 0
	for _, bank := range banks {
		n := len(bank)
		if n < 2 {
			// no valid pair
			continue
		}

		// slice of maximum digit to the right of each index
		maxAfter := make([]int, n)
		mva := 0
		// scan backward; maxAfter[i] should be max digit in bank[i+1:]
		for i := n - 1; i >= 0; i-- {
			maxAfter[i] = mva
			d := int(bank[i] - '0')
			if d > mva {
				mva = d
			}
		}

		maxVal := 0
		for i := 0; i < n-1; i++ {
			d := int(bank[i] - '0')
			val := d*10 + maxAfter[i]
			if val > maxVal {
				maxVal = val
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
