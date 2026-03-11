package aoc2025_test

import (
	"os"
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestReadLines(t *testing.T) {
	// create a temporary file with known contents
	tmp, err := os.CreateTemp("", "lines-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	content := "first\nsecond\nthird\n"
	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	got := aoc2025.ReadLines(tmp.Name())
	want := []string{"first", "second", "third"}
	if len(got) != len(want) {
		t.Fatalf("expected %d lines, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}
