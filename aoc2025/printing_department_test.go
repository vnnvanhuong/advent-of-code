package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestPrintingDepartment1(t *testing.T) {
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

func TestPrintingDepartment2(t *testing.T) {
	t.Run("empty grid returns zero", func(t *testing.T) {
		if got := aoc2025.PrintingDepartment2(nil); got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("single roll removed in one round", func(t *testing.T) {
		grid := []string{"@"}
		if got := aoc2025.PrintingDepartment2(grid); got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})

	t.Run("sparse row all removed at once", func(t *testing.T) {
		grid := []string{"@.@.@"}
		// each roll has 0 neighbours, all removed in round 1
		if got := aoc2025.PrintingDepartment2(grid); got != 3 {
			t.Errorf("expected 3, got %d", got)
		}
	})

	t.Run("3x3 block fully eroded in two rounds", func(t *testing.T) {
		grid := []string{
			"@@@",
			"@@@",
			"@@@",
		}
		// round 1: 4 corners removed (3 neighbours each)
		// round 2: remaining 5 cells — each now has <=3 neighbours, all removed
		// total: 9
		if got := aoc2025.PrintingDepartment2(grid); got != 9 {
			t.Errorf("expected 9, got %d", got)
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
		want := 43
		if got := aoc2025.PrintingDepartment2(grid); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
