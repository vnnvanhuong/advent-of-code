package main

import (
	"fmt"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func main() {
	fmt.Println("Hello, Advent of Code 2025!")
	rotations := aoc2025.SecretEntranceInput("aoc2025/secret_entrance.txt")
	fmt.Println("Secret Entrance 1:", aoc2025.SecretEntrance1(rotations))
	fmt.Println("Secret Entrance 2:", aoc2025.SecretEntrance2(rotations))
}
