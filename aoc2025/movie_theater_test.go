package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestMovieTheater(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		input := []string{
			"7,1",
			"11,1",
			"11,7",
			"9,7",
			"9,5",
			"2,5",
			"2,3",
			"7,3",
		}

		const want = 50
		if got := aoc2025.MovieTheater(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("two points horizontal line", func(t *testing.T) {
		input := []string{
			"1,1",
			"5,1",
		}

		// Width = |1-5| + 1 = 5, Height = |1-1| + 1 = 1. Area = 5 * 1 = 5
		const want = 5
		if got := aoc2025.MovieTheater(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("two points vertical line", func(t *testing.T) {
		input := []string{
			"2,2",
			"2,6",
		}

		// Width = |2-2| + 1 = 1, Height = |2-6| + 1 = 5. Area = 1 * 5 = 5
		const want = 5
		if got := aoc2025.MovieTheater(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single point", func(t *testing.T) {
		input := []string{
			"3,3",
		}

		// Need two points to form a rectangle, so max area should be 0
		const want = 0
		if got := aoc2025.MovieTheater(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("large coordinates", func(t *testing.T) {
		input := []string{
			"0,0",
			"1000,1000",
		}

		// Width = 1001, Height = 1001. Area = 1001 * 1001 = 1002001
		const want = 1002001
		if got := aoc2025.MovieTheater(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}

func TestMovieTheater2(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		input := []string{
			"7,1",
			"11,1",
			"11,7",
			"9,7",
			"9,5",
			"2,5",
			"2,3",
			"7,3",
		}

		const want = 24
		if got := aoc2025.MovieTheater2(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("C-shaped polygon", func(t *testing.T) {
		input := []string{
			"0,0",
			"10,0",
			"10,10",
			"0,10",
			"0,8",
			"8,8",
			"8,2",
			"0,2",
		}

		// The polygon is C-shaped.
		// A rectangle from (0,0) to (10,10) has area 121, but it's not valid because its center (5,5) is outside.
		// Valid rectangles:
		// (0,0) to (10,2) -> area 11 * 3 = 33
		// (0,8) to (10,10) -> area 11 * 3 = 33
		// (8,2) to (10,8) -> area 3 * 7 = 21
		const want = 33
		if got := aoc2025.MovieTheater2(input); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
