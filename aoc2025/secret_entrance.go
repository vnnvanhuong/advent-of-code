package aoc2025

import (
	"math"
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
// including mid-rotation crossings.
// Credits: https://z-nerd.com/posts/2025/12/aoc-2025-day-one/
func SecretEntrance2(rotations []string) int {
	dialPosition := 50
	count := 0
	startAtZero := false

	for _, rotation := range rotations {
		direction := rotation[0]
		distance, _ := strconv.Atoi(rotation[1:])
		if direction == 'L' {
			distance = -distance
		}

		// move the dial
		dialPosition += distance

		// count the number of times the dial crosses over 0
		count += int(math.Floor(math.Abs(float64(dialPosition)) / 100))

		// if dial position lands on 0, count++
		if dialPosition == 0 {
			count++
		}

		// edge case: add 1 when cross over 0 the negative direction
		if dialPosition < 0 && !startAtZero {
			count++
		}

		// reset the dial position within bound of 0-99
		dialPosition = (dialPosition%100 + 100) % 100
		startAtZero = dialPosition == 0
	}

	return count
}
