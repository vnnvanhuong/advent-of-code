# Day 11: Reactor Walkthrough

## Problem Summary

**The Goal:** Count how many **different directed routes** data can take from the device labeled `you` to the device labeled `out`, following only forward edges (each line lists outgoing connections).

**How it Works:** The input is a **directed graph** in adjacency-list form. Each path is a sequence of edges; when the graph has no cycles on relevant routes, the total count is finite and can be aggregated with dynamic programming.

## Logical Solution

1. **Parse:** For each non-empty line, split on the first `:`. The left token is the device name; the right side is whitespace-separated successor names. Build `map[name][]successor`.

2. **Count paths:** Let `P(v)` be the number of paths from `v` to `out`.
   - Base: `P(out) = 1` (one trivial path at the sink; we stop when we hit `out`).
   - Recurrence: `P(v) = sum over w in successors(v) of P(w)`.
   - Answer: `P(you)`.

3. **Implementation (memoized DFS):**
   - Pseudocode:

```
function count_from(v):
    if v == "out":
        return 1
    if v is on recursion stack:
        error "cycle"
    if v in memo:
        return memo[v]
    mark v on stack
    total = 0
    for each w in graph[v]:
        total += count_from(w)
    unmark v from stack
    memo[v] = total
    return total
```

4. **Complexity:** Time `O(V + E)`, space `O(V + E)` for the graph plus `O(V)` for memo and stack flags.

## Dry Run (example from puzzle)

- From `out`: return 1 (not stored in memo in our implementation; callers add contributions).
- Leaves like `eee`, `fff`, `ggg` each have only `out` → each contributes 1 from that edge; `P(eee)=P(fff)=P(ggg)=1`.
- `ddd` → `ggg` → 1 path; `P(ddd)=1`.
- `bbb` → `ddd` and `eee` → `P(bbb)=1+1=2`.
- `ccc` → `ddd`, `eee`, `fff` → `P(ccc)=3`.
- `you` → `bbb`, `ccc` → `P(you)=2+3=5`. Matches the statement.

## Implementation and Testing

Solution lives in `reactor.go` with tests in `reactor_test.go` (example → 5, no sink → 0, `you: out` → 1).

**Part one answer (puzzle input): [REDACTED]**

## Optimization Notes

Memoization already gives linear time in the size of the graph. Topological sort + bottom-up DP would also be `O(V + E)` and avoids recursion depth if that were a concern (not needed here).

## Takeaway

- **Path counting in a DAG:** Number of paths from a source to a sink equals the sum over outgoing edges of path counts from each neighbor—classic DP on a DAG.
- **Converging paths:** Shared subgraphs (diamonds) are handled automatically by summing at the merge point via memoization.
- **Cycles:** If simple paths were required, the problem would differ; here the puzzle implies a finite count, so treating a back-edge as an error (or detecting infinite families) keeps the model honest.

## Part Two

Not yet specified in the puzzle text.
