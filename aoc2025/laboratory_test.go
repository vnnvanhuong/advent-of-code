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

func TestLaboratory2(t *testing.T) {
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

		const want = 40
		if got := aoc2025.Laboratory2(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("no splitters", func(t *testing.T) {
		manifold := []string{
			"..S..",
			".....",
			".....",
		}

		const want = 1
		if got := aoc2025.Laboratory2(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single splitter", func(t *testing.T) {
		manifold := []string{
			"..S..",
			"..^..",
			".....",
		}

		const want = 2
		if got := aoc2025.Laboratory2(manifold); got != want {
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

		// Row 0: S at (3,0) -> 1 path
		// Row 1: ^ at (3,1) -> splits to (2,2) and (4,2)
		// Row 2: ^ at (2,2) and (4,2) -> splits to (1,3), (3,3) and (3,3), (5,3)
		// Row 3: . at (1,3), ^ at (3,3), . at (5,3)
		//        The 2 paths at (3,3) hit the splitter, each splitting into 2 -> 4 paths from here.
		//        The 1 path at (1,3) continues straight.
		//        The 1 path at (5,3) continues straight.
		// Total paths exiting = 1 + 4 + 1 = 6
		const want = 6
		if got := aoc2025.Laboratory2(manifold); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
