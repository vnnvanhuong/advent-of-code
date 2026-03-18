package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestLaboratory(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		manifold := []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}

		const want = 21
		if got := aoc2025.Laboratory(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("no splitters", func(t *testing.T) {
		manifold := []string{
			"..S..",
			".....",
			".....",
		}

		const want = 0
		if got := aoc2025.Laboratory(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single splitter", func(t *testing.T) {
		manifold := []string{
			"..S..",
			"..^..",
			".....",
		}

		const want = 1
		if got := aoc2025.Laboratory(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("overlapping beams", func(t *testing.T) {
		manifold := []string{
			"...S...",
			"...^...",
			"..^.^..",
			"...^...",
		}

		// The first splitter splits into 2 beams.
		// These 2 beams hit the 2 splitters on the next row.
		// They each spawn 2 beams (4 total).
		// The two inner beams overlap and hit the middle splitter on the last row.
		// Total splitters hit: 1 + 2 + 1 = 4
		const want = 4
		if got := aoc2025.Laboratory(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
