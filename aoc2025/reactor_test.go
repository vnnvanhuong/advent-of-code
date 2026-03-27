package aoc2025_test

import (
	"strings"
	"testing"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func TestReactorExample(t *testing.T) {
	const example = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	got := aoc2025.ReactorFromLines(strings.Split(example, "\n"))
	const want int64 = 5
	if got != want {
		t.Errorf("ReactorFromLines(example) = %d, want %d", got, want)
	}
}

func TestReactorNoPath(t *testing.T) {
	const noOut = `you: aaa
aaa: bbb`
	if n := aoc2025.ReactorFromLines(strings.Split(noOut, "\n")); n != 0 {
		t.Errorf("want 0 paths when out unreachable, got %d", n)
	}
}

func TestReactorDirect(t *testing.T) {
	const input = `you: out`
	if n := aoc2025.ReactorFromLines(strings.Split(input, "\n")); n != 1 {
		t.Errorf("want 1, got %d", n)
	}
}

func TestReactor2Example(t *testing.T) {
	const example = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	got := aoc2025.Reactor2FromLines(strings.Split(example, "\n"))
	const want int64 = 2
	if got != want {
		t.Errorf("Reactor2FromLines(example) = %d, want %d", got, want)
	}
}
