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
