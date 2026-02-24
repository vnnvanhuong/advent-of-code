package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestGiftShop1(t *testing.T) {

	t.Run("test case 1", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{
			{Start: 11, End: 22},
		}
		actual := aoc2025.GiftShop1(ranges)
		expected := 33

		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		ranges := []aoc2025.GiftRange{
			{Start: 998, End: 1012},
		}

		actual := aoc2025.GiftShop1(ranges)
		expected := 999 + 1010
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
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
		actual := aoc2025.GiftShop1(ranges)
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})
}
