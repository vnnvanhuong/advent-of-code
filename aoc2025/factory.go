package aoc2025

import (
	"strconv"
	"strings"
)

func parseMachine(line string) (int, []int) {
	// Example line: [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
	parts := strings.Split(line, "] ")
	targetStr := strings.TrimPrefix(parts[0], "[")

	targetMask := 0
	for i, char := range targetStr {
		if char == '#' {
			targetMask |= (1 << i)
		}
	}

	remaining := parts[1]
	buttonStrs := strings.Split(strings.Split(remaining, " {")[0], ") (")

	buttons := make([]int, 0)
	for _, bStr := range buttonStrs {
		bStr = strings.TrimPrefix(bStr, "(")
		bStr = strings.TrimSuffix(bStr, ")")

		if bStr == "" {
			continue
		}

		nums := strings.Split(bStr, ",")
		mask := 0
		for _, numStr := range nums {
			bit, _ := strconv.Atoi(numStr)
			mask |= (1 << bit)
		}
		buttons = append(buttons, mask)
	}

	return targetMask, buttons
}

func minPressesForMachine(target int, buttons []int) int {
	n := len(buttons)
	minPresses := n + 1 // Initialize with a value larger than max possible

	// Iterate through all 2^n subsets
	for subset := 0; subset < (1 << n); subset++ {
		currentMask := 0
		presses := 0

		for i := 0; i < n; i++ {
			if (subset & (1 << i)) != 0 {
				currentMask ^= buttons[i]
				presses++
			}
		}

		if currentMask == target {
			if presses < minPresses {
				minPresses = presses
			}
		}
	}

	return minPresses
}

func Factory(input []string) int {
	totalPresses := 0
	for _, line := range input {
		target, buttons := parseMachine(line)
		totalPresses += minPressesForMachine(target, buttons)
	}
	return totalPresses
}
