package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestSecretEntrance1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rotations := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
		expected := 3
		actual := aoc2025.SecretEntrance1(rotations)
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		rotations := []string{"L50"}
		expected := 1
		actual := aoc2025.SecretEntrance1(rotations)
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		rotations := []string{"R86"}
		expected := 0
		actual := aoc2025.SecretEntrance1(rotations)
		if actual != expected {
			t.Errorf("Expected: %d, Got: %d", expected, actual)
		}
	})
}
