# Day 6: Trash Compactor - Walkthrough

*2026-03-17T03:09:24Z by Showboat dev*
<!-- showboat-id: 45977a4e-2077-4504-b35d-7eef09f800ad -->

## Problem Summary

We are given a wide text worksheet that contains many arithmetic problems placed side by side.

Each problem consists of:
- A vertical stack of integers.
- A final bottom row containing either `+` or `*`.
- At least one completely blank separator column between neighboring problems.
- Arbitrary left/right alignment of the numbers within a problem, which we should ignore.

Our goal is to:
1. Parse the worksheet into separate vertical problems.
2. Read the numbers belonging to each problem.
3. Read the operator at the bottom of each problem.
4. Evaluate each problem.
5. Add all individual results into one grand total.

## Worked Example

From the puzzle statement, this worksheet:

```text
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + 
```

represents four separate problems:
- `123 * 45 * 6 = 33210`
- `328 + 64 + 98 = 490`
- `51 * 387 * 215 = 4243455`
- `64 + 23 + 314 = 401`

The grand total is:
- `33210 + 490 + 4243455 + 401 = 4277556`

## Key Parsing Observations

- We should think of the input as a character grid, not a line-by-line list of equations.
- Problem boundaries are vertical and are identified by blank separator columns.
- Number alignment does not matter, so the eventual parser should recover contiguous digit groups inside each problem region.
- The last row provides the operator for each problem.

## TDD Red Phase

Before writing any implementation, I added `trash_compactor_test.go` with focused failing tests for a proposed public entrypoint `TrashCompactor([]string) int`.

The tests currently cover:
1. The full sample from the puzzle description, expecting `4277556`.
2. A single addition problem.
3. A single multiplication problem.
4. Two narrow problems separated by exactly one blank column.
5. Misaligned numbers inside adjacent problems.

These tests were chosen to force the future implementation to handle both evaluation and parsing concerns, especially separator columns and ignored alignment.

At this point there is still no implementation, so the correct red-phase result is a failing test run due to `aoc2025.TrashCompactor` being undefined.

```bash
go test ./aoc2025 -run TestTrashCompactor
```

```output
# nguyenvanhuong.vn/adventofcode/aoc2025_test [nguyenvanhuong.vn/adventofcode/aoc2025.test]
aoc2025/trash_compactor_test.go:19:21: undefined: aoc2025.TrashCompactor
aoc2025/trash_compactor_test.go:32:21: undefined: aoc2025.TrashCompactor
aoc2025/trash_compactor_test.go:46:21: undefined: aoc2025.TrashCompactor
aoc2025/trash_compactor_test.go:59:21: undefined: aoc2025.TrashCompactor
aoc2025/trash_compactor_test.go:72:21: undefined: aoc2025.TrashCompactor
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025 [build failed]
FAIL
```

## TDD Green Phase

Implemented `TrashCompactor([]string) int` in `trash_compactor.go` using the column-segmentation approach described earlier.

### Algorithm walkthrough

1. Pad every row to the same width so the worksheet can be treated as a rectangular grid.
2. Scan columns from left to right.
3. Group consecutive non-blank columns into problem ranges.
4. For each problem range:
   - Read the bottom-row operator (`+` or `*`).
   - Parse every contiguous digit run from the rows above within that range.
   - Evaluate the numbers using the operator.
5. Add the result of every problem into the grand total.

### Why this matches the puzzle

- A fully blank column separates adjacent problems.
- Number alignment inside a problem is irrelevant.
- So the natural first step is to isolate each problem by columns, then parse its numbers locally.

### Complexity

Let `R` be the number of rows and `C` the maximum row width.

- Time: `O(R * C)`
- Space: `O(R * C)` for the padded grid

### Notes

This first green implementation is intentionally simple and direct. It handles the sample, both operators, narrow adjacent problems, and misaligned numbers.

```bash
go test ./aoc2025 -run TestTrashCompactor
```

```output
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.006s
```

## Main Program Integration

Updated `main.go` to read `aoc2025/trash_compactor.txt` and print the result of `TrashCompactor(...)` when running the repo entrypoint.

For reproducibility, the command to run is:

```bash
go run .
```

The walkthrough intentionally omits the numeric Trash Compactor answer, per request.

