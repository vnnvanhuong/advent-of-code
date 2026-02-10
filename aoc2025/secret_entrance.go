package aoc2025

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// SecretEntrance1 simulates a circular dial with 100 positions (0-99).
// Starting at position 50, each rotation moves the pointer left ('L') or
// right ('R') by the given distance, wrapping around the dial modularly.
// The function counts how many times the pointer lands exactly on position 0
// (the secret entrance) after processing each rotation.
//
// Example walkthrough:
// L68 -> P = ((50 - 68) % 100 + 100) % 100 = 82
// L30 -> P = ((82 - 30) % 100 + 100) % 100 = 52
// R48 -> P = ((52 + 48) % 100 + 100) % 100 = 0
func SecretEntrance1(rotations []string) int {
	p := 50
	count := 0
	for _, rotation := range rotations {
		direction := rotation[0]
		distance, _ := strconv.Atoi(rotation[1:])
		if direction == 'L' {
			distance = -distance
		}
		p = ((p+distance)%100 + 100) % 100
		if p == 0 {
			count++
		}
	}

	return count
}

// SecretEntrance2 counts every click that causes the dial to point at 0,
// including mid-rotation crossings. Uses floor division to count how many
// multiples of 100 are crossed during each rotation on the unwrapped number line.
func SecretEntrance2(rotations []string) int {
	count := 0
	return count
}

func SecretEntranceInput(filename string) []string {
	// read lines from file and append to a slice
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
