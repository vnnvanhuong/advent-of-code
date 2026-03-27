package aoc2025_test

import (
	"strings"
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestChristmasTreeFarm1Example(t *testing.T) {
	const example = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`
	lines := strings.Split(example, "\n")
	got := aoc2025.ChristmasTreeFarm(lines)
	if got != 2 {
		t.Fatalf("ChristmasTreeFarm(example) = %d, want 2", got)
	}
}
