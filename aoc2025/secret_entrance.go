package aoc2025

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// rotations = ["L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"]
// P = 50
// count = 0

// logical solution:
// for each rotation, determine the direction and the distance
// if the direction is L (left), minus the distance from P. Else, if the direction is R (right), add the distance to P.
// normalize p to be between 0 and 100.
// increment count by 1 if P is 0 or 100
func SecretEntrance1(rotations []string) int {
	p := 50
	count := 0
	for _, rotation := range rotations {
		direction := rotation[0]
		distance, _ := strconv.Atoi(rotation[1:])
		if direction == 'L' {
			p = (p - distance) % 100
		} else {
			p = (p + distance) % 100
		}

		if p == 0 || p == 100 {
			count++
		}
	}

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
