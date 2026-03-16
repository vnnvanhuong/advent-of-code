package aoc2025

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// CafeteriaInput parses the puzzle input into ranges and IDs.
// The file has two sections separated by a blank line:
//   - Lines before the blank: "low-high" inclusive ranges
//   - Lines after the blank: individual ingredient IDs
func CafeteriaInput(filename string) ([][2]int, []int) {
	lines := ReadLines(filename)

	var ranges [][2]int
	var ids []int
	pastBlank := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			pastBlank = true
			continue
		}
		if !pastBlank {
			parts := strings.SplitN(trimmed, "-", 2)
			lo, err1 := strconv.Atoi(parts[0])
			hi, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				panic(fmt.Sprintf("bad range line: %q", line))
			}
			ranges = append(ranges, [2]int{lo, hi})
		} else {
			id, err := strconv.Atoi(trimmed)
			if err != nil {
				panic(fmt.Sprintf("bad id line: %q", line))
			}
			ids = append(ids, id)
		}
	}

	return ranges, ids
}

// mergeRanges sorts and merges overlapping/adjacent inclusive ranges into
// a minimal set of non-overlapping intervals. This is the shared core for
// both Part 1 and Part 2.
//
// Algorithm:
//
// Step 1: Sort the ranges by their start value.
//
//	Input ranges may overlap or be unordered. Sorting by start lets us
//	merge them in a single left-to-right pass.
//
//	Example: [16-20, 3-5, 12-18, 10-14]
//	     →   [3-5, 10-14, 12-18, 16-20]   (sorted)
//
// Step 2: Merge overlapping/adjacent ranges into non-overlapping intervals.
//
//	Walk through the sorted list. If the current range overlaps with or
//	touches the last merged interval (current.start <= last.end + 1),
//	extend the last interval's end. Otherwise append a new interval.
//
//	[3-5, 10-14, 12-18, 16-20]
//	 → [3-5]                          start with first range
//	 → [3-5, 10-14]                   10 > 5+1, new interval
//	 → [3-5, 10-18]                   12 <= 14+1, extend end to max(14,18)=18
//	 → [3-5, 10-20]                   16 <= 18+1, extend end to max(18,20)=20
//
// Time:  O(R log R)
// Space: O(R)
func mergeRanges(ranges [][2]int) [][2]int {
	if len(ranges) == 0 {
		return nil
	}

	sorted := make([][2]int, len(ranges))
	copy(sorted, ranges)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	// merged holds non-overlapping intervals in ascending order.
	//
	// Invariant: for all i < len(merged)-1,
	//   merged[i][1] < merged[i+1][0]   (no overlap, no adjacency)
	merged := [][2]int{sorted[0]}
	for _, r := range sorted[1:] {
		last := &merged[len(merged)-1]
		if r[0] <= last[1]+1 {
			// Overlapping or adjacent — extend the current interval.
			//   last:  [==========]
			//   r:          [===========]
			//   result:[=================]
			if r[1] > last[1] {
				last[1] = r[1]
			}
		} else {
			// Gap between last and r — start a new interval.
			//   last: [====]
			//   r:              [====]
			merged = append(merged, r)
		}
	}

	return merged
}

// Cafeteria counts how many of the available ingredient IDs are "fresh",
// meaning they fall within at least one of the given inclusive ranges.
//
// Uses mergeRanges to build non-overlapping intervals, then binary searches
// each ID against them.
//
// Time:  O(R log R + N log R)  where R = ranges, N = IDs
// Space: O(R) for the merged intervals
func Cafeteria(ranges [][2]int, ids []int) int {
	if len(ranges) == 0 || len(ids) == 0 {
		return 0
	}

	merged := mergeRanges(ranges)

	// Binary search each ID against merged intervals.
	//
	// For a given id, we want the rightmost interval whose start <= id.
	// If that interval's end >= id, the id is fresh.
	//
	// Example with merged = [3-5, 10-20]:
	//   id=1  → no interval starts <= 1          → spoiled
	//   id=5  → interval [3-5], 5 <= 5           → fresh
	//   id=8  → interval [3-5], 8 > 5            → spoiled
	//   id=11 → interval [10-20], 11 <= 20       → fresh
	count := 0
	for _, id := range ids {
		// sort.Search finds the smallest index i where merged[i][0] > id.
		// The candidate interval is therefore at index i-1.
		i := sort.Search(len(merged), func(i int) bool {
			return merged[i][0] > id
		})
		if i > 0 && id <= merged[i-1][1] {
			count++
		}
	}

	return count
}

// Cafeteria2 counts the total number of unique ingredient IDs that all
// fresh ranges cover. The available IDs list is irrelevant for Part 2.
//
// Uses mergeRanges to collapse overlapping ranges, then sums the size of
// each non-overlapping interval: end - start + 1.
//
// Example with merged = [3-5, 10-20]:
//
//	(5-3+1) + (20-10+1) = 3 + 11 = 14
//
// Time:  O(R log R)   — dominated by the sort inside mergeRanges
// Space: O(R)
func Cafeteria2(ranges [][2]int) int {
	merged := mergeRanges(ranges)

	total := 0
	for _, iv := range merged {
		total += iv[1] - iv[0] + 1
	}
	return total
}
