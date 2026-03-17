package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestTrashCompactor(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		worksheet := []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
			"*   +   *   + ",
		}

		const want = 4277556
		if got := aoc2025.TrashCompactor(worksheet); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single addition problem", func(t *testing.T) {
		worksheet := []string{
			"12",
			" 3",
			"+ ",
		}

		const want = 15
		if got := aoc2025.TrashCompactor(worksheet); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single multiplication problem", func(t *testing.T) {
		worksheet := []string{
			" 12",
			"3  ",
			" 4 ",
			" * ",
		}

		const want = 144
		if got := aoc2025.TrashCompactor(worksheet); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("multiple narrow problems separated by one blank column", func(t *testing.T) {
		worksheet := []string{
			"1 2",
			"2 3",
			"+ *",
		}

		const want = 9
		if got := aoc2025.TrashCompactor(worksheet); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("alignment inside each problem is ignored", func(t *testing.T) {
		worksheet := []string{
			"12  7",
			" 3 89",
			"+  +",
		}

		const want = 111
		if got := aoc2025.TrashCompactor(worksheet); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
