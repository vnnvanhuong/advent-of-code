# Day 8: Playground — Walkthrough

## Problem Summary

**Input:** A list of junction boxes, each with a 3D position `(X, Y, Z)`.

**Process:**

1. Consider all possible pairs of junction boxes and compute the straight-line (Euclidean) distance between each pair.
2. Sort all pairs by distance, shortest first.
3. Process the first **1000** closest pairs in order. For each pair:
   - If the two boxes are in **different** circuits, connect them (merge their circuits).
   - If they are already in the **same** circuit, do nothing — but this still counts as one of the 1000 pairs processed.

**Output:** After processing the 1000 closest pairs, find the sizes of the **three largest circuits** and multiply them together.

## Key Observations

- This is essentially a partial **Kruskal's algorithm** — we sort edges by weight and process them with a **Union-Find** data structure, but we stop after a fixed number of edges (1000) rather than when the spanning tree is complete.
- Redundant connections (pairs already in the same circuit) still consume one of the 1000 slots.
- Distance is standard 3D Euclidean: `sqrt((x2-x1)² + (y2-y1)² + (z2-z1)²)`. Since we only need to compare/sort distances, we can use **squared distance** to avoid the `sqrt`.

## Example Walkthrough

Given 20 junction boxes, process the 10 shortest pairs:

| Step | Pair (by coordinates)                        | Same circuit? | Result                          |
|------|----------------------------------------------|---------------|---------------------------------|
| 1    | `162,817,812` — `425,690,689`                | No            | Merge → circuit of size 2       |
| 2    | `162,817,812` — `431,825,988`                | No            | Merge → circuit of size 3       |
| 3    | `906,360,560` — `805,96,715`                 | No            | Merge → circuit of size 2       |
| 4    | `431,825,988` — `425,690,689`                | **Yes**        | Nothing (redundant)             |
| 5–10 | *(remaining closest pairs)*                  | …             | …                               |

After 10 pairs (one redundant), the circuits are:

- Sizes: **5, 4, 2, 2**, 1, 1, 1, 1, 1, 1, 1 (11 circuits total)
- Three largest: 5 × 4 × 2 = **40**

## Test Cases

Function signature: `Playground(boxes []string, connections int) int`

| # | Test Name                      | Boxes | Connections | Expected | Rationale                                                        |
|---|--------------------------------|-------|-------------|----------|------------------------------------------------------------------|
| 1 | Sample from puzzle description | 20    | 10          | **40**   | Directly from the problem (5×4×2)                                |
| 2 | Two boxes, single connection   | 2     | 1           | **2**    | Single circuit of 2; top-3 = 2×1×1                               |
| 3 | Three boxes, two connections   | 3     | 2           | **3**    | All merge into one circuit; top-3 = 3×1×1                        |
| 4 | Redundant connections consumed | 4     | 3           | **3**    | 3rd pair is redundant but still counts; top-3 = 3×1×1            |
| 5 | Zero connections               | 3     | 0           | **1**    | Every box is its own circuit; top-3 = 1×1×1                      |
| 6 | Fewer than three circuits      | 2     | 0           | **1**    | Only 2 circuits exist; pad missing slots with 1 → 1×1×1          |

### Red phase (TDD)

All 6 tests fail with `undefined: aoc2025.Playground` — the function does not exist yet. This confirms we are in the **red** state before implementing.

### Green phase (TDD)

All 6 tests pass after implementing `Playground` in `playground.go`. Red → Green confirmed.

## Algorithmic Approach

1. **Parse** each `"X,Y,Z"` string into integer coordinates.
2. **Generate all pairs** `(i, j)` with squared Euclidean distance (`dx² + dy² + dz²`).
3. **Sort** pairs by distance ascending.
4. **Process the first `connections` pairs** using Union-Find (path compression + union by rank):
   - For each pair, call `union(a, b)`. If `find(a) == find(b)`, it's a no-op.
5. **Collect circuit sizes** from the Union-Find structure by visiting each unique root.
6. **Sort sizes descending**, take the top 3 (padding with 1 if fewer), and return their product.

### Dry-run: Test #4 — Redundant connections consumed

Boxes: `(0,0,0)`, `(1,0,0)`, `(2,0,0)`, `(100,100,100)` — 3 connections.

Pairs sorted by dist²: `(0,1)=1`, `(1,2)=1`, `(0,2)=4`, `(2,3)=29604`, ...

| Step | Pair  | find(a)==find(b)? | Action           | Circuits         |
|------|-------|-------------------|------------------|------------------|
| 1    | (0,1) | No                | Merge → {0,1}    | {0,1}, {2}, {3}  |
| 2    | (1,2) | No                | Merge → {0,1,2}  | {0,1,2}, {3}     |
| 3    | (0,2) | **Yes**           | Skip (redundant) | {0,1,2}, {3}     |

Top-3 sizes: 3, 1, 1 → product = **3** ✓

### Time & Space Complexity

- **Pairs:** For `n` junction boxes, there are `n*(n-1)/2` pairs — O(n²).
- **Sort:** O(n² log n) for sorting all pairs.
- **Union-Find:** O(k · α(n)) ≈ O(k) for `k` connections — nearly constant per operation.
- **Space:** O(n²) for storing all pairs; O(n) for Union-Find.

---

## Part Two

### Problem Summary

Keep connecting closest pairs until **all** junction boxes form a single circuit. Return the **product of the X coordinates** of the two boxes in the pair that causes the final merge.

### Test Cases (Part Two)

Function signature: `Playground2(boxes []string) int`

| # | Test Name                          | Boxes | Expected  | Rationale                                                              |
|---|------------------------------------|-------|-----------|------------------------------------------------------------------------|
| 1 | Sample from puzzle description     | 20    | **25272** | Last merge: `216,146,977` — `117,168,530` → 216 × 117                 |
| 2 | Two boxes                          | 2     | **21**    | Only one pair `(3,0,0)`—`(7,0,0)` → 3 × 7                            |
| 3 | Three boxes collinear              | 3     | **50**    | Last merge: `(5,0,0)`—`(10,0,0)` → 5 × 10                            |
| 4 | Four boxes with redundant before final | 4 | **300**   | Redundant `(0,2)` skipped; last merge: `(3,0,0)`—`(100,0,0)` → 3×100 |
| 5 | Single box                         | 1     | **0**     | Already one circuit; no pair exists                                    |

### Red phase (TDD)

All 5 tests fail with `undefined: aoc2025.Playground2`. Red state confirmed.

### Green phase (TDD)

All 11 tests pass (6 Part One + 5 Part Two). Red → Green confirmed.

### Algorithmic Approach (Part Two)

Same shared infrastructure as Part One (parsing, edge generation, sorting, Union-Find), but:

1. Process pairs in sorted order.
2. On each **successful** union (different components), decrement component count.
3. When component count reaches **1**, return `points[i].x * points[j].x` for that edge.

The key difference from Part One: we don't process a fixed number of edges — we stop as soon as all boxes are unified.

### Dry-run: Test #4 — Four boxes with redundant pair before final merge

Boxes: `(1,0,0)`, `(2,0,0)`, `(3,0,0)`, `(100,0,0)`

| Step | Pair  | dist² | Merged? | Components | Notes          |
|------|-------|-------|---------|------------|----------------|
| 1    | (0,1) | 1     | Yes     | 3          |                |
| 2    | (1,2) | 1     | Yes     | 2          |                |
| 3    | (0,2) | 4     | No      | 2          | Redundant      |
| 4    | (2,3) | 9409  | Yes     | **1**      | **Final merge** |

Last pair: boxes `(3,0,0)` and `(100,0,0)` → 3 × 100 = **300** ✓

### Puzzle Results

- **Part One** (1000 connections): **[REDACTED]**
- **Part Two** (all connected): **[REDACTED]**

### Takeaway

Both parts are applications of **Kruskal's algorithm** with Union-Find:
- Part One: partial Kruskal — process a fixed number of edges, report circuit sizes.
- Part Two: full Kruskal — run until one component remains, report the final merging edge.

The shared Union-Find with path compression and union by rank keeps per-operation cost at O(α(n)), making both parts efficient even with O(n²) edges.
