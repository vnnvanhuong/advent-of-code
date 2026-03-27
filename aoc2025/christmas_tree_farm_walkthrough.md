# Day 12: Christmas Tree Farm - Walkthrough

### Step 1 — Problem summary

Parse polyomino shapes from the first section and, for each region line `WxH: counts`, decide whether all required copies of each shape can be placed on a `H×W` grid without overlap on `#` cells. Rotations and reflections are allowed; `.` in a shape does not reserve grid cells.

### Step 2 — Test cases

- Example from the statement: three regions (`4x4` with two shape-4 only; two `12x5` lines differing by one extra shape-4) → **2** regions succeed.

### Step 3 — Approach

1. Convert each shape grid to a set of `#` cell coordinates, normalized so the top-left `#` is at `(0,0)`.
2. Build the **unique** orientations of each shape: four rotations of the base, then flip horizontally and four rotations (dedupe by a canonical key of sorted cells).
3. For each region, expand counts into a list of piece ids (one entry per physical present), sort by **descending** piece size (cell count) to prune the search earlier.
4. **Backtracking**: depth-first placement; for the next piece, try every orientation and every anchor `(r0,c0)` where all cells stay inside the grid and land on currently free cells; mark, recurse, unmark.

### Step 4 — Dry run (example)

- Region 1: two pentominoes (shape 4) in `4×4` — search finds a valid packing.
- Region 2: six pieces in `12×5` — search succeeds.
- Region 3: one more shape-4 than region 2 — total area still fits, but no non-overlapping placement exists; backtracking exhausts and returns false.

### Step 5 — Complexity

Let `P` = number of pieces, `A = W×H`, `O` = max orientations per shape (≤ 8), `S` = max cells per piece. Naive upper bound per region is roughly `O((O·A)^P)` in the worst case; with pruning and small AoC instances this is acceptable. Space: `O(A)` for the occupancy grid plus `O(P)` recursion depth.

### Step 6 — Possible optimizations

- Tighter placement bounds per orientation (only scan anchors where the oriented bounding box fits).
- Column/row bitmask or bitset for small grids; symmetry breaking for identical pieces (optional).
- Early reject if sum of `#` cells exceeds `W×H` (already done).

### Step 7 — Refinements

Heuristic ordering of pieces by size (larger first) implemented in the solution.

### Step 8 — Takeaway

This is **exact polyomino packing** (finite CSP / backtracking). Recognizing rotations/reflections as a small fixed set of transforms and normalizing coordinates keeps the code simple.

### Puzzle answer

**[REDACTED]** — Run `main.go`; do not commit real answers.