# Lobby Problem Analysis

*2026-03-10T02:49:27Z by Showboat dev*
<!-- showboat-id: 95e37acd-e4d5-4fc5-ae08-a4acde89be3e -->

## Part 1
Describe puzzle setup and initial brute-force approach.

Explain character-to-int conversion with subtracting '0' and why.

Discussed optimizations: suffix max scan, greedy digit selection, linear algorithms and relevant problems.

```bash
cd aoc2025 && go test -run Lobby
```

```output
PASS
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.004s
```

## Part 2
Part two extends the problem: instead of choosing two batteries per bank, select exactly twelve batteries to maximize the 12-digit joltage value. The goal becomes finding the lexicographically largest subsequence of length twelve from each line, then summing those values across banks.

```bash
go test ./aoc2025/ -run TestLobby2
```

```output
--- FAIL: TestLobby2 (0.00s)
    --- FAIL: TestLobby2/example_banks (0.00s)
        lobby_test.go:64: expected 3121910778619, got 0
    --- FAIL: TestLobby2/exactly_twelve (0.00s)
        lobby_test.go:81: expected 987654321012 got 0
FAIL
FAIL	nguyenvanhuong.vn/adventofcode/aoc2025	0.006s
FAIL
```

Implemented Lobby2 using the greedy window scan algorithm. We now select the largest possible digit for each of the 12 positions within shrinking windows.

```bash
go test ./aoc2025/ -run TestLobby2
```

```output
ok  	nguyenvanhuong.vn/adventofcode/aoc2025	0.006s
```
