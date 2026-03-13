package aoc2025

// PrintingDepartment returns the number of paper rolls ('@') in the grid that
// have fewer than four neighbouring rolls among the eight surrounding cells.
// The grid is modelled as a slice of equal‑length strings; positions outside
// the bounds are treated as empty.
func PrintingDepartment(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// Direction table
	// The 8 direction offsets are all combinations of -1, 0, 1 for row and column deltas, excluding (0,0).
	// The X is the current cell
	// (-1,-1) (-1,0) (-1,1)
	// ( 0,-1)   X    ( 0,1)
	// ( 1,-1) ( 1,0) ( 1,1)
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// scan every cell in the grid
	count := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] != '@' {
				continue
			}

			// count the number of adjacent rolls
			// For the current @ cell, iterate through all 8 directions.
			// Compute the neighbour coordinates (nr, nc).
			// The bounds check ar >= 0 && ar < rows && ac >= 0 && ac < cols treats anything outside the grid as empty (not a roll).
			// If the neighbour is in bounds and is @, bump the adj counter.
			adj := 0
			for _, d := range dirs {
				ar, ac := r+d[0], c+d[1]
				if ar >= 0 && ar < rows && ac >= 0 && ac < cols && grid[ar][ac] == '@' {
					adj++
				}
			}

			if adj < 4 {
				count++
			}
		}
	}

	return count
}

// BruteforcePrintingDepartment2 uses Approach 1: round-by-round full scan.
// Each round scans the entire grid, collects accessible rolls, removes them
// simultaneously, and repeats. Time: O(K * R * C).
func BruteforcePrintingDepartment2(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	g := make([][]byte, rows)
	for i, s := range grid {
		g[i] = []byte(s)
	}

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0

	for {
		removal := [][2]int{}
		for r := range rows {
			for c := range cols {
				if g[r][c] != '@' {
					continue
				}

				adj := 0
				for _, d := range dirs {
					ar, ac := r+d[0], c+d[1]
					if ar >= 0 && ar < rows && ac >= 0 && ac < cols && g[ar][ac] == '@' {
						adj++
					}
				}

				if adj < 4 {
					removal = append(removal, [2]int{r, c})
				}
			}
		}

		if len(removal) == 0 {
			break
		}

		for _, pos := range removal {
			r, c := pos[0], pos[1]
			g[r][c] = '.'
		}

		count += len(removal)
	}

	return count
}

// PrintingDepartment2 uses Approach 3: single-pass BFS peeling.
// Precompute neighbour counts, seed a queue with all accessible rolls,
// then process: remove, decrement neighbours, enqueue newly accessible.
// Each cell is processed at most once. Time: O(R * C).
func PrintingDepartment2(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	isRoll := make([][]bool, rows)
	adj := make([][]int, rows)
	for r := range rows {
		isRoll[r] = make([]bool, cols)
		adj[r] = make([]int, cols)
		for c := range cols {
			if grid[r][c] == '@' {
				isRoll[r][c] = true
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if !isRoll[r][c] {
				continue
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && isRoll[nr][nc] {
					adj[r][c]++
				}
			}
		}
	}

	queue := [][2]int{}
	for r := range rows {
		for c := range cols {
			if isRoll[r][c] && adj[r][c] < 4 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	total := 0
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		r, c := pos[0], pos[1]
		if !isRoll[r][c] {
			continue
		}
		isRoll[r][c] = false
		total++
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && isRoll[nr][nc] {
				adj[nr][nc]--
				if adj[nr][nc] < 4 {
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	return total
}
