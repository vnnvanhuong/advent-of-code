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

