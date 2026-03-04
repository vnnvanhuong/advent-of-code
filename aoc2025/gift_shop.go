package aoc2025

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type GiftRange struct {
	Start int
	End   int
}

// GiftShop1 calculates the total invalid IDs from a list of gift ranges
// GiftShop1 calculates the total invalid IDs from a list of gift ranges
// using the Part One rule: an ID is invalid if it consists of some sequence
// of digits repeated **exactly twice** (e.g. 55, 6464, 123123).
func GiftShop1(ranges []GiftRange) int {
	totalInvalid := 0
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalidPart1(id) {
				totalInvalid += id
			}
		}
	}
	return totalInvalid
}

// GiftShop2 implements the Part Two rule: invalid IDs are made of some sequence
// of digits repeated at least twice (so any number of repetitions ≥2 counts).
func GiftShop2(ranges []GiftRange) int {
	totalInvalid := 0
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalidPart2(id) {
				totalInvalid += id
			}
		}
	}
	return totalInvalid
}

// isInvalidPart1 checks for exactly two repetitions
func isInvalidPart1(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)
	// only lengths that are exactly twice a substring
	for l := 1; l <= n/2; l++ {
		if n == 2*l && strings.Repeat(s[:l], 2) == s {
			return true
		}
	}
	return false
}

// isInvalidPart2 checks for any number of repetitions >= 2.
func isInvalidPart2(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l == 0 && strings.Repeat(s[:l], n/l) == s {
			return true
		}
	}
	return false
}

func parseGiftRange(raw string) GiftRange {
	parts := strings.Split(raw, "-")
	return GiftRange{
		Start: parseInt(parts[0]),
		End:   parseInt(parts[1]),
	}
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// GiftShopInput reads and parses the input file
// file contains only one line with a list of gift ranges separated by commas
func GiftShopInput(filename string) []GiftRange {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	lines := strings.Split(string(fileContent), ",")
	ranges := []GiftRange{}
	for _, line := range lines {
		ranges = append(ranges, parseGiftRange(line))
	}

	return ranges
}
