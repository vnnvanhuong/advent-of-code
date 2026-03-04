package aoc2025_test

import (
	"strconv"
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

// tests for part one behavior
func TestGiftShop1(t *testing.T) {
	t.Run("small range", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{{Start: 11, End: 22}}
		actual := aoc2025.GiftShop1(ranges)
		expected := 33
		if actual != expected {
			t.Errorf("Expected %d got %d", expected, actual)
		}
	})

	t.Run("111 is not invalid (three repeats)", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{{Start: 111, End: 111}}
		actual := aoc2025.GiftShop1(ranges)
		if actual != 0 {
			t.Errorf("expected 0 for 111, got %d", actual)
		}
	})

	t.Run("full example from description", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{
			{Start: 11, End: 22},
			{Start: 95, End: 115},
			{Start: 998, End: 1012},
			{Start: 1188511880, End: 1188511890},
			{Start: 222220, End: 222224},
			{Start: 1698522, End: 1698528},
			{Start: 446443, End: 446449},
			{Start: 38593856, End: 38593862},
			{Start: 565653, End: 565659},
			{Start: 824824821, End: 824824827},
			{Start: 2121212118, End: 2121212124},
		}
		expected := 1227775554
		actual := aoc2025.GiftShop1(ranges)
		if actual != expected {
			t.Errorf("Expected %d, Got %d", expected, actual)
		}
	})
}

// tests for part two behavior
func TestGiftShop2(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{{Start: 11, End: 22}}
		actual := aoc2025.GiftShop2(ranges)
		expected := 33
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{{Start: 998, End: 1012}}
		actual := aoc2025.GiftShop2(ranges)
		expected := 999 + 1010
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("full example", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{
			{Start: 11, End: 22},
			{Start: 95, End: 115},
			{Start: 998, End: 1012},
			{Start: 1188511880, End: 1188511890},
			{Start: 222220, End: 222224},
			{Start: 1698522, End: 1698528},
			{Start: 446443, End: 446449},
			{Start: 38593856, End: 38593862},
			{Start: 565653, End: 565659},
			{Start: 824824821, End: 824824827},
			{Start: 2121212118, End: 2121212124},
		}
		expected := 4174379265
		actual := aoc2025.GiftShop2(ranges)
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("edge cases and helpers", func(t *testing.T) {
		tests := []struct {
			id      int
			invalid bool
		}{
			{11, true},     // 1 repeated twice
			{111, true},    // 1 repeated three times
			{123123, true}, // "123" twice
			{121212, true}, // "12" three times
			{1234, false},  // not a repeat
			{1213, false},  // not a pure repetition
		}

		for _, tt := range tests {
			name := strconv.Itoa(tt.id)
			t.Run(name, func(t *testing.T) {
				ranges := []aoc2025.GiftRange{{Start: tt.id, End: tt.id}}
				actual := aoc2025.GiftShop2(ranges)
				expected := 0
				if tt.invalid {
					expected = tt.id
				}
				if actual != expected {
					t.Errorf("id %d expected %d got %d", tt.id, expected, actual)
				}
			})
		}
	})
}
