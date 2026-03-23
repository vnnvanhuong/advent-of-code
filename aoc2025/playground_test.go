package aoc2025_test

import (
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestPlayground(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		boxes := []string{
			"162,817,812",
			"57,618,57",
			"906,360,560",
			"592,479,940",
			"352,342,300",
			"466,668,158",
			"542,29,236",
			"431,825,988",
			"739,650,466",
			"52,470,668",
			"216,146,977",
			"819,987,18",
			"117,168,530",
			"805,96,715",
			"346,949,466",
			"970,615,88",
			"941,993,340",
			"862,61,35",
			"984,92,344",
			"425,690,689",
		}

		const want = 40
		if got := aoc2025.Playground(boxes, 10); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("two boxes single connection", func(t *testing.T) {
		boxes := []string{
			"0,0,0",
			"1,1,1",
		}

		// 1 connection merges both into a circuit of size 2.
		// Only 1 circuit exists; top-3 sizes are 2, 1, 1 → product = 2.
		const want = 2
		if got := aoc2025.Playground(boxes, 1); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("three boxes two connections", func(t *testing.T) {
		boxes := []string{
			"0,0,0",
			"1,0,0",
			"2,0,0",
		}

		// Pairs sorted by distance:
		//   (0,0,0)-(1,0,0) dist=1
		//   (1,0,0)-(2,0,0) dist=1
		//   (0,0,0)-(2,0,0) dist=2
		// 2 connections: both merge → single circuit of size 3.
		// Top-3 sizes: 3, 1, 1 → product = 3.
		const want = 3
		if got := aoc2025.Playground(boxes, 2); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("redundant connections still consumed", func(t *testing.T) {
		boxes := []string{
			"0,0,0",
			"1,0,0",
			"2,0,0",
			"100,100,100",
		}

		// Pairs sorted by distance:
		//   (0,0,0)-(1,0,0) dist=1        → merge {0,1}
		//   (1,0,0)-(2,0,0) dist=1        → merge {0,1,2}
		//   (0,0,0)-(2,0,0) dist=2        → redundant (same circuit)
		// After 3 connections: circuit {0,1,2} size 3, circuit {100,100,100} size 1.
		// Top-3 sizes: 3, 1, 1 → product = 3.
		// The 3rd connection was redundant but still counted.
		const want = 3
		if got := aoc2025.Playground(boxes, 3); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("zero connections", func(t *testing.T) {
		boxes := []string{
			"0,0,0",
			"1,1,1",
			"2,2,2",
		}

		// No connections made. Each box is its own circuit of size 1.
		// Top-3 sizes: 1, 1, 1 → product = 1.
		const want = 1
		if got := aoc2025.Playground(boxes, 0); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("fewer boxes than three circuits", func(t *testing.T) {
		boxes := []string{
			"0,0,0",
			"5,5,5",
		}

		// No connections. Two circuits of size 1.
		// Top-3 sizes: 1, 1, 1 (pad with 1s) → product = 1.
		const want = 1
		if got := aoc2025.Playground(boxes, 0); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}

func TestPlayground2(t *testing.T) {
	t.Run("sample from puzzle description", func(t *testing.T) {
		boxes := []string{
			"162,817,812",
			"57,618,57",
			"906,360,560",
			"592,479,940",
			"352,342,300",
			"466,668,158",
			"542,29,236",
			"431,825,988",
			"739,650,466",
			"52,470,668",
			"216,146,977",
			"819,987,18",
			"117,168,530",
			"805,96,715",
			"346,949,466",
			"970,615,88",
			"941,993,340",
			"862,61,35",
			"984,92,344",
			"425,690,689",
		}

		// Last merge connects 216,146,977 and 117,168,530.
		// Answer: 216 * 117 = 25272.
		const want = 25272
		if got := aoc2025.Playground2(boxes); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("two boxes", func(t *testing.T) {
		boxes := []string{
			"3,0,0",
			"7,0,0",
		}

		// Only one pair; merging it forms a single circuit immediately.
		// Answer: 3 * 7 = 21.
		const want = 21
		if got := aoc2025.Playground2(boxes); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("three boxes collinear", func(t *testing.T) {
		boxes := []string{
			"2,0,0",
			"5,0,0",
			"10,0,0",
		}

		// Pairs sorted by dist²:
		//   (2,0,0)-(5,0,0)  = 9   → merge → 2 components
		//   (5,0,0)-(10,0,0) = 25  → merge → 1 component  ← last merge
		// Answer: 5 * 10 = 50.
		const want = 50
		if got := aoc2025.Playground2(boxes); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("four boxes with redundant pair before final merge", func(t *testing.T) {
		boxes := []string{
			"1,0,0",
			"2,0,0",
			"3,0,0",
			"100,0,0",
		}

		// Pairs sorted by dist²:
		//   (1,0,0)-(2,0,0) = 1     → merge {0,1}       → 3 components
		//   (2,0,0)-(3,0,0) = 1     → merge {0,1,2}     → 2 components
		//   (1,0,0)-(3,0,0) = 4     → redundant          → 2 components
		//   (3,0,0)-(100,0,0) = 9409 → merge {0,1,2,3}  → 1 component ← last merge
		// Answer: 3 * 100 = 300.
		const want = 300
		if got := aoc2025.Playground2(boxes); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("single box", func(t *testing.T) {
		boxes := []string{
			"42,0,0",
		}

		// Already one circuit; no merging needed. No pair exists.
		// Return 0 as there is no connecting pair.
		const want = 0
		if got := aoc2025.Playground2(boxes); got != want {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}
