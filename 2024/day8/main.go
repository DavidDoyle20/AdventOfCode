package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var puzzle []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, line)
	}

	fmt.Println("Part One: ", partOne(puzzle))
	fmt.Println("Part Two: ", partTwo(puzzle))
}

func partOne(puzzle []string) int {
	var count int
	return count
}

func partTwo(puzzle []string) int {
	var count int
	return count
}

func getPositions(puzzle []string) {
	// A -> [(1,2),(3,4)]
}
