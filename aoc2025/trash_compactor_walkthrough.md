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

## Part Two Problem Summary

Part two keeps the same worksheet layout and the same operators, but changes how each number is read.

In part one, each problem was read by rows inside its column block.

In part two, each number must instead be read **right-to-left by columns** inside a problem block:
- each vertical column inside a problem contributes one digit to each number
- digits are ordered from most significant at the top to least significant at the bottom within that column
- the problem itself is read from the **rightmost column toward the leftmost column**
- blank separator columns between problems still define the problem boundaries
- the bottom-row symbol is still the operator for the whole problem

So the major change is not how to split problems, but how to decode the numbers inside each problem.

### Worked example reinterpretation

Using the same sample worksheet:

```text
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +
```

The four problems now become:
- Rightmost problem: `4 + 431 + 623 = 1058`
- Second from right: `175 * 581 * 32 = 3253600`
- Third from right: `8 + 248 + 369 = 625`
- Leftmost problem: `356 * 24 * 1 = 8544`

The sample total for part two is therefore `3263827`.

## Part Two Key Insight

We can likely reuse the part one boundary detection, because separator columns and operators still work the same way.

The new work is inside each problem block:
- scan the columns of that block from right to left
- in each column, collect the non-space digits from top to bottom
- each collected vertical digit string becomes one number
- then apply the same `+` or `*` evaluation as before

This means part two is a decoding change, not a worksheet segmentation change.

## Part Two TDD Red Phase

Added a new `TestTrashCompactor2` suite in `trash_compactor_test.go` for a proposed `TrashCompactor2([]string) int` entrypoint.

The tests cover:
1. The full part two sample from the puzzle description, expecting `3263827`.
2. A dense two-column addition case, proving numbers are formed by columns instead of rows.
3. A dense two-column multiplication case.
4. A sparse-column addition case, proving spaces inside a column are ignored when forming a number.
5. Multiple problems separated by a blank column, showing part one segmentation should still be reusable.

As expected for the red phase, these tests currently fail because `TrashCompactor2` has not been implemented yet.

```bash
go test ./aoc2025 -run TestTrashCompactor2
```

```output
# nguyenvanhuong.vn/adventofcode/aoc2025_test [nguyenvanhuong.vn/adventofcode/aoc2025.test]
aoc2025/trash_compactor_test.go:88:21: undefined: aoc2025.TrashCompactor2
aoc2025/trash_compactor_test.go:101:21: undefined: aoc2025.TrashCompactor2
aoc2025/trash_compactor_test.go:114:21: undefined: aoc2025.TrashCompactor2
aoc2025/trash_compactor_test.go:127:21: undefined: aoc2025.TrashCompactor2
aoc2025/trash_compactor_test.go:140:21: undefined: aoc2025.TrashCompactor2
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025 [build failed]
FAIL
```

## Part Two Logical Solution

The cleanest solution is to keep the part one worksheet segmentation and replace only the number-decoding step.

For each problem block:
1. Reuse the left/right bounds found from blank separator columns.
2. Read the operator from the bottom row exactly as before.
3. Scan the columns in that block from right to left.
4. For each column, collect all non-space digits from top to bottom.
5. If a column contributes any digits, convert that vertical digit string into one number.
6. Evaluate the resulting numbers with the problem operator and add the result to the grand total.

This works because part two changes the orientation of number parsing, but does not change the worksheet boundaries or the evaluation rules.

## Part Two TDD Green Phase

Implemented `TrashCompactor2([]string) int` in `trash_compactor.go`.

### Implementation approach

- Reused `padWorksheet`, `problemRanges`, and `findOperator` from part one.
- Added `parseNumbersByColumns`, which walks each problem block from right to left and reads each column top to bottom.
- Reused the evaluation logic through a shared `evaluateNumbers` helper.

### Complexity

Let `R` be the number of rows and `C` the worksheet width.

- Time: `O(R * C)`
- Space: `O(R * C)` with the padded grid

This keeps part two as a small extension of the part one parser rather than a separate code path for worksheet segmentation.

```bash
go test ./aoc2025 -run TestTrashCompactor2
```

```output
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.006s
```

## Part Two Main Program Integration

Updated `main.go` to print both `TrashCompactor(...)` and `TrashCompactor2(...)` from `aoc2025/trash_compactor.txt` when running the repository entrypoint.

For reproducibility, the command remains:

```bash
go run .
```

This walkthrough continues to omit the numeric Trash Compactor answers from the document.

