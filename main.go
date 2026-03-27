package main

import (
	"fmt"

	"nguyenvanhuong.vn/adventofcode/aoc2025"
)

func main() {
	fmt.Println("Hello, Advent of Code 2025!")

	// rotations := aoc2025.ReadLines("aoc2025/secret_entrance.txt")
	// fmt.Println("Secret Entrance 1:", aoc2025.SecretEntrance1(rotations))
	// fmt.Println("Secret Entrance 2:", aoc2025.SecretEntrance2(rotations))

	// ranges := aoc2025.GiftShopInput("aoc2025/gift_shop.txt")
	// fmt.Println("Gift Shop 1:", aoc2025.GiftShop1(ranges))
	// fmt.Println("Gift Shop 2", aoc2025.GiftShop2(ranges))

	// banks := aoc2025.ReadLines("aoc2025/lobby.txt")
	// fmt.Println("Lobby 1:", aoc2025.Lobby1(banks))
	// fmt.Println("Lobby 2:", aoc2025.Lobby2(banks))

	// grid := aoc2025.ReadLines("aoc2025/printing_department.txt")
	// fmt.Println("Printing Department 1:", aoc2025.PrintingDepartment(grid))
	// fmt.Println("Printing Department 2:", aoc2025.PrintingDepartment2(grid))

	// cafRanges, cafIDs := aoc2025.CafeteriaInput("aoc2025/cafeteria.txt")
	// fmt.Println("Cafeteria 1:", aoc2025.Cafeteria(cafRanges, cafIDs))
	// fmt.Println("Cafeteria 2:", aoc2025.Cafeteria2(cafRanges))

	// worksheet := aoc2025.ReadLines("aoc2025/trash_compactor.txt")
	// fmt.Println("Trash Compactor 1:", aoc2025.TrashCompactor(worksheet))
	// fmt.Println("Trash Compactor 2:", aoc2025.TrashCompactor2(worksheet))

	// manifold := aoc2025.ReadLines("aoc2025/laboratory.txt")
	// fmt.Println("Laboratory 1:", aoc2025.Laboratory(manifold))
	// fmt.Println("Laboratory 2:", aoc2025.Laboratory2(manifold))

	// boxes := aoc2025.ReadLines("aoc2025/playground.txt")
	// fmt.Println("Playground 1:", aoc2025.Playground(boxes, 1000))
	// fmt.Println("Playground 2:", aoc2025.Playground2(boxes))

	// movieTheaterInput := aoc2025.ReadLines("aoc2025/movie_theater.txt")
	// fmt.Println("Movie Theater 1:", aoc2025.MovieTheater(movieTheaterInput))
	// fmt.Println("Movie Theater 2:", aoc2025.MovieTheater2(movieTheaterInput))

	// factoryInput := aoc2025.ReadLines("aoc2025/factory.txt")
	// fmt.Println("Factory 1:", aoc2025.Factory(factoryInput))
	// fmt.Println("Factory 2:", aoc2025.Factory2(factoryInput))

	// reactorLines := aoc2025.ReadLines("aoc2025/reactor.txt")
	// fmt.Println("Reactor 1:", aoc2025.ReactorFromLines(reactorLines))
	// fmt.Println("Reactor 2:", aoc2025.Reactor2FromLines(reactorLines))

	christmasLines := aoc2025.ReadLines("aoc2025/christmas_tree_farm.txt")
	fmt.Println("Christmas Tree Farm:", aoc2025.ChristmasTreeFarmFromLines(christmasLines))
}
