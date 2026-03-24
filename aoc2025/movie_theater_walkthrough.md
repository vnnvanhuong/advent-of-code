# Day 9: Movie Theater - Walkthrough

*2026-03-24T02:19:52Z by Showboat dev*
<!-- showboat-id: fa246ad1-c2da-4408-b175-e44c1e11bf0b -->

## Problem Summary

**Input:** A list of 2D coordinates representing red tiles on a grid.

**Process:**

1. Consider all possible pairs of red tiles.
2. For each pair, treat the two tiles as opposite corners of an axis-aligned rectangle.
3. Compute the area of each rectangle as `(|x1 - x2| + 1) * (|y1 - y2| + 1)` (tile-inclusive).

**Output:** The largest area among all such rectangles.

## Key Observations

- The rectangle area includes both endpoint tiles, so the formula is `(|x1 - x2| + 1) * (|y1 - y2| + 1)`, not just `|x1 - x2| * |y1 - y2|`.
- Any two red tiles can be chosen as opposite corners, regardless of orientation.
- The brute-force O(N^2) approach checking all pairs is sufficient since N (number of red tiles) is small.

## Example Walkthrough

Given 8 red tiles, here are some rectangle candidates:

| Pair                    | Width          | Height         | Area |
|-------------------------|----------------|----------------|------|
| (2,5) and (9,7)        | \|2-9\|+1 = 8   | \|5-7\|+1 = 3   | 24   |
| (7,1) and (11,7)       | \|7-11\|+1 = 5  | \|1-7\|+1 = 7   | 35   |
| (7,3) and (2,3)        | \|7-2\|+1 = 6   | \|3-3\|+1 = 1   | 6    |
| (2,5) and (11,1)       | \|2-11\|+1 = 10 | \|5-1\|+1 = 5   | 50   |

The largest rectangle has area **50**, formed by (2,5) and (11,1).

## Test Cases

Function signature: `MovieTheater(input []string) int`

| # | Test Name                       | Tiles | Expected    | Rationale                                                           |
|---|---------------------------------|-------|-------------|---------------------------------------------------------------------|
| 1 | Sample from puzzle description  | 8     | **50**      | Directly from the problem: (2,5)-(11,1) = 10*5                     |
| 2 | Two points horizontal line      | 2     | **5**       | (1,1)-(5,1): width 5, height 1                                     |
| 3 | Two points vertical line        | 2     | **5**       | (2,2)-(2,6): width 1, height 5                                     |
| 4 | Single point                    | 1     | **0**       | Cannot form a rectangle with one tile                               |
| 5 | Large coordinates               | 2     | **1002001** | (0,0)-(1000,1000): 1001*1001                                       |

### Red phase (TDD)

All 5 tests fail with `undefined: aoc2025.MovieTheater` or return 0 from the stub. This confirms we are in the **red** state before implementing.

### Green phase (TDD)

All 5 tests pass after implementing `MovieTheater` in `movie_theater.go`. Red -> Green confirmed.

## Algorithmic Approach

1. **Parse** each `"X,Y"` string into integer coordinates.
2. **Generate all pairs** `(i, j)` where `i < j`.
3. **Compute area** for each pair: `(|x1 - x2| + 1) * (|y1 - y2| + 1)`.
4. **Track the maximum** area across all pairs.

### Dry-run: Test #2 — Two points horizontal line

Tiles: `(1,1)`, `(5,1)`

Only one pair exists:
- dx = |1-5| = 4, dy = |1-1| = 0
- Area = (4+1) * (0+1) = 5 * 1 = **5** ✓

### Time & Space Complexity

- **Pairs:** For `n` tiles, there are `n*(n-1)/2` pairs — O(n²).
- **Per pair:** O(1) computation.
- **Total:** O(n²) time, O(n) space for storing parsed points.

---

## Part Two

### Problem Summary

The red tiles now form the vertices of a **rectilinear polygon** (each consecutive pair of tiles is connected by a straight horizontal or vertical line of green tiles, and the list wraps around). The interior of this polygon is also green. A rectangle is only valid if **all** its tiles are red or green (i.e., the rectangle must be fully contained within the polygon).

**Output:** The largest area of a rectangle with two red tile corners that fits entirely inside the polygon.

### Key Observations (Part Two)

- Checking every tile inside a candidate rectangle is too slow for large coordinates (up to 100,000).
- A rectangle is fully contained in the rectilinear polygon if and only if:
  1. Its **center point** is inside the polygon (ray casting).
  2. **No polygon edge** strictly intersects the interior of the rectangle.
- For horizontal edges: edge at `y` from `x1` to `x2` intersects rectangle `[rx1,rx2] x [ry1,ry2]` if `ry1 < y < ry2` and the x-intervals overlap.
- For vertical edges: edge at `x` from `y1` to `y2` intersects rectangle `[rx1,rx2] x [ry1,ry2]` if `rx1 < x < rx2` and the y-intervals overlap.
- This makes each rectangle check O(E) where E is the number of polygon edges, instead of O(Area).

### Test Cases (Part Two)

Function signature: `MovieTheater2(input []string) int`

| # | Test Name                       | Tiles | Expected | Rationale                                                                |
|---|---------------------------------|-------|----------|--------------------------------------------------------------------------|
| 1 | Sample from puzzle description  | 8     | **24**   | Directly from the problem: (9,5)-(2,3) = 8*3                            |
| 2 | C-shaped polygon                | 8     | **33**   | Full span blocked by concavity; best fit is (0,0)-(10,2) = 11*3         |

### Red phase (TDD)

Both tests fail with `undefined: aoc2025.MovieTheater2`. Red state confirmed.

### Green phase (TDD)

All 7 tests pass (5 Part One + 2 Part Two). Red -> Green confirmed.

### Algorithmic Approach (Part Two)

Same O(N^2) pair enumeration as Part One, but with an added containment check per pair:

1. **Parse** coordinates into a polygon vertex list.
2. **For each pair** of vertices `(i, j)`:
   a. Compute the candidate rectangle bounds `[minX, maxX] x [minY, maxY]`.
   b. **Center check:** Use ray casting (count vertical edges to the right of the center with matching y-range) to determine if the center is inside the polygon. If not, skip.
   c. **Edge intersection check:** For each polygon edge, check if it strictly crosses through the interior of the rectangle. If any edge does, the rectangle is invalid.
   d. If valid, compute area and update the maximum.

### Dry-run: Test #2 - C-shaped polygon

Vertices: `(0,0)`, `(10,0)`, `(10,10)`, `(0,10)`, `(0,8)`, `(8,8)`, `(8,2)`, `(0,2)`

Candidate: (0,0) to (10,10) - area 121
- Center (5,5): ray cast right finds 0 vertical edges to the right with matching y -> outside. **Rejected.**

Candidate: (0,0) to (10,0) - area 11*1 = 11
- Center (5,0): on boundary -> inside. No edges cross interior (height is 0). **Valid.** Max = 11.

Candidate: (0,0) to (0,2) - area 1*3 = 3
- Degenerate (width 1). **Valid.** Max still 11.

Candidate: (10,0) to (0,2) - area 11*3 = 33
- Center (5,1): ray cast right -> vertical edge at x=10 with y in [0,10] -> 1 intersection -> inside.
- Edge check: horizontal edge y=8 -> 0 < 8 < 2? No. Horizontal edge y=2 -> 0 < 2 < 2? No. Vertical edge x=8 -> 0 < 8 < 10, y in [2,8], 2 > 0 and 8 > 2? No, max_y=2 so edge y-range is [2,8], and rect y-range is [0,2], so 8 > 0 and 2 > 2? No (not strict). **Valid.** Max = 33.

Final answer: **33**

### Time and Space Complexity (Part Two)

- **Pairs:** O(N^2) where N is the number of red tiles.
- **Per pair:** O(N) to check center + edge intersections (N edges in the polygon).
- **Total:** O(N^3) time, O(N) space.

For N ~ 500 tiles, this is ~125 million operations, which completes in seconds.

### Puzzle Results

- **Part One:** **[REDACTED]**
- **Part Two:** **[REDACTED]**

### Takeaway

- Part One is a straightforward **brute-force enumeration** of all pairs with a simple area formula.
- Part Two introduces a **computational geometry** constraint: checking if a rectangle is contained within a rectilinear polygon. The key insight is that instead of checking every tile, we only need to verify (1) the center is inside via **ray casting** and (2) no polygon edge crosses the rectangle interior, reducing per-check cost from O(Area) to O(N).
