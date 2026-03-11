package aoc2025

// no imports required; all file I/O has been moved to input.go

// Lobby1 computes the total output joltage for part one of the lobby puzzle.
// Each string in banks represents a line of digit-labeled batteries. Exactly two
// batteries must be turned on per bank and the output is the two-digit number
// formed by their labels in order. The function returns the sum of each bank's
// maximum possible joltage.
func Lobby1(banks []string) int {
	// brute-force two-pointer search per bank
	// for each line of digits we scan all ordered pairs and keep the largest
	// two-digit value encountered.
	total := 0
	for _, bank := range banks {
		maxVal := 0
		// nested loops over indices
		for i := 0; i < len(bank)-1; i++ {
			d1 := int(bank[i] - '0')
			for j := i + 1; j < len(bank); j++ {
				d2 := int(bank[j] - '0')
				val := d1*10 + d2
				if val > maxVal {
					maxVal = val
				}
			}
		}
		total += maxVal
	}
	return total
}

// PrefixSum_Lobby1 is an O(n) alternative to Lobby1. It precalculates,
// for each index i, the largest digit appearing *after* i, then uses that to
// compute the best two‑digit value for each possible first battery.  By
// restricting the search to i < len(bank)-1 we avoid invalid pairs (there is
// no second battery after the last position).
func PrefixSum_Lobby1(banks []string) int {
	total := 0
	for _, bank := range banks {
		n := len(bank)
		if n < 2 {
			// no valid pair
			continue
		}

		// slice of maximum digit to the right of each index
		maxAfter := make([]int, n)
		mva := 0
		// scan backward; maxAfter[i] should be max digit in bank[i+1:]
		for i := n - 1; i >= 0; i-- {
			maxAfter[i] = mva
			d := int(bank[i] - '0')
			if d > mva {
				mva = d
			}
		}

		maxVal := 0
		for i := 0; i < n-1; i++ {
			d := int(bank[i] - '0')
			val := d*10 + maxAfter[i]
			if val > maxVal {
				maxVal = val
			}
		}
		total += maxVal
	}
	return total
}

// Lobby2 choose exactly twelve batteries per bank to
// maximize the resulting 12-digit joltage number.  If a bank has fewer than
// twelve batteries it contributes 0.
//
// The algorithm is a simple greedy window scan: for each of the twelve
// positions we pick the largest available digit that still leaves enough
// characters to finish the sequence.
func Lobby2(banks []string) int64 {
	total := int64(0)
	const k = 12
	for _, bank := range banks {
		n := len(bank)
		if n < k {
			continue
		}
		pos := 0
		var val int64
		for picks := k; picks > 0; picks-- {
			// we must pick one digit from bank[pos : n-picks+1]
			end := n - picks + 1
			best := -1
			bestIdx := pos
			for i := pos; i < end; i++ {
				d := int(bank[i] - '0')
				if d > best {
					best = d
					bestIdx = i
					// early exit: cannot do better than 9
					if best == 9 {
						break
					}
				}
			}
			val = val*10 + int64(best)
			pos = bestIdx + 1
		}
		total += val
	}
	return total
}
