package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	row int
	col int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hikingMap [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var intArr []int
		for _, c := range line {
			intArr = append(intArr, int(c-'0'))
		}
		hikingMap = append(hikingMap, intArr)
	}
	fmt.Println("Part One: ", partOne(hikingMap))
	fmt.Println("Part Two: ", partTwo(hikingMap))
}

func partOne(hikingMap [][]int) int {
	var sum int
	for i := range hikingMap {
		for j := range hikingMap[i] {
			if hikingMap[i][j] == 0 {
				var coords []Coordinate
				sum += pathFind(hikingMap, i, j, &coords)
			}
		}
	}
	return sum
}

func partTwo(hikingMap [][]int) int {
	var sum int
	for i := range hikingMap {
		for j := range hikingMap[i] {
			if hikingMap[i][j] == 0 {
				sum += pathFindDistinct(hikingMap, i, j)
			}
		}
	}
	return sum
}

func pathFind(hikingMap [][]int, row int, col int, tailCoords *[]Coordinate) int {
	head := hikingMap[row][col]
	tempScore := 0

	if head == 9 {
		coord := newCoord(row, col)
		if containsCoord(*tailCoords, coord) {
			return 0
		} else {
			*tailCoords = append(*tailCoords, coord)
			return 1
		}
	}
	// can move up
	if row != 0 && hikingMap[row-1][col] == head+1 {
		tempScore += pathFind(hikingMap, row-1, col, tailCoords)
	}
	// can move down
	if row != len(hikingMap)-1 && hikingMap[row+1][col] == head+1 {
		tempScore += pathFind(hikingMap, row+1, col, tailCoords)
	}
	// can move left
	if col != 0 && hikingMap[row][col-1] == head+1 {
		tempScore += pathFind(hikingMap, row, col-1, tailCoords)
	}
	// can move right
	if col != len(hikingMap[0])-1 && hikingMap[row][col+1] == head+1 {
		tempScore += pathFind(hikingMap, row, col+1, tailCoords)
	}

	return tempScore
}

func pathFindDistinct(hikingMap [][]int, row int, col int) int {
	head := hikingMap[row][col]
	tempScore := 0

	if head == 9 {
		return 1
	}
	// can move up
	if row != 0 && hikingMap[row-1][col] == head+1 {
		tempScore += pathFindDistinct(hikingMap, row-1, col)
	}
	// can move down
	if row != len(hikingMap)-1 && hikingMap[row+1][col] == head+1 {
		tempScore += pathFindDistinct(hikingMap, row+1, col)
	}
	// can move left
	if col != 0 && hikingMap[row][col-1] == head+1 {
		tempScore += pathFindDistinct(hikingMap, row, col-1)
	}
	// can move right
	if col != len(hikingMap[0])-1 && hikingMap[row][col+1] == head+1 {
		tempScore += pathFindDistinct(hikingMap, row, col+1)
	}

	return tempScore
}

func newCoord(row int, col int) Coordinate {
	return Coordinate{row: row, col: col}
}

func containsCoord(coordinates []Coordinate, coord Coordinate) bool {
	for _, c := range coordinates {
		if c.row == coord.row && c.col == coord.col {
			return true
		}
	}
	return false
}

func printMap(arr [][]int) {
	for _, line := range arr {
		fmt.Println(line)
	}
}
