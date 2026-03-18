package aoc2025

// Laboratory calculates the total number of splitters hit by a tachyon beam.
func Laboratory(manifold []string) int {
	if len(manifold) == 0 {
		return 0
	}

	height := len(manifold)
	width := len(manifold[0])

	// Find the starting position 'S'
	var startX, startY int
	found := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if manifold[y][x] == 'S' {
				startX, startY = x, y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return 0 // No start found
	}

	type pos struct {
		x, y int
	}

	queue := []pos{{startX, startY}}
	visited := make(map[pos]bool)
	visited[pos{startX, startY}] = true
	hitSplitters := make(map[pos]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		nextY := curr.y + 1
		nextX := curr.x

		// Check bounds
		if nextY >= height || nextX < 0 || nextX >= width {
			continue
		}

		nextPos := pos{nextX, nextY}
		char := manifold[nextY][nextX]

		if char == '.' || char == 'S' {
			if !visited[nextPos] {
				visited[nextPos] = true
				queue = append(queue, nextPos)
			}
		} else if char == '^' {
			hitSplitters[nextPos] = true

			leftPos := pos{nextX - 1, nextY}
			if !visited[leftPos] {
				visited[leftPos] = true
				queue = append(queue, leftPos)
			}

			rightPos := pos{nextX + 1, nextY}
			if !visited[rightPos] {
				visited[rightPos] = true
				queue = append(queue, rightPos)
			}
		}
	}

	return len(hitSplitters)
}

// Laboratory2 calculates the total number of timelines a single tachyon particle ends up on.
func Laboratory2(manifold []string) int {
	if len(manifold) == 0 {
		return 0
	}

	height := len(manifold)
	width := len(manifold[0])

	// Find the starting position 'S'
	var startX, startY int
	found := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if manifold[y][x] == 'S' {
				startX, startY = x, y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return 0 // No start found
	}

	// dp[x] stores the number of timelines at column x on the current row
	dp := make([]int, width)
	dp[startX] = 1

	for y := startY; y < height; y++ {
		nextDP := make([]int, width)
		for x := 0; x < width; x++ {
			if dp[x] > 0 {
				char := manifold[y][x]
				if char == '.' || char == 'S' {
					nextDP[x] += dp[x]
				} else if char == '^' {
					if x-1 >= 0 {
						nextDP[x-1] += dp[x]
					}
					if x+1 < width {
						nextDP[x+1] += dp[x]
					}
				}
			}
		}
		dp = nextDP
	}

	totalTimelines := 0
	for _, count := range dp {
		totalTimelines += count
	}

	return totalTimelines
}
