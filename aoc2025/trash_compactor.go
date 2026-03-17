package aoc2025

import "strconv"

// TrashCompactor parses the worksheet, evaluates each vertical problem,
// and returns the grand total of all problem results.
func TrashCompactor(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	grid := padWorksheet(lines)
	ranges := problemRanges(grid)

	total := 0
	for _, r := range ranges {
		total += evaluateProblem(grid, r[0], r[1])
	}

	return total
}

// TrashCompactor2 parses the worksheet using the part-two rules, where each
// number is read as a vertical column and problems are decoded right-to-left.
func TrashCompactor2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	grid := padWorksheet(lines)
	ranges := problemRanges(grid)

	total := 0
	for _, r := range ranges {
		operator := findOperator(grid[len(grid)-1], r[0], r[1])
		numbers := parseNumbersByColumns(grid[:len(grid)-1], r[0], r[1])
		total += evaluateNumbers(operator, numbers)
	}

	return total
}

func padWorksheet(lines []string) []string {
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	grid := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < width {
			padding := make([]byte, width-len(line))
			for j := range padding {
				padding[j] = ' '
			}
			grid[i] = line + string(padding)
			continue
		}
		grid[i] = line
	}

	return grid
}

func problemRanges(grid []string) [][2]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}

	width := len(grid[0])
	var ranges [][2]int
	start := -1

	for col := 0; col < width; col++ {
		active := false
		for _, row := range grid {
			if row[col] != ' ' {
				active = true
				break
			}
		}

		if active && start == -1 {
			start = col
		}
		if !active && start != -1 {
			ranges = append(ranges, [2]int{start, col - 1})
			start = -1
		}
	}

	if start != -1 {
		ranges = append(ranges, [2]int{start, width - 1})
	}

	return ranges
}

func evaluateProblem(grid []string, left, right int) int {
	operator := findOperator(grid[len(grid)-1], left, right)
	numbers := parseNumbers(grid[:len(grid)-1], left, right)
	return evaluateNumbers(operator, numbers)
}

func evaluateNumbers(operator byte, numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if operator == '+' {
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum
	}

	product := 1
	for _, n := range numbers {
		product *= n
	}
	return product
}

func findOperator(row string, left, right int) byte {
	for col := left; col <= right; col++ {
		if row[col] == '+' || row[col] == '*' {
			return row[col]
		}
	}
	return 0
}

func parseNumbers(rows []string, left, right int) []int {
	var numbers []int

	for _, row := range rows {
		col := left
		for col <= right {
			for col <= right && (row[col] < '0' || row[col] > '9') {
				col++
			}
			if col > right {
				break
			}

			start := col
			for col <= right && row[col] >= '0' && row[col] <= '9' {
				col++
			}

			value, _ := strconv.Atoi(row[start:col])
			numbers = append(numbers, value)
		}
	}

	return numbers
}

func parseNumbersByColumns(rows []string, left, right int) []int {
	var numbers []int

	for col := right; col >= left; col-- {
		digits := make([]byte, 0, len(rows))
		for _, row := range rows {
			if row[col] >= '0' && row[col] <= '9' {
				digits = append(digits, row[col])
			}
		}

		if len(digits) == 0 {
			continue
		}

		value, _ := strconv.Atoi(string(digits))
		numbers = append(numbers, value)
	}

	return numbers
}
