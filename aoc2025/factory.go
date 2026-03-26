package aoc2025

import (
	"strconv"
	"strings"
)

const maxCounters = 16

type counterVec [maxCounters]int

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

func parseJoltage(line string) ([][]int, []int) {
	parts := strings.Split(line, "] ")
	remaining := parts[1]

	sections := strings.SplitN(remaining, " {", 2)
	buttonSection := sections[0]
	targetSection := strings.TrimSuffix(sections[1], "}")

	buttonStrs := strings.Split(buttonSection, ") (")
	var buttons [][]int
	for _, bStr := range buttonStrs {
		bStr = strings.TrimPrefix(bStr, "(")
		bStr = strings.TrimSuffix(bStr, ")")
		if bStr == "" {
			continue
		}
		nums := strings.Split(bStr, ",")
		var indices []int
		for _, numStr := range nums {
			idx, _ := strconv.Atoi(numStr)
			indices = append(indices, idx)
		}
		buttons = append(buttons, indices)
	}

	targetStrs := strings.Split(targetSection, ",")
	var targets []int
	for _, tStr := range targetStrs {
		t, _ := strconv.Atoi(tStr)
		targets = append(targets, t)
	}

	return buttons, targets
}

// solveMachineJoltage uses binary decomposition to find the minimum total
// button presses. Each button's press count is decomposed into binary bits.
// At each bit level, we choose a subset of buttons to "activate" (0 or 1),
// matching the current parity of the remaining goal. The halved remainder
// is solved recursively at double cost per press.
func solveMachineJoltage(buttons [][]int, targets []int) int {
	numCounters := len(targets)
	numButtons := len(buttons)

	coeffs := make([]counterVec, numButtons)
	for j, btn := range buttons {
		for _, idx := range btn {
			coeffs[j][idx] = 1
		}
	}

	patternsByParity := make(map[counterVec]map[counterVec]int)
	for subset := 0; subset < (1 << numButtons); subset++ {
		var pattern counterVec
		cost := 0
		for j := 0; j < numButtons; j++ {
			if subset&(1<<j) != 0 {
				cost++
				for k := 0; k < numCounters; k++ {
					pattern[k] += coeffs[j][k]
				}
			}
		}

		var parity counterVec
		for k := 0; k < numCounters; k++ {
			parity[k] = pattern[k] % 2
		}

		if patternsByParity[parity] == nil {
			patternsByParity[parity] = make(map[counterVec]int)
		}
		if existing, ok := patternsByParity[parity][pattern]; !ok || cost < existing {
			patternsByParity[parity][pattern] = cost
		}
	}

	memo := make(map[counterVec]int)
	var solve func(goal counterVec) int
	solve = func(goal counterVec) int {
		allZero := true
		for k := 0; k < numCounters; k++ {
			if goal[k] != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			return 0
		}

		if cached, ok := memo[goal]; ok {
			return cached
		}

		var parity counterVec
		for k := 0; k < numCounters; k++ {
			parity[k] = goal[k] % 2
		}

		answer := 1_000_000
		if patterns, ok := patternsByParity[parity]; ok {
			for pattern, cost := range patterns {
				valid := true
				var newGoal counterVec
				for k := 0; k < numCounters; k++ {
					if pattern[k] > goal[k] {
						valid = false
						break
					}
					newGoal[k] = (goal[k] - pattern[k]) / 2
				}
				if !valid {
					continue
				}

				total := cost + 2*solve(newGoal)
				if total < answer {
					answer = total
				}
			}
		}

		memo[goal] = answer
		return answer
	}

	var goal counterVec
	copy(goal[:], targets)
	return solve(goal)
}

func Factory2(input []string) int {
	total := 0
	for _, line := range input {
		buttons, targets := parseJoltage(line)
		total += solveMachineJoltage(buttons, targets)
	}
	return total
}
