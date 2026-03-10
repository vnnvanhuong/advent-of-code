package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestLobby1(t *testing.T) {
	t.Run("example banks", func(t *testing.T) {
		banks := []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}
		expected := 357
		actual := aoc2025.PrefixSum_Lobby1(banks)
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("single short bank", func(t *testing.T) {
		banks := []string{"12"} // only two batteries, output 12
		expected := 12
		actual := aoc2025.PrefixSum_Lobby1(banks)
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("edge cases", func(t *testing.T) {
		// empty and single-digit banks should contribute 0
		tests := []struct {
			banks    []string
			expected int
		}{
			{[]string{""}, 0},
			{[]string{"5"}, 0},
		}
		for _, tt := range tests {
			actual := aoc2025.PrefixSum_Lobby1(tt.banks)
			if actual != tt.expected {
				t.Errorf("banks %v expected %d got %d", tt.banks, tt.expected, actual)
			}
		}
	})
}

func TestLobby2(t *testing.T) {
	t.Run("placeholder", func(t *testing.T) {
		// behavior for part two not yet specified
		banks := []string{"123"}
		_ = aoc2025.Lobby2(banks)
	})
}
