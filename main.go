package main

import (
	"fmt"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func main() {
	fmt.Println("Hello, Advent of Code 2025!")

	rotations := aoc2025.ReadLines("aoc2025/secret_entrance.txt")
	fmt.Println("Secret Entrance 1:", aoc2025.SecretEntrance1(rotations))
	fmt.Println("Secret Entrance 2:", aoc2025.SecretEntrance2(rotations))

	ranges := aoc2025.GiftShopInput("aoc2025/gift_shop.txt")
	fmt.Println("Gift Shop 1:", aoc2025.GiftShop1(ranges))
	fmt.Println("Gift Shop 2", aoc2025.GiftShop2(ranges))

	banks := aoc2025.ReadLines("aoc2025/lobby.txt")
	fmt.Println("Lobby 1:", aoc2025.Lobby1(banks))
	fmt.Println("Lobby 2:", aoc2025.Lobby2(banks))

	grid := aoc2025.ReadLines("aoc2025/printing_department.txt")
	fmt.Println("Printing Department 1:", aoc2025.PrintingDepartment(grid))
	fmt.Println("Printing Department 2:", aoc2025.PrintingDepartment2(grid))

	cafRanges, cafIDs := aoc2025.CafeteriaInput("aoc2025/cafeteria.txt")
	fmt.Println("Cafeteria 1:", aoc2025.Cafeteria(cafRanges, cafIDs))
	fmt.Println("Cafeteria 2:", aoc2025.Cafeteria2(cafRanges))
}
