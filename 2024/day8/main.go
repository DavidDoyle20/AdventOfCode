package main

import (
	"bufio"
	"fmt"
	"os"
)

var coordinates map[rune][]Coordinate

type Coordinate struct {
	X int
	Y int
}

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

	positions := getPositions(puzzle)
	antennaCoordinates := make([]Coordinate, 0)

	for _, values := range positions {
		for i, val := range values {
			for j := 0; j != i && j < len(values); j++ {
				xDist := val.X - values[j].X
				yDist := val.Y - values[j].Y

				antennaX := val.X + xDist
				antennaY := val.Y + yDist

				if 0 <= antennaX && antennaX < len(puzzle[0]) && 0 <= antennaY && antennaY < len(puzzle) {
					coord := Coordinate{
						X: antennaX,
						Y: antennaY,
					}
					if !containsCoordinate(antennaCoordinates, coord) {
						antennaCoordinates = append(antennaCoordinates, coord)
						count++
					}
				}

				antennaX = values[j].X - xDist
				antennaY = values[j].Y - yDist

				if 0 <= antennaX && antennaX < len(puzzle[0]) && 0 <= antennaY && antennaY < len(puzzle) {
					//fmt.Println(string(key), antennaX, antennaY)
					coord := Coordinate{
						X: antennaX,
						Y: antennaY,
					}
					if !containsCoordinate(antennaCoordinates, coord) {
						antennaCoordinates = append(antennaCoordinates, coord)
						count++
					}
				}
			}
		}
	}
	//printAntenna(puzzle)
	return count
}

func partTwo(puzzle []string) int {
	var count int

	positions := getPositions(puzzle)
	antennaCoordinates := make([]Coordinate, 0)

	for key, values := range positions {
		for i, val := range values {
			for j := 0; j != i && j < len(values); j++ {
				xDist := val.X - values[j].X
				yDist := val.Y - values[j].Y

				coord := Coordinate{
					X: val.X,
					Y: val.Y,
				}
				if !containsCoordinate(antennaCoordinates, coord) {
					antennaCoordinates = append(antennaCoordinates, coord)
					count++
				}

				antennaX := val.X + xDist
				antennaY := val.Y + yDist

				for 0 <= antennaX && antennaX < len(puzzle[0]) && 0 <= antennaY && antennaY < len(puzzle) {
					coord := Coordinate{
						X: antennaX,
						Y: antennaY,
					}
					if !containsCoordinate(antennaCoordinates, coord) {
						fmt.Println(string(key), antennaX, antennaY, xDist, yDist, val.X, val.Y)
						antennaCoordinates = append(antennaCoordinates, coord)
						count++
					}
					antennaX += xDist
					antennaY += yDist
				}

				antennaX = val.X - xDist
				antennaY = val.Y - yDist

				for 0 <= antennaX && antennaX < len(puzzle[0]) && 0 <= antennaY && antennaY < len(puzzle) {
					coord := Coordinate{
						X: antennaX,
						Y: antennaY,
					}
					if !containsCoordinate(antennaCoordinates, coord) {
						fmt.Println(string(key), antennaX, antennaY)
						antennaCoordinates = append(antennaCoordinates, coord)
						count++
					}
					antennaX -= xDist
					antennaY -= yDist
				}
			}
		}
	}
	printAntenna(antennaCoordinates, puzzle)
	return count
}

func getPositions(puzzle []string) map[rune][]Coordinate {
	data := make(map[rune][]Coordinate)
	for i, line := range puzzle {
		for j, c := range line {
			if c != '.' {
				coord := Coordinate{
					X: i,
					Y: j,
				}
				data[c] = append(data[c], coord)
			}
		}
	}
	fmt.Println(data)
	return data
}

func printAntenna(coordinates []Coordinate, puzzle []string) {
	for i, line := range puzzle {
		for j, c := range line {
			coord := Coordinate{
				X: i,
				Y: j,
			}
			if containsCoordinate(coordinates, coord) {
				if c != '.' {
					fmt.Print("|", string(c))
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}

func containsCoordinate(coordinates []Coordinate, coord Coordinate) bool {
	for _, c := range coordinates {
		if c.X == coord.X && c.Y == coord.Y {
			return true
		}
	}
	return false
}
