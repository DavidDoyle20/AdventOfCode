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

func partOne(data []string) int {
	var count int
	for i, line := range data {
		for j := range line {
			if line[j] == 'X' {
				for r := -1; r < 2; r++ {
					for c := -1; c < 2; c++ {
						if r == 0 && c == 0 {
							continue
						}
						if isValid(data, i, j, r, c) {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func partTwo(data []string) int {
	var count int
	for i, line := range data {
		for j := range line {
			if line[j] == 'A' {
				if isX(data, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func isX(data []string, r int, c int) bool {
	var firstCross bool
	var secondCross bool
	if r == 0 || c == 0 || r == len(data[0])-1 || c == len(data[0])-1 {
		return false
	}
	if data[r-1][c-1] == 'S' && data[r+1][c+1] == 'M' || data[r-1][c-1] == 'M' && data[r+1][c+1] == 'S' {
		firstCross = true
	}
	if data[r-1][c+1] == 'S' && data[r+1][c-1] == 'M' || data[r-1][c+1] == 'M' && data[r+1][c-1] == 'S' {
		secondCross = true
	}
	return firstCross && secondCross
}

func isValid(data []string, r int, c int, offsetY int, offsetX int) bool {
	var xmas string
	xmas = "XMAS"

	if r+(offsetY*len(xmas)) > len(data) || c+(offsetX*len(xmas)) > len(data[0]) || r+(offsetY*len(xmas)) < -1 || c+(offsetX*len(xmas)) < -1 {
		return false
	}
	for i := range len(xmas) {
		if xmas[i] != data[r+(i*offsetY)][c+(i*offsetX)] {
			return false
		}
	}
	return true
}
