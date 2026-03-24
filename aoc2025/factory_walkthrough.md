# Day 10: Factory Walkthrough

*2026-03-24T07:23:43Z by Showboat dev*
<!-- showboat-id: d3739020-63ea-49bf-bf8d-8274c3b49150 -->

## Problem Summary

**The Goal:**
You need to figure out the absolute minimum number of button presses required to turn on the correct pattern of lights for a list of machines, and then add up those minimums for your final answer.

**How it Works:**
- **The Setup:** You have a list of machines (one per line). Every machine has a row of indicator lights that all start in the **OFF** (`.`) position.
- **The Target:** Each machine has a specific goal pattern of lights that need to be **ON** (`#`) or **OFF** (`.`), shown in the square brackets (e.g., `[.##.]`).
- **The Buttons:** Each machine has a set of buttons, shown in parentheses (e.g., `(0,2)`). Pressing a button "toggles" a specific group of lights (flips them from OFF to ON, or ON to OFF). The numbers represent the positions of the lights it flips (starting at 0 for the first light).
- **The Distraction:** The numbers in the curly braces at the end of each line (e.g., `{3,5,4,7}`) are "joltage requirements" and should be completely ignored for this part of the puzzle.

**The Strategy:**
Because pressing a button twice just flips the lights back to their original state, you only ever need to press any given button either `0` or `1` time. You need to find the smallest combination of buttons to press for each machine to match its target pattern, and then sum those minimum presses together.

## Logical Solution

1. **Parsing the Input:**
   - Extract the target state from the square brackets `[...]`. We can represent this as an integer where the $i$-th bit is `1` if the $i$-th character is `#`, and `0` if it's `.`.
   - Extract the button schematics from the parentheses `(...)`. Each button can also be represented as an integer (a bitmask) where the $i$-th bit is `1` if the button toggles the $i$-th light.
   - Ignore the joltage requirements in the curly braces `{...}`.

2. **Modeling the Problem:**
   - Pressing a button toggles its corresponding lights. In binary, toggling is equivalent to the bitwise XOR operation (`^`).
   - Pressing a button twice returns the lights to their previous state (since $x \oplus x = 0$). Therefore, to minimize button presses, we should press each button at most once.
   - The problem reduces to finding a subset of the available buttons such that the XOR sum of their bitmasks equals the target state bitmask, while minimizing the size of this subset.

3. **Algorithm (Brute-Force Search):**
   - For each machine, let $B$ be the number of buttons. Since $B$ is small (up to ~13 based on the input), we can iterate through all possible subsets of buttons.
   - There are $2^B$ possible subsets. We can represent each subset as an integer from $0$ to $2^B - 1$.
   - For each subset:
     - Calculate the XOR sum of the bitmasks of the included buttons.
     - If the XOR sum matches the target state, count the number of buttons in the subset (which is the number of set bits in the subset integer).
     - Keep track of the minimum number of buttons pressed across all valid subsets.
   - Sum the minimum button presses for all machines to get the final answer.

4. **Complexity:**
   - **Time Complexity:** For each machine, we check $2^B$ subsets. If there are $M$ machines, the time complexity is $O(M \cdot 2^B)$. With $M = 165$ and $B \le 13$, $165 \times 8192 \approx 1.3 \times 10^6$ operations, which will execute in milliseconds.
   - **Space Complexity:** $O(B)$ per machine to store the button masks, so overall $O(M \cdot B)$ to store the parsed input, which is extremely small.

## Dry Run

Let's dry-run the first example machine:
`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}`

**1. Parsing:**
- Target: `[.##.]` -> `0110` in binary -> `6` in decimal.
- Buttons:
  - `b0`: `(3)` -> `1000` in binary -> `8`
  - `b1`: `(1,3)` -> `1010` in binary -> `10`
  - `b2`: `(2)` -> `0100` in binary -> `4`
  - `b3`: `(2,3)` -> `1100` in binary -> `12`
  - `b4`: `(0,2)` -> `0101` in binary -> `5`
  - `b5`: `(0,1)` -> `0011` in binary -> `3`

**2. Execution:**
We are looking for a combination of these buttons that XORs to `6` (`0110`), with the minimum number of buttons.

Let's check subsets of size 1:
- `b0`: `8` != `6`
- `b1`: `10` != `6`
- `b2`: `4` != `6`
- `b3`: `12` != `6`
- `b4`: `5` != `6`
- `b5`: `3` != `6`

Let's check some subsets of size 2:
- `b4` ^ `b5`: `5 ^ 3` = `0101 ^ 0011` = `0110` = `6`.
  - Wait, let's re-verify the bit positions. The problem says "0 means the first light, 1 means the second light".
  - If we read left-to-right, `0` is the LSB or MSB?
  - Let's assume `0` is the LSB (rightmost bit).
  - Target `[.##.]`: index 0 is `.`, index 1 is `#`, index 2 is `#`, index 3 is `.`.
  - So bits are: `bit 0 = 0`, `bit 1 = 1`, `bit 2 = 1`, `bit 3 = 0`.
  - Binary representation: `0110` (which is `6`).
  - `b4`: `(0,2)` -> `bit 0 = 1`, `bit 2 = 1` -> `0101` (which is `5`).
  - `b5`: `(0,1)` -> `bit 0 = 1`, `bit 1 = 1` -> `0011` (which is `3`).
  - `b4 ^ b5` = `0101 ^ 0011` = `0110` = `6`.
  - This matches the target!

The number of buttons pressed is 2. This matches the example's expected output of 2.

The logic holds up perfectly.

## Implementation and Testing

I have implemented the solution in Go. The tests pass successfully.

The time and space complexity match our initial analysis:
- **Time Complexity:** $O(M \cdot 2^B)$ where $M$ is the number of machines and $B$ is the number of buttons per machine.
- **Space Complexity:** $O(B)$ per machine to store the button masks.

## Takeaway

The key lesson from this problem is recognizing when a problem can be modeled using bitwise operations.
- **State Representation:** Representing a series of binary states (on/off) as bits in an integer is highly efficient.
- **Toggling as XOR:** Recognizing that "toggling" a state is equivalent to the bitwise XOR (`^`) operation simplifies the logic immensely.
- **Brute-Force Feasibility:** By analyzing the constraints (the number of buttons per machine was small, $\le 13$), we determined that a brute-force approach iterating through all $2^B$ subsets was perfectly feasible and would run in milliseconds. This avoided the need for more complex algorithms like Gaussian elimination over GF(2), which could also solve this but would be overkill given the small input size.

