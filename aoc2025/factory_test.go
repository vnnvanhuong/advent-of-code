package aoc2025_test

import (
	"strings"
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestFactoryPartOne(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Machine 1",
			input:    "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			expected: 2,
		},
		{
			name:     "Machine 2",
			input:    "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
			expected: 3,
		},
		{
			name:     "Machine 3",
			input:    "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			expected: 2,
		},
		{
			name: "All Example Machines",
			input: `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := aoc2025.Factory(strings.Split(tt.input, "\n"))
			if result != tt.expected {
				t.Errorf("Factory() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFactoryPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Machine 1 - joltage",
			input:    "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			expected: 10,
		},
		{
			name:     "Machine 2 - joltage",
			input:    "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
			expected: 12,
		},
		{
			name:     "Machine 3 - joltage",
			input:    "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			expected: 11,
		},
		{
			name: "All Example Machines - joltage",
			input: `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`,
			expected: 33,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := aoc2025.Factory2(strings.Split(tt.input, "\n"))
			if result != tt.expected {
				t.Errorf("Factory2() = %v, want %v", result, tt.expected)
			}
		})
	}
}
