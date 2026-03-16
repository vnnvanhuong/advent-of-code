# Day 3: Lobby - Walkthrough

*2026-03-16T10:40:01Z by Showboat dev*
<!-- showboat-id: e5be2949-90e3-43d1-b691-2d9af52b76b3 -->

## Problem Summary

Batteries are arranged in banks (one string of digits per bank). You must turn on a fixed number of batteries per bank, maintaining their original order. The joltage output is the number formed by the selected digits concatenated.

**Part 1:** Turn on exactly **2** batteries per bank to maximize the resulting 2-digit number. Return the sum of all banks' maximum joltages.

**Part 2:** Turn on exactly **12** batteries per bank to maximize the resulting 12-digit number. Return the sum of all banks' maximum joltages.

## Worked Example (from problem statement)

Input banks:

    987654321111111
    811111111111119
    234234234234278
    818181911112111

### Part 1 -- pick 2 batteries for max 2-digit number

| Bank | Best pair (positions) | Joltage | Reasoning |
|------|-----------------------|---------|-----------|
| 987654321111111 | pos 0,1 (digits 9,8) | 98 | 9 is the largest first digit; 8 is the largest digit after it |
| 811111111111119 | pos 0,14 (digits 8,9) | 89 | 8 first, then 9 is the only digit > 1 after it |
| 234234234234278 | pos 13,14 (digits 7,8) | 78 | No digit before 7 beats 7 as first digit with an 8+ after it |
| 818181911112111 | pos 6,11 (digits 9,2) | 92 | 9 is the largest first digit; best digit after position 6 is 2 |

**Total: 98 + 89 + 78 + 92 = 357**

### Part 2 -- pick 12 batteries for max 12-digit number

The strategy is greedy: for each digit position (most significant first), pick the largest available digit that still leaves enough remaining characters.

| Bank | Drop 3 of 15 | Result | Reasoning |
|------|-------------|--------|-----------|
| 987654321111111 | drop three trailing 1s | 987654321111 | Already nearly sorted descending |
| 811111111111119 | drop three interior 1s | 811111111119 | Keep 8 at front and 9 at end |
| 234234234234278 | drop 2,3,2 near the start | 434234234278 | Greedy skips the first "2,3" to grab "4" earlier |
| 818181911112111 | drop three 1s near the front | 888911112111 | Greedy picks 8,8,8,9 from the first 7 positions |

**Total: 987654321111 + 811111111119 + 434234234278 + 888911112111 = 3,121,910,778,619**

## Additional Test Cases

### Part 1

| Case | Input | Expected | Rationale |
|------|-------|----------|-----------|
| Minimal bank | "12" | 12 | Only one possible pair |
| Empty bank | "" | 0 | No valid pair |
| Single digit | "5" | 0 | Can't form a 2-digit number |

### Part 2

| Case | Input | Expected | Rationale |
|------|-------|----------|-----------|
| Bank shorter than 12 | "12345678901" (11 chars) | 0 | Not enough batteries |
| Exactly 12 digits | "987654321012" | 987654321012 | Must use all -- no choice |
| Full example (4 banks) | see above | 3121910778619 | Sum of four 12-digit maximums |

## Part 1 -- Approach & Implementation

The solution provides **two** implementations, progressing from brute force to optimized.

### Approach 1: Brute Force (`Lobby1`) -- O(n^2)

Try all ordered pairs (i, j) where i < j. Compute d_i * 10 + d_j and keep the maximum.

```go
for i := 0; i < len(bank)-1; i++ {
    d1 := int(bank[i] - '0')
    for j := i + 1; j < len(bank); j++ {
        d2 := int(bank[j] - '0')
        val := d1*10 + d2
        if val > maxVal { maxVal = val }
    }
}
```

Simple and correct, but O(n^2) per bank.

### Approach 2: Suffix-Max Precomputation (`PrefixSum_Lobby1`) -- O(n)

Key insight: the best two-digit number starting with digit d_i is d_i * 10 + (largest digit after position i). We can precompute "largest digit after each index" in one backward pass.

```go
// Backward pass: maxAfter[i] = max digit in bank[i+1:]
maxAfter := make([]int, n)
mva := 0
for i := n - 1; i >= 0; i-- {
    maxAfter[i] = mva
    d := int(bank[i] - '0')
    if d > mva { mva = d }
}

// Forward pass: best pair for each starting position
maxVal := 0
for i := 0; i < n-1; i++ {
    d := int(bank[i] - '0')
    val := d*10 + maxAfter[i]
    if val > maxVal { maxVal = val }
}
```

Two linear passes: O(n) time, O(n) space.

### Dry-run: "818181911112111"

Backward pass builds maxAfter (max digit to the right of each position):

    Index:    0  1  2  3  4  5  6  7  8  9  10 11 12 13 14
    Digit:    8  1  8  1  8  1  9  1  1  1  1  2  1  1  1
    maxAfter: 9  9  9  9  9  9  2  2  2  2  2  1  1  1  0

Forward pass:
- i=0: 8*10 + 9 = 89
- i=2: 8*10 + 9 = 89
- i=4: 8*10 + 9 = 89
- i=6: 9*10 + 2 = **92** (winner)
- i=11: 2*10 + 1 = 21

Max = 92. Correct!

```bash
cd aoc2025 && go test -v -run TestLobby1
```

```output
=== RUN   TestLobby1
=== RUN   TestLobby1/example_banks
=== RUN   TestLobby1/single_short_bank
=== RUN   TestLobby1/edge_cases
--- PASS: TestLobby1 (0.00s)
    --- PASS: TestLobby1/example_banks (0.00s)
    --- PASS: TestLobby1/single_short_bank (0.00s)
    --- PASS: TestLobby1/edge_cases (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

## Part 2 -- Approach & Implementation

### Approach: Greedy Window Scan

Selecting k=12 digits (in order) to form the largest possible number is a classic **greedy subsequence** problem. The key insight: to maximize the leftmost (most significant) digit first, then the next, and so on.

For each of the 12 picks, we have a **window** of valid positions: from the current position `pos` to `n - picks + 1`. This window constraint ensures enough characters remain for the remaining picks. Within the window, choose the largest digit (breaking ties by taking the leftmost to preserve maximum flexibility).

### Code walkthrough (`Lobby2`)

```go
const k = 12
for _, bank := range banks {
    n := len(bank)
    if n < k { continue }

    pos := 0
    var val int64
    for picks := k; picks > 0; picks-- {
        end := n - picks + 1    // right boundary of search window
        best := -1
        bestIdx := pos
        for i := pos; i < end; i++ {
            d := int(bank[i] - '0')
            if d > best {
                best = d
                bestIdx = i
                if best == 9 { break }  // early exit optimization
            }
        }
        val = val*10 + int64(best)
        pos = bestIdx + 1    // advance past chosen digit
    }
    total += val
}
```

Notable details:

1. **`int64` accumulator:** 12-digit numbers exceed int32 range (max ~2.1 billion), so int64 is required.
2. **Early exit on 9:** Since digits are 1-9, finding a 9 means we can't do better -- skip the rest of the window. This is a significant speedup in practice.
3. **Window shrinks as we pick:** Each pick narrows the valid window because the right boundary `n - picks + 1` grows as `picks` decreases.

### Detailed dry-run: "234234234234278" (n=15, k=12)

Must drop exactly 3 of the 15 digits.

| Pick | Window [pos, end) | Digits in window | Best | bestIdx | val so far |
|------|-------------------|-----------------|------|---------|------------|
| 12 | [0, 4) | 2,3,4,2 | 4 | 2 | 4 |
| 11 | [3, 5) | 2,3 | 3 | 4 | 43 |
| 10 | [5, 6) | 4 | 4 | 5 | 434 |
| 9 | [6, 7) | 2 | 2 | 6 | 4342 |
| 8 | [7, 8) | 3 | 3 | 7 | 43423 |
| 7 | [8, 9) | 4 | 4 | 8 | 434234 |
| 6 | [9, 10) | 2 | 2 | 9 | 4342342 |
| 5 | [10, 11) | 3 | 3 | 10 | 43423423 |
| 4 | [11, 12) | 4 | 4 | 11 | 434234234 |
| 3 | [12, 13) | 2 | 2 | 12 | 4342342342 |
| 2 | [13, 14) | 7 | 7 | 13 | 43423423427 |
| 1 | [14, 15) | 8 | 8 | 14 | 434234234278 |

Result: **434234234278**. The greedy approach skipped the initial "2,3" to grab "4" as the first digit, then was forced through the remaining positions.

```bash
cd aoc2025 && go test -v -run TestLobby2
```

```output
=== RUN   TestLobby2
=== RUN   TestLobby2/example_banks
=== RUN   TestLobby2/short_banks
=== RUN   TestLobby2/exactly_twelve
--- PASS: TestLobby2 (0.00s)
    --- PASS: TestLobby2/example_banks (0.00s)
    --- PASS: TestLobby2/short_banks (0.00s)
    --- PASS: TestLobby2/exactly_twelve (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

## Complexity Summary

Let B = number of banks, n = bank length, k = number of batteries to pick.

| | Time | Space |
|---|------|-------|
| Lobby1 (brute force) | O(B * n^2) | O(1) |
| PrefixSum_Lobby1 | O(B * n) | O(n) for maxAfter array |
| Lobby2 (greedy) | O(B * n * k) | O(1) |

For Part 2, worst case is O(B * n * 12). The early-exit on digit 9 makes this much faster in practice -- if the bank contains 9s, many window scans terminate after just a few comparisons.

## Minor Code Review Notes

1. **`Lobby1` is unused:** Tests call `PrefixSum_Lobby1` exclusively. The brute-force `Lobby1` is dead code. Consider removing it or adding tests that exercise both implementations to confirm they agree.
2. **Naming convention:** `PrefixSum_Lobby1` uses an underscore, which is not idiomatic Go. Go convention would be `PrefixSumLobby1` or `Lobby1PrefixSum`.
3. **Part 2 could use a suffix-max optimization:** Similar to `PrefixSum_Lobby1`, a precomputed sparse table or suffix-max-with-index could make each greedy pick O(1) instead of scanning the window. This would reduce Part 2 from O(n*k) to O(n + k) per bank. Overkill here, but worth noting as a pattern.

## Takeaway

This problem illustrates the **greedy subsequence selection** pattern. The key lessons:

1. **Suffix-max precomputation:** For Part 1's "pick the best pair" problem, scanning all O(n^2) pairs is unnecessary. Precomputing the max digit to the right of each position reduces the problem to a single O(n) forward pass. This is a widely applicable trick whenever you need "the best element after position i."

2. **Greedy window narrowing for k-subsequence:** Part 2's algorithm is the standard greedy approach for "select k elements from a sequence to maximize the formed number." The window `[pos, n-k+1)` ensures feasibility while the greedy choice (largest digit in window) ensures optimality. This works because selecting a larger digit at a more significant position always dominates.

3. **From O(n^2) to O(n) -- recognizing structure:** The jump from `Lobby1` to `PrefixSum_Lobby1` is a textbook example of replacing nested iteration with precomputation. The suffix-max array captures all the information the inner loop was recomputing redundantly.

