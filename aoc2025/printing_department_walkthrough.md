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

## Part 2

### Problem Summary: 

Part 2 extends the single-pass count into an iterative removal process. Each round, identify every roll that is currently accessible (fewer than 4 neighbours), remove all of them at once, then repeat on the updated grid. Keep going until no more rolls are accessible. The answer is the **total number of rolls removed across all rounds**. In the example, the rounds remove 13, 12, 7, 5, 2, 1, 1, 1, 1 rolls for a total of **43**.

### Red Phase: add Part 2 tests with a stub that returns 0.

```bash
cd aoc2025 && go test -run TestPrintingDepartment2 -v
```

```output
=== RUN   TestPrintingDepartment2
=== RUN   TestPrintingDepartment2/empty_grid_returns_zero
=== RUN   TestPrintingDepartment2/single_roll_removed_in_one_round
    printing_department_test.go:66: expected 1, got 0
=== RUN   TestPrintingDepartment2/sparse_row_all_removed_at_once
    printing_department_test.go:74: expected 3, got 0
=== RUN   TestPrintingDepartment2/3x3_block_fully_eroded_in_two_rounds
    printing_department_test.go:88: expected 9, got 0
=== RUN   TestPrintingDepartment2/sample_from_puzzle_description
    printing_department_test.go:107: expected 43, got 0
--- FAIL: TestPrintingDepartment2 (0.00s)
    --- PASS: TestPrintingDepartment2/empty_grid_returns_zero (0.00s)
    --- FAIL: TestPrintingDepartment2/single_roll_removed_in_one_round (0.00s)
    --- FAIL: TestPrintingDepartment2/sparse_row_all_removed_at_once (0.00s)
    --- FAIL: TestPrintingDepartment2/3x3_block_fully_eroded_in_two_rounds (0.00s)
    --- FAIL: TestPrintingDepartment2/sample_from_puzzle_description (0.00s)
FAIL
exit status 1
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025	0.009s
```

### Proposed Approaches (simple to complex)

**Approach 1 -- Round-by-round full scan.** Each round, scan every cell, collect all accessible rolls, remove them simultaneously, repeat until nothing is removable. Time: O(K * R * C) where K is number of rounds. Pros: dead simple, mirrors the problem statement, easy to debug. Cons: re-scans the entire grid each round including empty and interior cells.

**Approach 2 -- Round-by-round with candidate tracking.** Maintain a candidate set of cells that might become accessible. After removing a batch, only neighbours of removed cells become candidates for the next round. Time: O(T * 8) where T is total rolls removed. Pros: avoids scanning empty/interior cells, still gives round-by-round structure. Cons: more bookkeeping, cells can appear in the candidate set multiple times.

**Approach 3 -- Single-pass BFS peeling.** Precompute neighbour counts, seed a queue with all initially accessible rolls. Dequeue a cell, remove it, decrement its neighbours' counts, enqueue any whose count drops below 4. Each cell processed at most once. Time: O(R * C). Pros: optimal, elegant, same pattern as topological sort / Kahn's algorithm. Cons: requires an auxiliary count array, does not naturally separate rounds.

**Recommendation:** Approach 3 gives optimal O(R*C) time with the same O(R*C) space as the others. The code is only slightly more complex than Approach 1 and follows a well-known graph peeling pattern.

### Green Phase: Approach 1 implemented and all tests pass.

```bash
cd aoc2025 && go test -run TestPrintingDepartment -v
```

```output
=== RUN   TestPrintingDepartment1
=== RUN   TestPrintingDepartment1/empty_grid_returns_zero
=== RUN   TestPrintingDepartment1/one_roll_is_accessible
=== RUN   TestPrintingDepartment1/dense_3x3_block_has_only_corners_accessible
=== RUN   TestPrintingDepartment1/sample_from_puzzle_description
--- PASS: TestPrintingDepartment1 (0.00s)
    --- PASS: TestPrintingDepartment1/empty_grid_returns_zero (0.00s)
    --- PASS: TestPrintingDepartment1/one_roll_is_accessible (0.00s)
    --- PASS: TestPrintingDepartment1/dense_3x3_block_has_only_corners_accessible (0.00s)
    --- PASS: TestPrintingDepartment1/sample_from_puzzle_description (0.00s)
=== RUN   TestPrintingDepartment2
=== RUN   TestPrintingDepartment2/empty_grid_returns_zero
=== RUN   TestPrintingDepartment2/single_roll_removed_in_one_round
=== RUN   TestPrintingDepartment2/sparse_row_all_removed_at_once
=== RUN   TestPrintingDepartment2/3x3_block_fully_eroded_in_two_rounds
=== RUN   TestPrintingDepartment2/sample_from_puzzle_description
--- PASS: TestPrintingDepartment2 (0.00s)
    --- PASS: TestPrintingDepartment2/empty_grid_returns_zero (0.00s)
    --- PASS: TestPrintingDepartment2/single_roll_removed_in_one_round (0.00s)
    --- PASS: TestPrintingDepartment2/sparse_row_all_removed_at_once (0.00s)
    --- PASS: TestPrintingDepartment2/3x3_block_fully_eroded_in_two_rounds (0.00s)
    --- PASS: TestPrintingDepartment2/sample_from_puzzle_description (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.006s
```

### Refactor: 
Renamed `PrintingDepartment2` to `BruteforcePrintingDepartment2` and implemented Approach 3 (BFS peeling) as the new PrintingDepartment2. Precompute neighbour counts, seed queue with all initially accessible rolls, then process: dequeue, remove, decrement neighbours, enqueue any that drop below threshold. Each cell processed at most once.

```bash
cd aoc2025 && go test -run TestPrintingDepartment -v
```

```output
=== RUN   TestPrintingDepartment1
=== RUN   TestPrintingDepartment1/empty_grid_returns_zero
=== RUN   TestPrintingDepartment1/one_roll_is_accessible
=== RUN   TestPrintingDepartment1/dense_3x3_block_has_only_corners_accessible
=== RUN   TestPrintingDepartment1/sample_from_puzzle_description
--- PASS: TestPrintingDepartment1 (0.00s)
    --- PASS: TestPrintingDepartment1/empty_grid_returns_zero (0.00s)
    --- PASS: TestPrintingDepartment1/one_roll_is_accessible (0.00s)
    --- PASS: TestPrintingDepartment1/dense_3x3_block_has_only_corners_accessible (0.00s)
    --- PASS: TestPrintingDepartment1/sample_from_puzzle_description (0.00s)
=== RUN   TestPrintingDepartment2
=== RUN   TestPrintingDepartment2/empty_grid_returns_zero
=== RUN   TestPrintingDepartment2/single_roll_removed_in_one_round
=== RUN   TestPrintingDepartment2/sparse_row_all_removed_at_once
=== RUN   TestPrintingDepartment2/3x3_block_fully_eroded_in_two_rounds
=== RUN   TestPrintingDepartment2/sample_from_puzzle_description
--- PASS: TestPrintingDepartment2 (0.00s)
    --- PASS: TestPrintingDepartment2/empty_grid_returns_zero (0.00s)
    --- PASS: TestPrintingDepartment2/single_roll_removed_in_one_round (0.00s)
    --- PASS: TestPrintingDepartment2/sparse_row_all_removed_at_once (0.00s)
    --- PASS: TestPrintingDepartment2/3x3_block_fully_eroded_in_two_rounds (0.00s)
    --- PASS: TestPrintingDepartment2/sample_from_puzzle_description (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

### Part 2 Complexity

**Brute-force (Approach 1):** Time O(K * R * C) where K is the number of erosion rounds. Space O(R * C) for the mutable byte grid. 

**BFS peeling (Approach 3):** Time O(R * C) -- each cell enqueued and processed at most once. Space O(R * C) for the adjacency count array and queue. Both produce identical results; the BFS approach eliminates redundant full-grid scans.

## Takeaway

Part 1 is a straightforward **grid neighbour-counting** problem (same family as minesweeper, Game of Life). Part 2 turns it into **iterative erosion / graph peeling** -- repeatedly removing nodes whose degree falls below a threshold. The optimal approach is a BFS queue seeded with initially removable cells, processing each cell exactly once as its neighbours' counts are decremented. This is the same algorithmic pattern as **Kahn's algorithm** for topological sort and **k-core decomposition** in graph theory, where nodes with degree below k are peeled away layer by layer. Key lessons: (1) when a brute-force simulation re-scans unchanged regions, a queue-based approach can eliminate redundant work; (2) precomputing derived state (neighbour counts) and maintaining it incrementally is cheaper than recomputing from scratch each round.
