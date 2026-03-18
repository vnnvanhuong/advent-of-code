# Day 7: Laboratories Walkthrough

*2026-03-18T03:24:39Z by Showboat dev*
<!-- showboat-id: d4cda085-2aa5-43b3-b9c7-eea101ab4ea3 -->

## Problem Summary

**Input:** A 2D grid representing a tachyon manifold. It contains empty space (`.`), a single starting point (`S`), and splitters (`^`).

**Rules:**
- A tachyon beam starts at `S` and travels continuously downward.
- The beam passes freely through empty space (`.`).
- When a beam encounters a splitter (`^`), it stops and spawns two new beams: one immediately to the left of the splitter, and one immediately to the right.
- These new beams also continue traveling downward.
- If multiple beams overlap (e.g., two splitters dump a beam into the same space between them), they effectively merge into a single beam at that location.

**Goal:** Calculate the total number of times a tachyon beam is split. In other words, count the total number of splitters (`^`) that are hit by at least one beam.

## Test Cases

I've written the test cases in `laboratory_test.go` using the Red/Green TDD approach. Here are the test cases defined:

1. **Sample from puzzle description:** The main example provided in the problem, expecting 21 splitters hit.
2. **No splitters:** A simple grid with just the start `S` and empty space `.`, expecting 0 splitters hit.
3. **Single splitter:** A grid with one `S` and one `^` directly below it, expecting 1 splitter hit.
4. **Overlapping beams:** A small grid designed to test the overlapping logic where two splitters drop beams into the same space, expecting 4 splitters hit.

Running the tests currently fails (Red phase) because the `Laboratory` function hasn't been implemented yet:

```bash
$ go test -v ./aoc2025 -run TestLaboratory
# nguyenvanhuong.vn/adventofcode/aoc2025_test [nguyenvanhuong.vn/adventofcode/aoc2025.test]
aoc2025\laboratory_test.go:31:21: undefined: aoc2025.Laboratory
aoc2025\laboratory_test.go:44:21: undefined: aoc2025.Laboratory
aoc2025\laboratory_test.go:57:21: undefined: aoc2025.Laboratory
aoc2025\laboratory_test.go:76:21: undefined: aoc2025.Laboratory
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025 [build failed]
FAIL
```

## Logical Solution

To solve this problem, we can simulate the flow of the tachyon beams using a Breadth-First Search (BFS) approach.

1.  **Find the Start:** First, we scan the grid to find the coordinates `(x, y)` of the starting point `S`.
2.  **State Tracking:** We will use a queue to keep track of the active beams. Each beam is represented by its current position `(x, y)`. We also need a `visited` set to keep track of positions we've already processed to avoid redundant calculations when beams overlap. Finally, we need a `hit_splitters` set to store the unique coordinates of the splitters that get hit.
3.  **Simulation Loop:**
    *   Pop a beam position `(x, y)` from the queue.
    *   Calculate the next position by moving down: `next_x = x`, `next_y = y + 1`.
    *   Check bounds: If `next_y` is past the bottom of the grid, or `next_x` is outside the horizontal bounds, the beam exits the manifold, so we just continue to the next beam in the queue.
    *   Check the grid character at `(next_x, next_y)`:
        *   If it's empty space `.` (or `S`), the beam continues straight. We add `(next_x, next_y)` to the queue if it hasn't been visited.
        *   If it's a splitter `^`:
            *   We add `(next_x, next_y)` to our `hit_splitters` set.
            *   The beam stops and spawns two new beams to its immediate left and right: `(next_x - 1, next_y)` and `(next_x + 1, next_y)`.
            *   We add these two new positions to the queue if they haven't been visited.
4.  **Result:** Once the queue is empty, all beams have exited the manifold. The answer is the size of the `hit_splitters` set.

### Complexity
*   **Time Complexity:** $O(W \times H)$ where $W$ is the width and $H$ is the height of the grid. In the worst case, beams could visit every cell in the grid. Since we use a `visited` set, each cell is processed at most once.
*   **Space Complexity:** $O(W \times H)$ to store the `visited` set, the `hit_splitters` set, and the queue, which in the worst case could grow proportionally to the grid size.

## Implementation and Testing

I've implemented the `Laboratory` function in `laboratory.go` using the BFS approach described above.

The tests now pass successfully (Green phase):

```bash
$ go test -v ./aoc2025 -run TestLaboratory
=== RUN   TestLaboratory
=== RUN   TestLaboratory/sample_from_puzzle_description
=== RUN   TestLaboratory/no_splitters
=== RUN   TestLaboratory/single_splitter
=== RUN   TestLaboratory/overlapping_beams
--- PASS: TestLaboratory (0.00s)
    --- PASS: TestLaboratory/sample_from_puzzle_description (0.00s)
    --- PASS: TestLaboratory/no_splitters (0.00s)
    --- PASS: TestLaboratory/single_splitter (0.00s)
    --- PASS: TestLaboratory/overlapping_beams (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	1.018s
```

### Time and Space Complexity
*   **Time Complexity:** $O(W \times H)$ where $W$ is the grid width and $H$ is the grid height. Each cell is added to the queue at most once because of the `visited` map.
*   **Space Complexity:** $O(W \times H)$ for the `visited` map, `hitSplitters` map, and the `queue`.

The solution is already quite optimal as it only visits each cell once. We could potentially optimize space slightly by using a 2D boolean array instead of a map for `visited`, but for typical Advent of Code grid sizes, a map is usually fine and easier to write.

### Takeaway
This problem is a classic grid traversal simulation. The key to solving it efficiently is recognizing that beams can overlap and merge. By using a `visited` set to track which cells have already been processed, we avoid exponential explosion of paths and keep the time complexity linear with respect to the grid size. This is a common pattern in BFS/DFS problems where multiple paths can lead to the same state.

## Running the Solution

Finally, we hook up the `Laboratory` function to the main program, reading the actual puzzle input from `laboratory.txt`:

```go
	manifold := aoc2025.ReadLines("aoc2025/laboratory.txt")
	fmt.Println("Laboratory 1:", aoc2025.Laboratory(manifold))
```

Running the program gives us the final answer for Part One:

```bash
$ go run main.go
...
Laboratory 1: [REDACTED]
```

## Part Two: Problem Summary

**New Concept:** Quantum Tachyon Manifold (Many-Worlds Interpretation)

**Rules:**
- Instead of beams merging when they overlap, we are tracking a single *particle*.
- When the particle hits a splitter (`^`), the timeline splits into two. In one timeline, the particle goes left; in the other, it goes right.
- Timelines *do not merge*. If two different paths lead to the same location at the same time, they are still considered separate timelines.
- The particle continues until it exits the manifold (falls off the bottom).

**Goal:** Calculate the total number of distinct timelines active after the particle completes its journey. In other words, count the total number of paths the particle can take from the start `S` to the bottom of the grid.

## Part Two: Test Cases

I've added the test cases for Part Two in `laboratory_test.go` using the Red/Green TDD approach. Here are the test cases defined:

1. **Sample from puzzle description:** The main example provided in the problem, expecting 40 timelines.
2. **No splitters:** A simple grid with just the start `S` and empty space `.`, expecting 1 timeline (the particle just goes straight down).
3. **Single splitter:** A grid with one `S` and one `^` directly below it, expecting 2 timelines (one left, one right).
4. **Overlapping beams:** The same small grid used in Part One, but now we count timelines. We expect 6 timelines to exit the manifold.

Running the tests currently fails (Red phase) because the `Laboratory2` function hasn't been implemented yet:

```bash
$ go test -v ./aoc2025 -run TestLaboratory2
# nguyenvanhuong.vn/adventofcode/aoc2025_test [nguyenvanhuong.vn/adventofcode/aoc2025.test]
aoc2025\laboratory_test.go:104:21: undefined: aoc2025.Laboratory2
aoc2025\laboratory_test.go:117:21: undefined: aoc2025.Laboratory2
aoc2025\laboratory_test.go:130:21: undefined: aoc2025.Laboratory2
aoc2025\laboratory_test.go:152:21: undefined: aoc2025.Laboratory2
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025 [build failed]
FAIL
```

## Part Two: Logical Solution

To solve this problem, we can use Dynamic Programming (DP) to count the number of paths reaching each cell. Since the particle only moves downward, we can process the grid row by row from top to bottom.

1.  **State Tracking:** We maintain an array `dp` where `dp[x]` represents the number of timelines currently at column `x` on the current row.
2.  **Initialization:** We find the starting position `S` at `(startX, startY)`. We initialize `dp[startX] = 1` and all other entries to `0`. We start processing from row `startY`.
3.  **Simulation Loop:**
    *   For each row `y` from `startY` to `height - 1`:
        *   Create a new array `next_dp` initialized to `0` to store the number of timelines reaching the next row.
        *   For each column `x` from `0` to `width - 1`:
            *   If `dp[x] > 0`:
                *   If `manifold[y][x] == '.'` or `'S'`: The particle continues straight down. We add `dp[x]` to `next_dp[x]`.
                *   If `manifold[y][x] == '^'`: The particle splits. We add `dp[x]` to `next_dp[x-1]` and `dp[x]` to `next_dp[x+1]`.
        *   Update `dp = next_dp`.
4.  **Result:** After processing all rows, the `dp` array contains the number of timelines that exited the manifold at each column. The total number of timelines is the sum of all values in the `dp` array.

### Complexity
*   **Time Complexity:** $O(W \times H)$ where $W$ is the width and $H$ is the height of the grid. We process each cell exactly once.
*   **Space Complexity:** $O(W)$ to store the `dp` and `next_dp` arrays for the current and next row. This is an improvement over Part One's space complexity since we only need to keep track of the current row's state.

## Part Two: Implementation and Testing

I've implemented the `Laboratory2` function in `laboratory.go` using the Dynamic Programming approach described above.

The tests now pass successfully (Green phase):

```bash
$ go test -v ./aoc2025 -run TestLaboratory2
=== RUN   TestLaboratory2
=== RUN   TestLaboratory2/sample_from_puzzle_description
=== RUN   TestLaboratory2/no_splitters
=== RUN   TestLaboratory2/single_splitter
=== RUN   TestLaboratory2/overlapping_beams
--- PASS: TestLaboratory2 (0.00s)
    --- PASS: TestLaboratory2/sample_from_puzzle_description (0.00s)
    --- PASS: TestLaboratory2/no_splitters (0.00s)
    --- PASS: TestLaboratory2/single_splitter (0.00s)
    --- PASS: TestLaboratory2/overlapping_beams (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.969s
```

### Time and Space Complexity
*   **Time Complexity:** $O(W \times H)$ where $W$ is the grid width and $H$ is the grid height. We iterate through each cell of the grid exactly once.
*   **Space Complexity:** $O(W)$ for the `dp` and `nextDP` arrays. We only need to store the state of the current row to compute the next row, making it very memory efficient.

### Takeaway
Part Two introduces a shift from tracking *unique visited locations* to tracking *total number of paths*. This is a classic indicator to switch from a simple BFS/DFS with a `visited` set to a Dynamic Programming approach. Because the movement is strictly downward (a Directed Acyclic Graph - DAG), we can process it row by row, accumulating the number of paths that reach each cell. This keeps both time and space complexity very low, even when the number of paths (timelines) grows exponentially large.

## Running the Solution (Part Two)

We update the main program to also print the result for Part Two:

```go
	manifold := aoc2025.ReadLines("aoc2025/laboratory.txt")
	fmt.Println("Laboratory 1:", aoc2025.Laboratory(manifold))
	fmt.Println("Laboratory 2:", aoc2025.Laboratory2(manifold))
```

Running the program gives us the final answer for Part Two:

```bash
$ go run main.go
...
Laboratory 2: [REDACTED]
```

