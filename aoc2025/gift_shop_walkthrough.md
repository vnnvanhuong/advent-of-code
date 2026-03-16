# Day 2: Gift Shop - Walkthrough

*2026-03-16T10:30:01Z by Showboat dev*
<!-- showboat-id: 5b8b3c2b-84e9-4bde-b846-1de16e79c931 -->

## Problem Summary

Given a list of inclusive numeric ranges (e.g. `11-22, 998-1012`), identify "invalid" product IDs within those ranges based on digit-repetition patterns.

**Part 1:** An ID is invalid if its digits consist of some substring repeated **exactly twice** (e.g. `55` = "5" x2, `6464` = "64" x2, `123123` = "123" x2). Sum all invalid IDs across all ranges.

**Part 2:** An ID is invalid if its digits consist of some substring repeated **at least twice** (e.g. `111` = "1" x3, `1212121212` = "12" x5 also count now). Sum all invalid IDs across all ranges.

## Worked Example (from problem statement)

### Part 1 -- exactly-twice repetitions

| Range | Invalid IDs | Reason |
|-------|-------------|--------|
| 11-22 | 11, 22 | "1"x2, "2"x2 |
| 95-115 | 99 | "9"x2 |
| 998-1012 | 1010 | "10"x2 |
| 1188511880-1188511890 | 1188511885 | "11885"x2 |
| 222220-222224 | 222222 | "222"x2 |
| 1698522-1698528 | (none) | no doubled pattern |
| 446443-446449 | 446446 | "446"x2 |
| 38593856-38593862 | 38593859 | "3859"x2 (wait -- "38593859" is "3859"x2? Let's check: len=8, half="3859", "3859"+"3859"="38593859" -- yes) |
| 565653-565659 | (none for Part 1) | 565656 = "56"x3, not exactly twice |
| 824824821-824824827 | (none for Part 1) | 824824824 = "824"x3, not exactly twice |
| 2121212118-2121212124 | (none for Part 1) | 2121212121 = "21"x5, not exactly twice |

**Sum = 11 + 22 + 99 + 1010 + 1188511885 + 222222 + 446446 + 38593859 = 1,227,775,554**

### Part 2 -- at-least-twice repetitions (new invalid IDs highlighted)

| Range | New Invalid IDs in Part 2 | Pattern |
|-------|--------------------------|---------|
| 95-115 | **111** | "1"x3 |
| 998-1012 | **999** | "9"x3 |
| 565653-565659 | **565656** | "56"x3 |
| 824824821-824824827 | **824824824** | "824"x3 |
| 2121212118-2121212124 | **2121212121** | "21"x5 |

**Part 2 sum = 1,227,775,554 + 111 + 999 + 565656 + 824824824 + 2121212121 = 4,174,379,265**

## Additional Test Cases

### Part 1

| Case | Input | Expected | Rationale |
|------|-------|----------|-----------|
| Small range 11-22 | `{11, 22}` | 33 (11+22) | Both endpoints are doubled single digits |
| 111 is NOT invalid | `{111, 111}` | 0 | "1"x3 is three repeats, not exactly two |

### Part 2

| Case | Input | Expected | Rationale |
|------|-------|----------|-----------|
| Small range 11-22 | `{11, 22}` | 33 | Same as Part 1 (two repeats qualifies) |
| 998-1012 | `{998, 1012}` | 2009 (999+1010) | 999="9"x3 now counts; 1010="10"x2 still counts |
| 111 IS invalid | `{111, 111}` | 111 | "1"x3 now qualifies |
| 123123 | `{123123, 123123}` | 123123 | "123"x2 |
| 121212 | `{121212, 121212}` | 121212 | "12"x3 |
| 1234 is valid | `{1234, 1234}` | 0 | No repeating pattern |
| 1213 is valid | `{1213, 1213}` | 0 | Looks close but not a pure repetition |

## Part 1 -- Approach & Implementation

### Approach: Brute-Force with String-Repeat Check

For each ID in each range, convert to a string and try every possible "half-length" prefix. If the string length is exactly double the prefix length and the prefix repeated twice equals the full string, the ID is invalid.

This works because AoC ranges are small (typically ~10 IDs each), so iterating every ID is fast.

### Code walkthrough

**Outer loop** (`GiftShop1`):

```go
for _, r := range ranges {
    for id := r.Start; id <= r.End; id++ {
        if isInvalidPart1(id) {
            totalInvalid += id
        }
    }
}
```

**Core check** (`isInvalidPart1`):

```go
func isInvalidPart1(id int) bool {
    s := strconv.Itoa(id)
    n := len(s)
    for l := 1; l <= n/2; l++ {
        if n == 2*l && strings.Repeat(s[:l], 2) == s {
            return true
        }
    }
    return false
}
```

The key constraint is `n == 2*l`: this restricts matches to **exactly** two repetitions. For example, `111` (n=3) never satisfies `3 == 2*l` for any integer l, so it correctly returns false for Part 1.

Note that the loop `for l := 1; l <= n/2` actually only has one value of `l` that can satisfy `n == 2*l`, namely `l = n/2`. So the loop could be simplified to a single check:

```go
if n%2 == 0 {
    half := n / 2
    return strings.Repeat(s[:half], 2) == s
}
return false
```

But the current form is harmless since the loop body short-circuits on the only matching `l`.

```bash
cd aoc2025 && go test -v -run TestGiftShop1
```

```output
=== RUN   TestGiftShop1
=== RUN   TestGiftShop1/small_range
=== RUN   TestGiftShop1/111_is_not_invalid_(three_repeats)
=== RUN   TestGiftShop1/full_example_from_description
--- PASS: TestGiftShop1 (0.00s)
    --- PASS: TestGiftShop1/small_range (0.00s)
    --- PASS: TestGiftShop1/111_is_not_invalid_(three_repeats) (0.00s)
    --- PASS: TestGiftShop1/full_example_from_description (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

## Part 2 -- Approach & Implementation

### What changed

Part 1 required exactly two repetitions (string length must be 2x the prefix). Part 2 relaxes this to *any* number of repetitions >= 2.

### Code walkthrough (`isInvalidPart2`)

```go
func isInvalidPart2(id int) bool {
    s := strconv.Itoa(id)
    n := len(s)
    for l := 1; l <= n/2; l++ {
        if n%l == 0 && strings.Repeat(s[:l], n/l) == s {
            return true
        }
    }
    return false
}
```

The two changes from Part 1:

1. **`n%l == 0`** replaces `n == 2*l` -- now any prefix length that evenly divides the total length is a candidate.
2. **`strings.Repeat(s[:l], n/l)`** repeats the prefix `n/l` times instead of a fixed 2.

The bound `l <= n/2` guarantees `n/l >= 2`, so single-copy (the number itself) is never accepted.

### Dry-run examples

**`111`** (n=3):
- l=1: 3%1==0, "1" repeated 3 times = "111" == "111" -- **invalid**

**`1212121212`** (n=10):
- l=1: 10%1==0, "1" x 10 = "1111111111" != "1212121212"
- l=2: 10%2==0, "12" x 5 = "1212121212" == "1212121212" -- **invalid**

**`1234`** (n=4):
- l=1: "1" x 4 = "1111" != "1234"
- l=2: "12" x 2 = "1212" != "1234"
- No match -- **valid**

### Connection to string periodicity

This check is equivalent to asking: does the string have a period that divides its length, with the period being at most half the length? This is related to the KMP failure function, which can find the minimal period of a string in O(n). However, for the short strings in this problem (at most ~10 digits), the brute-force `strings.Repeat` approach is perfectly adequate.

```bash
cd aoc2025 && go test -v -run TestGiftShop2
```

```output
=== RUN   TestGiftShop2
=== RUN   TestGiftShop2/test_case_1
=== RUN   TestGiftShop2/test_case_2
=== RUN   TestGiftShop2/full_example
=== RUN   TestGiftShop2/edge_cases_and_helpers
=== RUN   TestGiftShop2/edge_cases_and_helpers/11
=== RUN   TestGiftShop2/edge_cases_and_helpers/111
=== RUN   TestGiftShop2/edge_cases_and_helpers/123123
=== RUN   TestGiftShop2/edge_cases_and_helpers/121212
=== RUN   TestGiftShop2/edge_cases_and_helpers/1234
=== RUN   TestGiftShop2/edge_cases_and_helpers/1213
--- PASS: TestGiftShop2 (0.00s)
    --- PASS: TestGiftShop2/test_case_1 (0.00s)
    --- PASS: TestGiftShop2/test_case_2 (0.00s)
    --- PASS: TestGiftShop2/full_example (0.00s)
    --- PASS: TestGiftShop2/edge_cases_and_helpers (0.00s)
        --- PASS: TestGiftShop2/edge_cases_and_helpers/11 (0.00s)
        --- PASS: TestGiftShop2/edge_cases_and_helpers/111 (0.00s)
        --- PASS: TestGiftShop2/edge_cases_and_helpers/123123 (0.00s)
        --- PASS: TestGiftShop2/edge_cases_and_helpers/121212 (0.00s)
        --- PASS: TestGiftShop2/edge_cases_and_helpers/1234 (0.00s)
        --- PASS: TestGiftShop2/edge_cases_and_helpers/1213 (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.003s
```

## Complexity Summary

Let R = number of ranges, S = average range size, D = max digit count of any ID.

| | Time | Space |
|---|------|-------|
| Part 1 | O(R * S * D) | O(D) for string conversion |
| Part 2 | O(R * S * D^2) | O(D) for string conversion |

Per ID, Part 1 does at most one string comparison (only `l = n/2` can match). Part 2 tries up to D/2 prefix lengths, each requiring an O(D) string comparison, giving O(D^2) per ID.

Since AoC ranges are small (~10 IDs each) and IDs are at most ~10 digits, this is effectively O(R) in practice.

## Optimization: Generate-and-Filter

Instead of checking every ID in every range, we could **generate all invalid IDs** up to a maximum digit length and binary-search which ones fall in ranges:

1. For each prefix length `l` (1, 2, 3, ...), generate all numbers formed by repeating an `l`-digit prefix 2 times (Part 1) or 2..k times (Part 2).
2. For Part 1 with `l` digits, there are 9 * 10^(l-1) candidates (no leading zeros). For 10-digit IDs, that's ~450K total candidates across all lengths -- far fewer than iterating huge ranges.
3. Sort generated IDs, then for each range do a binary search to find and sum the ones inside.

This would be dramatically faster for ranges spanning millions of IDs, though it's overkill for the typical AoC input.

## Minor Code Review Notes

1. **Duplicate comment** on line 14-16 of `gift_shop.go`: `GiftShop1` has its doc comment repeated twice. The second copy can be removed.
2. **`isInvalidPart1` loop simplification**: The loop tries all `l` from 1 to n/2, but only `l == n/2` can satisfy `n == 2*l`. A direct `if n%2 == 0 && repeat(s[:n/2], 2) == s` would be clearer.
3. **Shared structure**: Both `GiftShop1` and `GiftShop2` have identical outer loops; they differ only in the predicate. A higher-order helper `sumInvalid(ranges, predicate)` could reduce duplication.

## Takeaway

This problem is a **string periodicity / pattern repetition** problem. The key lessons:

1. **Reduction to string operations:** Numeric patterns ("digits repeated") become simple string operations. Converting the ID to a string and using `strings.Repeat` is both clear and efficient for small inputs.

2. **Exactly-N vs at-least-N repetition:** The jump from Part 1 to Part 2 is a classic AoC escalation. Part 1's `n == 2*l` becomes Part 2's `n%l == 0` -- a minimal code change with significant behavioral impact. Recognizing that "exactly twice" is a special case of "evenly divides" is the key insight.

3. **Brute force is fine when ranges are small:** AoC often has inputs where brute force works despite seeming scary. The ranges here are tiny (~10 IDs), so the O(R * S * D^2) approach runs in microseconds. The generate-and-filter optimization exists but isn't needed.

