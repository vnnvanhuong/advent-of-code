package aoc2025

import (
	"math"
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

func minJoltagePresses(buttons [][]int, targets []int) int {
	m := len(targets)
	n := len(buttons)

	A := make([][]float64, m)
	for i := range A {
		A[i] = make([]float64, n)
	}
	for j, btn := range buttons {
		for _, idx := range btn {
			A[idx][j] = 1
		}
	}

	b := make([]float64, m)
	for i := range targets {
		b[i] = float64(targets[i])
	}

	c := make([]float64, n)
	for j := range c {
		c[j] = 1
	}

	x, _ := solveLPSimplex(A, b, c)
	if x == nil {
		return -1
	}

	total := 0.0
	for _, v := range x {
		total += v
	}
	return int(math.Round(total))
}

// solveLPSimplex solves: minimize c^T x, subject to Ax = b, x >= 0
// using the Big-M method with the simplex algorithm.
func solveLPSimplex(A [][]float64, b []float64, c []float64) ([]float64, float64) {
	m := len(A)
	n := len(c)
	bigM := 1e8
	totalVars := n + m

	tab := make([][]float64, m+1)
	for i := range tab {
		tab[i] = make([]float64, totalVars+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tab[i+1][j] = A[i][j]
		}
		tab[i+1][n+i] = 1.0
		tab[i+1][totalVars] = b[i]
	}

	for j := 0; j < n; j++ {
		var colSum float64
		for i := 0; i < m; i++ {
			colSum += A[i][j]
		}
		tab[0][j] = c[j] - bigM*colSum
	}
	var bSum float64
	for i := 0; i < m; i++ {
		bSum += b[i]
	}
	tab[0][totalVars] = -bigM * bSum

	basis := make([]int, m)
	for i := 0; i < m; i++ {
		basis[i] = n + i
	}

	for iter := 0; iter < 10000; iter++ {
		pivotCol := -1
		minRC := -1e-10
		for j := 0; j < totalVars; j++ {
			if tab[0][j] < minRC {
				minRC = tab[0][j]
				pivotCol = j
			}
		}
		if pivotCol == -1 {
			break
		}

		pivotRow := -1
		minRatio := math.MaxFloat64
		for i := 1; i <= m; i++ {
			if tab[i][pivotCol] > 1e-10 {
				ratio := tab[i][totalVars] / tab[i][pivotCol]
				if ratio < minRatio-1e-10 {
					minRatio = ratio
					pivotRow = i
				}
			}
		}
		if pivotRow == -1 {
			break
		}

		pivot := tab[pivotRow][pivotCol]
		for j := 0; j <= totalVars; j++ {
			tab[pivotRow][j] /= pivot
		}
		for i := 0; i <= m; i++ {
			if i != pivotRow {
				factor := tab[i][pivotCol]
				if factor != 0 {
					for j := 0; j <= totalVars; j++ {
						tab[i][j] -= factor * tab[pivotRow][j]
					}
				}
			}
		}
		basis[pivotRow-1] = pivotCol
	}

	for i := 0; i < m; i++ {
		if basis[i] >= n && tab[i+1][totalVars] > 1e-6 {
			return nil, math.Inf(1)
		}
	}

	x := make([]float64, n)
	for i := 0; i < m; i++ {
		if basis[i] < n {
			x[basis[i]] = tab[i+1][totalVars]
		}
	}

	objVal := 0.0
	for j := 0; j < n; j++ {
		objVal += c[j] * x[j]
	}

	return x, objVal
}

func Factory2(input []string) int {
	total := 0
	for _, line := range input {
		buttons, targets := parseJoltage(line)
		total += minJoltagePresses(buttons, targets)
	}
	return total
}
