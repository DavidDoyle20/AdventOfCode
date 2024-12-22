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

func partOne(puzzleMap []string) int {
	var count int
	posHistory := make([][]bool, len(puzzleMap))
	for z := range puzzleMap {
		posHistory[z] = make([]bool, len(puzzleMap[z]))
	}

	i, j := getStartingIndex(puzzleMap)
	for {
		currentCount := 0
		// up - cant move past [0][x]
		// i --> index of # or end of map
		for ; i >= 0; i-- {
			if i == 0 {
				return currentCount + count + 1
			}
			if puzzleMap[i-1][j] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}

		// right - cant move past [x][len-1]
		for ; j <= len(puzzleMap[i])-1; j++ {
			if j == len(puzzleMap[i])-1 {
				return currentCount + count + 1
			}
			if puzzleMap[i][j+1] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}
		// down - cant move past [len-1][x]
		for ; i <= len(puzzleMap)-1; i++ {
			if i == len(puzzleMap)-1 {
				return currentCount + count + 1
			}
			if puzzleMap[i+1][j] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}
		// left - cant move past [x][0]
		for ; j >= 0; j-- {
			if j == 0 {
				return currentCount + count + 1
			}
			if puzzleMap[i][j-1] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}

		// if no possible moves all spaces have been filled or there is a loop
		if currentCount == 0 {
			return count
		}
		count += currentCount
	}
}

// loop through original path and checks each pos in path
func partTwo(puzzleMap []string) int {
	var count int

	obstacleHistory := make([][]bool, len(puzzleMap))
	for z := range puzzleMap {
		obstacleHistory[z] = make([]bool, len(puzzleMap[z]))
	}

	i, j := getStartingIndex(puzzleMap)
	//fmt.Println(isLoop(puzzleMap, 6, 3))
	for {
		// up - cant move past [0][x]
		// i --> index of # or end of map
		for ; i >= 0; i-- {
			if i == 0 {
				return count
			}
			if puzzleMap[i-1][j] == '#' {
				break
			}
			if isLoop(puzzleMap, i-1, j) {
				if !obstacleHistory[i-1][j] {
					obstacleHistory[i-1][j] = true
					//fmt.Println("up     ", i-1, j)
					count++
				}
			}
		}

		// right - cant move past [x][len-1]
		for ; j <= len(puzzleMap[i])-1; j++ {
			if j == len(puzzleMap[i])-1 {
				return count
			}
			if puzzleMap[i][j+1] == '#' {
				break
			}
			if isLoop(puzzleMap, i, j+1) {
				if !obstacleHistory[i][j+1] {
					obstacleHistory[i][j+1] = true
					//fmt.Println("right  ", i, j+1)
					count++
				}
			}

		}
		// down - cant move past [len-1][x]
		for ; i <= len(puzzleMap)-1; i++ {
			if i == len(puzzleMap)-1 {
				return count
			}
			if puzzleMap[i+1][j] == '#' {
				break
			}
			if isLoop(puzzleMap, i+1, j) {
				if !obstacleHistory[i+1][j] {
					obstacleHistory[i+1][j] = true
					//fmt.Println("down   ", i+1, j)
					count++
				}
			}
		}
		// left - cant move past [x][0]
		for ; j >= 0; j-- {
			if j == 0 {
				return count
			}
			if puzzleMap[i][j-1] == '#' {
				break
			}
			if isLoop(puzzleMap, i, j-1) {
				if !obstacleHistory[i][j-1] {
					obstacleHistory[i][j-1] = true
					//fmt.Println("left   ", i, j-1)
					count++
				}
			}
		}
	}
	return count
}

func isLoop(puzzleMap []string, i int, j int) bool {
	// make obstacle map and add the new obstacle
	obstacleMap := make([]string, len(puzzleMap))
	copy(obstacleMap, puzzleMap)
	obstacleMap[i] = obstacleMap[i][:j] + "#" + obstacleMap[i][j+1:]

	// used to detect loops
	posHistory := make([][]bool, len(obstacleMap))
	for z := range obstacleMap {
		posHistory[z] = make([]bool, len(obstacleMap[z]))
	}

	i, j = getStartingIndex(puzzleMap)
	for {
		currentCount := 0
		// up - cant move past [0][x]
		// i --> index of # or end of map
		for ; i >= 0; i-- {
			// reached end -> false
			if i == 0 {
				return false
			}
			// obstacle -> continue
			if obstacleMap[i-1][j] == '#' {
				break
			}
			// if space not visited update count
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}

		// right - cant move past [x][len-1]
		for ; j <= len(obstacleMap[i])-1; j++ {
			if j == len(obstacleMap[i])-1 {
				return false
			}
			if obstacleMap[i][j+1] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}
		// down - cant move past [len-1][x]
		for ; i <= len(obstacleMap)-1; i++ {
			if i == len(obstacleMap)-1 {
				return false
			}
			if obstacleMap[i+1][j] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}
		// left - cant move past [x][0]
		for ; j >= 0; j-- {
			if j == 0 {
				return false
			}
			if obstacleMap[i][j-1] == '#' {
				break
			}
			if !posHistory[i][j] {
				currentCount += 1
				posHistory[i][j] = true
			}
		}

		// if no possible moves all spaces have been filled or there is a loop
		if currentCount == 0 {
			return true
		}
	}
}

func getStartingIndex(puzzleMap []string) (int, int) {
	for i := 0; i < len(puzzleMap); i++ {
		for j := 0; j < len(puzzleMap[i]); j++ {
			if puzzleMap[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}
