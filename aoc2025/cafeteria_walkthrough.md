# Day 5: Cafeteria - Walkthrough

*2026-03-16T03:03:02Z by Showboat dev*
<!-- showboat-id: b4ecd3de-6ed9-4d04-980e-91f824433bbe -->

## Problem Summary

We are given a database of ingredient IDs split into two sections separated by a blank line:

1. **Fresh ID ranges** - each line defines an inclusive range (e.g. 3-5 means IDs 3, 4, 5 are fresh). Ranges can overlap.
2. **Available ingredient IDs** - a list of individual IDs to check.

**Goal:** Count how many of the available ingredient IDs fall within at least one of the fresh ranges.

## Worked Example

Given ranges and IDs from the problem description:

**Ranges:** 3-5, 10-14, 16-20, 12-18

**IDs to check:** 1, 5, 8, 11, 17, 32

| ID | Fresh? | Reason |
|----|--------|--------|
| 1  | No     | Not in any range |
| 5  | Yes    | Falls in 3-5 |
| 8  | No     | Not in any range |
| 11 | Yes    | Falls in 10-14 |
| 17 | Yes    | Falls in 16-20 and 12-18 |
| 32 | No     | Not in any range |

**Answer:** 3 fresh IDs.

## Additional Test Cases

### Test 1 - Boundary values
Range: 5-10. ID 5 is fresh (lower bound), ID 10 is fresh (upper bound), ID 4 is spoiled (just below), ID 11 is spoiled (just above).

### Test 2 - Overlapping ranges
Ranges: 1-5, 3-7. ID 4 is fresh (appears in both ranges but should only count once).

### Test 3 - Empty input
No ranges or no IDs should produce 0.

### Test 4 - Single-element range
Range: 7-7. ID 7 is fresh, ID 6 and 8 are spoiled.

## Approach Options

### Approach 1: Brute Force - check each ID against every range
For each available ID, iterate over all ranges and check low <= id <= high. Time: O(N * R) where N = number of IDs, R = number of ranges. Simple and likely sufficient for typical AoC input sizes.

### Approach 2: Merge intervals + Binary Search
Sort and merge overlapping ranges into non-overlapping intervals. For each ID, binary search the merged intervals. Time: O(R log R + N log R). Better for very large inputs.

### Approach 3: Build a set of all fresh IDs
Expand every range into a set of IDs, then check membership. Time: O(total range span + N). Space could be large if ranges are wide.

## TDD Red Phase

Wrote 6 test cases in cafeteria_test.go covering:
1. Sample from puzzle description (3 fresh out of 6)
2. Boundary values (lower and upper bounds of a range)
3. Overlapping ranges (ID counted only once)
4. Empty ranges (returns 0)
5. Empty IDs (returns 0)
6. Single-element range (7-7)

All tests fail because Cafeteria() is not yet defined.

## TDD Green Phase

Implemented Cafeteria() in cafeteria.go using Approach 2: Merge Intervals + Binary Search.

### Algorithm walkthrough

**Step 1 - Sort** ranges by start value so we can merge in one pass.

**Step 2 - Merge** overlapping/adjacent ranges into non-overlapping intervals:
- Walk left to right. If the current range overlaps with the last merged interval (start <= last.end + 1), extend; otherwise append new interval.
- Example: [3-5, 10-14, 12-18, 16-20] merges to [3-5, 10-20].

**Step 3 - Binary search** each ID against the merged intervals:
- Use sort.Search to find the rightmost interval whose start <= id.
- If id <= that interval's end, the ingredient is fresh.

### Complexity
- Time: O(R log R + N log R) where R = number of ranges, N = number of IDs
- Space: O(R) for the merged interval list

### Test results
All 6 tests pass:
- sample from puzzle description
- boundary values
- overlapping ranges count each ID once
- empty ranges returns zero
- empty IDs returns zero
- single-element range

## Puzzle Answer

Added CafeteriaInput() parser to cafeteria.go and wired it into main.go.

Run the program to get the answer:
go run main.go

## Part Two - Problem Summary

Ignore the available IDs list. Count the total number of unique ingredient IDs that all the fresh ranges cover.

Since ranges overlap, we must merge them first to avoid double-counting.

**Example:**
- Ranges: 3-5, 10-14, 16-20, 12-18
- After merging: [3-5], [10-20]
- Count: (5-3+1) + (20-10+1) = 3 + 11 = 14

This reuses the merge-intervals step from Part 1. Instead of binary searching IDs, we sum the size of each merged interval: end - start + 1.

## Part Two - TDD Red Phase

Wrote 6 test cases for Cafeteria2() in cafeteria_test.go:
1. Sample from puzzle description - ranges [3-5, 10-14, 16-20, 12-18] cover 14 unique IDs
2. Non-overlapping ranges - [1-3, 10-12] cover 3+3 = 6 IDs
3. Fully overlapping ranges - [1-10, 3-7] cover 10 IDs (inner range adds nothing)
4. Adjacent ranges merge - [1-5, 6-10] merge to [1-10] = 10 IDs
5. Single-element range - [7-7] covers 1 ID
6. Empty ranges returns 0

All tests fail because Cafeteria2() is not yet defined.

## Part Two - TDD Green Phase

Refactored cafeteria.go to extract a shared mergeRanges() helper used by both parts:
- Cafeteria() (Part 1): mergeRanges + binary search each ID
- Cafeteria2() (Part 2): mergeRanges + sum interval sizes (end - start + 1)

### Test results
All 12 tests pass (6 for Part 1, 6 for Part 2).

### Complexity
- Time: O(R log R) dominated by sorting
- Space: O(R) for the merged interval list

Run the program to get the answer:
go run main.go

## Takeaway

This problem is a classic **interval merging** pattern. The key lesson: once you have sorted, non-overlapping intervals, many questions become trivial -- membership checks via binary search (Part 1) and total coverage via summation (Part 2). Recognizing this pattern early lets you solve both parts with a single shared O(R log R) preprocessing step.
