package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestCafeteria(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		ranges := [][2]int{{3, 5}, {10, 14}, {16, 20}, {12, 18}}
		ids := []int{1, 5, 8, 11, 17, 32}
		want := 3
		if got := aoc2025.Cafeteria(ranges, ids); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("boundary values", func(t *testing.T) {
		ranges := [][2]int{{5, 10}}
		ids := []int{4, 5, 10, 11}
		want := 2
		if got := aoc2025.Cafeteria(ranges, ids); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("overlapping ranges count each ID once", func(t *testing.T) {
		ranges := [][2]int{{1, 5}, {3, 7}}
		ids := []int{4}
		want := 1
		if got := aoc2025.Cafeteria(ranges, ids); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("empty ranges returns zero", func(t *testing.T) {
		ids := []int{1, 2, 3}
		want := 0
		if got := aoc2025.Cafeteria(nil, ids); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("empty IDs returns zero", func(t *testing.T) {
		ranges := [][2]int{{1, 10}}
		want := 0
		if got := aoc2025.Cafeteria(ranges, nil); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single-element range", func(t *testing.T) {
		ranges := [][2]int{{7, 7}}
		ids := []int{6, 7, 8}
		want := 1
		if got := aoc2025.Cafeteria(ranges, ids); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}

func TestCafeteria2(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		ranges := [][2]int{{3, 5}, {10, 14}, {16, 20}, {12, 18}}
		want := 14
		if got := aoc2025.Cafeteria2(ranges); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("non-overlapping ranges", func(t *testing.T) {
		ranges := [][2]int{{1, 3}, {10, 12}}
		want := 6
		if got := aoc2025.Cafeteria2(ranges); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("fully overlapping ranges", func(t *testing.T) {
		ranges := [][2]int{{1, 10}, {3, 7}}
		want := 10
		if got := aoc2025.Cafeteria2(ranges); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("adjacent ranges merge", func(t *testing.T) {
		ranges := [][2]int{{1, 5}, {6, 10}}
		want := 10
		if got := aoc2025.Cafeteria2(ranges); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single-element range", func(t *testing.T) {
		ranges := [][2]int{{7, 7}}
		want := 1
		if got := aoc2025.Cafeteria2(ranges); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("empty ranges returns zero", func(t *testing.T) {
		want := 0
		if got := aoc2025.Cafeteria2(nil); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
