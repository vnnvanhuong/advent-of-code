package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestAccessibleCount(t *testing.T) {
	t.Run("empty grid returns zero", func(t *testing.T) {
		if got := aoc2025.PrintingDepartment(nil); got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("one roll is accessible", func(t *testing.T) {
		grid := []string{"@"}
		if got := aoc2025.PrintingDepartment(grid); got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})

	t.Run("dense 3x3 block has only corners accessible", func(t *testing.T) {
		grid := []string{
			"@@@",
			"@@@",
			"@@@",
		}
		// corners each have 3 neighbours (<4) so they count, the other five
		// rolls are surrounded by four or more neighbors.
		if got := aoc2025.PrintingDepartment(grid); got != 4 {
			t.Errorf("expected 4, got %d", got)
		}
	})

	t.Run("sample from puzzle description", func(t *testing.T) {
		grid := []string{
			"..@@.@@@@.",
			"@@@.@.@.@@",
			"@@@@@.@.@@",
			"@.@@@@..@.",
			"@@.@@@@.@@",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"@.@@@.@@@@",
			".@@@@@@@@.",
			"@.@.@@@.@.",
		}
		want := 13
		if got := aoc2025.PrintingDepartment(grid); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
