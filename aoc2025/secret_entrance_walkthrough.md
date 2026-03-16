# Day 1: Secret Entrance - Walkthrough

*2026-03-16T10:15:39Z by Showboat dev*
<!-- showboat-id: 333c5e5e-07b9-41e1-b777-985b2c266dff -->

## Problem Summary

A circular dial has 100 positions numbered 0-99. Starting at position 50, we receive a sequence of rotations like `L68` (left 68 clicks) or `R48` (right 48 clicks). The dial wraps: turning left from 0 reaches 99, turning right from 99 reaches 0.

**Part 1:** Count how many rotations leave the dial pointing at 0 *after* the rotation completes.

**Part 2:** Count *every click* that causes the dial to point at 0 -- including mid-rotation passes through 0, not just the final position.

## Worked Example (from problem statement)

Input rotations: L68, L30, R48, L5, R60, L55, L1, L99, R14, L82

### Part 1 -- count landings on 0

| Step | Rotation | Calculation | Position | Lands on 0? |
|------|----------|-------------|----------|-------------|
| start | -- | -- | 50 | -- |
| 1 | L68 | (50 - 68 + 100) mod 100 | 82 | No |
| 2 | L30 | (82 - 30) mod 100 | 52 | No |
| 3 | R48 | (52 + 48) mod 100 | **0** | **Yes** |
| 4 | L5 | (0 - 5 + 100) mod 100 | 95 | No |
| 5 | R60 | (95 + 60) mod 100 | 55 | No |
| 6 | L55 | (55 - 55) mod 100 | **0** | **Yes** |
| 7 | L1 | (0 - 1 + 100) mod 100 | 99 | No |
| 8 | L99 | (99 - 99) mod 100 | **0** | **Yes** |
| 9 | R14 | (0 + 14) mod 100 | 14 | No |
| 10 | L82 | (14 - 82 + 100) mod 100 | 32 | No |

**Part 1 answer: 3**

### Part 2 -- also count mid-rotation zero crossings

| Step | Rotation | Passes through 0 mid-rotation? | Lands on 0? | Zeros this step |
|------|----------|-------------------------------|-------------|-----------------|
| 1 | L68 | Yes (50, 49, ..., 0, 99, ..., 82) | No | 1 |
| 2 | L30 | No | No | 0 |
| 3 | R48 | No (lands exactly on 0) | Yes | 1 |
| 4 | L5 | No (starts at 0, first click is 99) | No | 0 |
| 5 | R60 | Yes (95, 96, ..., 99, 0, 1, ..., 55) | No | 1 |
| 6 | L55 | No (lands exactly on 0) | Yes | 1 |
| 7 | L1 | No (starts at 0, first click is 99) | No | 0 |
| 8 | L99 | No (lands exactly on 0) | Yes | 1 |
| 9 | R14 | No (starts at 0, first click is 1) | No | 0 |
| 10 | L82 | Yes (14, 13, ..., 0, 99, ..., 32) | No | 1 |

**Part 2 answer: 6** (3 landings + 3 mid-rotation crossings)

## Additional Test Cases

### Part 1

| Case | Input | Expected | Rationale |
|------|-------|----------|-----------|
| Single rotation landing on 0 | `L50` from 50 | 1 | Exact subtraction to 0 |
| Single rotation not on 0 | `R86` from 50 | 0 | (50+86) mod 100 = 36, not 0 |

### Part 2

| Case | Input | Expected | Rationale |
|------|-------|----------|-----------|
| Large rotation | `R1000` from 50 | 10 | 1000 clicks crosses 0 ten times (at clicks 50, 150, 250, ..., 950), final position is 50 |
| Single wrap right | `R86` from 50 | 1 | (50+86)=136, wraps past 0 once, lands at 36 |
| Example sequence | 10 rotations | 6 | 3 landings + 3 mid-rotation crossings |

## Part 1 -- Approach & Implementation

### Approach: Modular Arithmetic

Each rotation is a signed offset on a mod-100 ring. Parse direction (`L` = negative, `R` = positive), add to current position, normalize with the double-mod idiom `((p + d) % 100 + 100) % 100` (handles negative remainders in Go), and check if the result is 0.

- **Time:** O(N) -- one pass over N rotations
- **Space:** O(1) -- only a position counter and a count

### Code walkthrough (`SecretEntrance1`)

```go
p := 50                              // dial starts at 50
count := 0
for _, rotation := range rotations {
    direction := rotation[0]          // 'L' or 'R'
    distance, _ := strconv.Atoi(rotation[1:])
    if direction == 'L' {
        distance = -distance          // left = subtract
    }
    p = ((p+distance)%100 + 100) % 100  // normalize to 0-99
    if p == 0 { count++ }
}
return count
```

The double-mod `((x % 100) + 100) % 100` is the standard Go trick because Go's `%` operator preserves the sign of the dividend. For example, `(-18 % 100)` yields `-18` in Go, not `82`. Adding 100 and taking mod again maps it into the range [0, 99].

```bash
cd aoc2025 && go test -v -run TestSecretEntrance1
```

```output
=== RUN   TestSecretEntrance1
=== RUN   TestSecretEntrance1/test_case_1
=== RUN   TestSecretEntrance1/test_case_2
=== RUN   TestSecretEntrance1/test_case_3
--- PASS: TestSecretEntrance1 (0.00s)
    --- PASS: TestSecretEntrance1/test_case_1 (0.00s)
    --- PASS: TestSecretEntrance1/test_case_2 (0.00s)
    --- PASS: TestSecretEntrance1/test_case_3 (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.003s
```

## Part 2 -- Approach & Implementation

### The challenge

A naive approach would simulate each click one-by-one, but a single rotation like `R1000` from position 50 means 1000 individual clicks -- and with thousands of rotations in the real input, this could be very slow. We need O(1) math per rotation.

### Key insight: unbounded position tracking

Instead of normalizing the position after each click, keep the position **unbounded** (allow it to go beyond 0-99) after adding the full distance. The number of times the dial crosses 0 can then be computed from how far the unbounded position is from zero.

Three cases contribute to zero-crossings:

1. **Full 100-boundary crossings:** `floor(abs(unboundedPos) / 100)` counts how many complete cycles the dial travels through.

2. **Exact landing on 0:** If `unboundedPos == 0`, the dial lands exactly on zero -- add 1.

3. **Negative direction edge case:** When `unboundedPos < 0` and the previous position was *not* 0, the dial crossed from positive territory through 0 into negative -- this crossing isn't captured by the floor formula, so add 1.

The `startAtZero` flag prevents a false positive when the dial begins a rotation *at* 0 and moves left: e.g., from 0, `L5` goes 99, 98, 97, 96, 95 -- it never revisits 0, but the unbounded position (-5) is negative. The flag suppresses the spurious +1.

### Code walkthrough (`SecretEntrance2`)

```go
dialPosition := 50
count := 0
startAtZero := false

for _, rotation := range rotations {
    direction := rotation[0]
    distance, _ := strconv.Atoi(rotation[1:])
    if direction == 'L' { distance = -distance }

    dialPosition += distance                    // unbounded move

    // Case 1: full cycles through 0
    count += int(math.Floor(math.Abs(float64(dialPosition)) / 100))

    // Case 2: exact landing on 0
    if dialPosition == 0 { count++ }

    // Case 3: crossed 0 going negative (from non-zero start)
    if dialPosition < 0 && !startAtZero { count++ }

    // Normalize back to 0-99 for next iteration
    dialPosition = (dialPosition%100 + 100) % 100
    startAtZero = dialPosition == 0
}
```

### Dry-run: `R1000` from position 50

- unbounded = 50 + 1000 = 1050
- Case 1: floor(1050 / 100) = **10** (the dial passes 100, 200, ..., 1000 -- each is a wrap through 0)
- Case 2: 1050 != 0
- Case 3: 1050 > 0
- **Count = 10** (final position: 1050 mod 100 = 50)

### Potential improvement

The `math.Floor(math.Abs(float64(...)) / 100)` expression converts to float64. For very large distances this could theoretically lose precision. A purely integer alternative:

```go
abs := dialPosition
if abs < 0 { abs = -abs }
count += abs / 100    // integer division truncates toward zero
```

This is both faster and avoids any floating-point concern, though for AoC-scale inputs the difference is negligible.

```bash
cd aoc2025 && go test -v -run TestSecretEntrance2
```

```output
=== RUN   TestSecretEntrance2
=== RUN   TestSecretEntrance2/test_case_1
=== RUN   TestSecretEntrance2/test_case_2
=== RUN   TestSecretEntrance2/test_case_3
--- PASS: TestSecretEntrance2 (0.00s)
    --- PASS: TestSecretEntrance2/test_case_1 (0.00s)
    --- PASS: TestSecretEntrance2/test_case_2 (0.00s)
    --- PASS: TestSecretEntrance2/test_case_3 (0.00s)
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.003s
```

## Complexity Summary

| | Time | Space |
|---|------|-------|
| Part 1 | O(N) | O(1) |
| Part 2 | O(N) | O(1) |

Both parts process each rotation in constant time -- Part 1 with straightforward modular arithmetic, Part 2 with an O(1) zero-crossing formula that avoids simulating individual clicks.

## Takeaway

This problem illustrates the **modular arithmetic on a circular structure** pattern. The key lessons:

1. **Double-mod idiom for negative numbers:** Go's `%` preserves sign, so `((x % m) + m) % m` is the standard way to get a non-negative remainder. This is a recurring pattern in any language with truncated division.

2. **Unbounded tracking for crossing counts:** Part 2's trick of keeping the position unbounded (not normalizing until the end of each rotation) lets you compute zero-crossings with simple integer division instead of simulating each click. This transforms an O(N * D) brute-force approach (where D is the average distance per rotation) into O(N).

3. **Edge-case bookkeeping with a state flag:** The `startAtZero` boolean elegantly handles the subtle case where the dial *starts* at 0 and moves left -- preventing a false zero-crossing count. This kind of state-tracking is a common technique when math formulas almost-but-not-quite cover all cases.
