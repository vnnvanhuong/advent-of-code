# Printing Department - Day 4

*2026-03-13T06:39:34Z by Showboat dev*
<!-- showboat-id: bb8bbe38-6a12-4a25-90b5-495285f5b8cb -->

## Problem Summary

We have a 2D grid where `@` marks a paper roll and `.` is empty. A forklift can access a roll only if **fewer than four** of its eight surrounding cells also contain rolls. Out-of-bounds cells count as empty. The task: count accessible rolls.

## Logical Solution

For each cell in the grid: skip if not an `@`. Count how many of the 8 neighbours are in-bounds and contain `@`. The 8 direction offsets are all combinations of -1, 0, 1 for row and column deltas, excluding (0,0). If the neighbour count is less than 4, the roll is accessible -- increment the result. Time complexity is O(rows * cols) since every cell inspects at most 8 neighbours. Space is O(1) beyond the input.

## Dry Run

Dense 3x3 grid: every cell is '@'. Corners have 3 neighbours -- accessible. Edges have 5 neighbours -- not accessible. Center has 8 -- not accessible. Result: 4 corners. Matches test expectation.

Puzzle sample spot-check: cell at row 0, col 2 is '@' with neighbours '@','@','@' giving count 3 -- accessible. Cell at row 0, col 7 is '@' with neighbours '@','@','@','@' giving count 4 -- not accessible. Both match the expected output diagram.

## Tests Before Implementation

The stub returns 0. Running the test suite should show failures for every case except the empty grid.

```bash
cd aoc2025 && go test -run TestAccessibleCount -v
```

```output
=== RUN   TestAccessibleCount
=== RUN   TestAccessibleCount/empty_grid_returns_zero
=== RUN   TestAccessibleCount/one_roll_is_accessible
    printing_department_test.go:19: expected 1, got 0
=== RUN   TestAccessibleCount/dense_3x3_block_has_only_corners_accessible
    printing_department_test.go:32: expected 4, got 0
=== RUN   TestAccessibleCount/sample_from_puzzle_description
    printing_department_test.go:51: expected 13, got 0
--- FAIL: TestAccessibleCount (0.00s)
    --- PASS: TestAccessibleCount/empty_grid_returns_zero (0.00s)
    --- FAIL: TestAccessibleCount/one_roll_is_accessible (0.00s)
    --- FAIL: TestAccessibleCount/dense_3x3_block_has_only_corners_accessible (0.00s)
    --- FAIL: TestAccessibleCount/sample_from_puzzle_description (0.00s)
FAIL
exit status 1
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

## Implementation

Iterate every cell. For each '@', scan the 8 neighbours. If the roll-neighbour count stays below 4, it is accessible.

```bash
cd aoc2025 && go test -run TestAccessibleCount -v
```

```output
=== RUN   TestAccessibleCount
=== RUN   TestAccessibleCount/empty_grid_returns_zero
=== RUN   TestAccessibleCount/one_roll_is_accessible
=== RUN   TestAccessibleCount/dense_3x3_block_has_only_corners_accessible
=== RUN   TestAccessibleCount/sample_from_puzzle_description
--- PASS: TestAccessibleCount (0.00s)
    --- PASS: TestAccessibleCount/empty_grid_returns_zero (0.00s)
    --- PASS: TestAccessibleCount/one_roll_is_accessible (0.00s)
    --- PASS: TestAccessibleCount/dense_3x3_block_has_only_corners_accessible (0.00s)
    --- PASS: TestAccessibleCount/sample_from_puzzle_description (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

## Complexity

**Time:** O(R * C) where R and C are the grid dimensions. Each cell checks a fixed 8 neighbours, so the inner loop is O(1). **Space:** O(1) beyond the input grid itself -- no auxiliary data structures needed.

## Optimization Notes

The solution is already optimal at O(R*C). Possible micro-optimizations: early-exit the inner neighbour loop once adj reaches 4, since any count >= 4 means the roll is not accessible. This saves a few comparisons in dense grids but does not change the asymptotic complexity. A sliding-window or prefix-sum approach could precompute neighbour counts, but the constant factor is worse for a fixed 3x3 kernel, making it not worthwhile here.

## Takeaway

This is a classic **grid neighbour-counting** problem -- the same pattern appears in Conway's Game of Life, minesweeper number generation, and cellular automata. The key insight is that checking a fixed-size neighbourhood around each cell keeps the per-cell work constant, giving a clean linear scan over the grid. Boundary handling by bounds-checking is simpler than padding the grid with sentinel values for small kernels like this.

