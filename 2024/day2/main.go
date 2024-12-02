package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var report []int

		if len(line) < 1 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		numStrings := strings.Fields(line)
		for _, n := range numStrings {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println("Invalid numbers on line:", line)
				continue
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	fmt.Println("Part One: ", partOne(reports))
	fmt.Println("Part Two: ", partTwo(reports))
}

// return number of safe reports
func partOne(reports [][]int) int {

	var count int
	var ascOrDesc bool

	for _, r := range reports {
		// compare first and last index to check direction
		if r[0] > r[len(r)-1] {
			// descending
			ascOrDesc = false
		} else {
			// ascending
			ascOrDesc = true
		}
		prev := r[0]
		for i := 1; i <= len(r); i++ {
			// if the end of a report is found without a false statement
			if i == len(r) {
				count++
				//fmt.Println(lineNo, r)
				break
			}
			if !compareValues(prev, r[i], ascOrDesc) {
				break
			}
			prev = r[i]
		}
	}
	return count
}

// true -> asc, false -> desc
func ascendingOrDescending(r []int) bool {
	// compare first and last index to check direction
	// need to fix this so the last one cant throw the direction off
	if len(r) == 0 {
		return false
	}

	if len(r) <= 3 {
		if r[0] > r[len(r)-1] {
			// descending
			return false
		} else {
			// ascending
			return true
		}
	} else {
		frontAvg := float64(r[0]+r[1]) / 2
		rearAvg := float64(r[(len(r)-1)]+r[len(r)-2]) / 2
		if frontAvg > rearAvg {
			return false
		} else {
			return true
		}
	}
}

func partTwo(reports [][]int) int {
	scroll := bruteForce(reports)
	var count int
	var ascOrDesc bool
	var _ int

	for lineNo, r := range reports {
		dampenerUsesLeft := 1

		ascOrDesc = ascendingOrDescending(r)

		prevIndex := 0
		for i := 1; i <= len(r); i++ {
			// if the end of a report is found without a false statement
			if i == len(r) {
				scroll = removeItem(scroll, lineNo)
				count++
				continue
			}
			if !compareValues(r[prevIndex], r[i], ascOrDesc) {
				if dampenerUsesLeft > 0 {
					dampenerUsesLeft--

					// more robust
					if i+1 < len(r) && !compareValues(prevIndex, r[i+1], ascOrDesc) {
						for z := range r {
							if isSafe(r, z) {
								count++
								break
							}
						}
						break
					}
					if prevIndex == 0 {
						nextNum := r[i+1]
						if compareValues(r[i], nextNum, ascOrDesc) {
							prevIndex = i
						}

					}
				} else {
					break
				}
			} else {
				prevIndex = i
			}
		}
	}
	fmt.Println(scroll, reports[376])
	return count
}

func compareValues(a int, b int, ascOrDesc bool) bool {
	if a == b {
		return false
	}
	diff := int(math.Abs(float64(a - b)))
	// ascending
	if ascOrDesc {
		return b > a && diff <= 3
	} else {
		// descending
		return b < a && diff <= 3
	}
}

func bruteForce(reports [][]int) []int {
	var safeReports []int
	var count int
	for lineNo, r := range reports {
		if isSafe(r, -1) {
			safeReports = append(safeReports, lineNo)
			count++
		} else {
			for i := 0; i < len(r); i++ {
				if isSafe(r, i) {
					safeReports = append(safeReports, lineNo)
					count++
					break
				}
				//skip i when looping
			}
		}
	}
	fmt.Println("Brute Force Count: ", count)
	return safeReports
}

func isSafe(r []int, skipIndex int) bool {
	// asc or desc
	ascOrDesc := ascendingOrDescending(r)

	var prevIndex int
	if skipIndex == 0 {
		prevIndex = 1
	} else {
		prevIndex = 0
	}
	for i := prevIndex + 1; i < len(r); i++ {
		if skipIndex == i {
			continue
		}
		if i == len(r) {
			return true
		}
		if !compareValues(r[prevIndex], r[i], ascOrDesc) {
			return false
		}
		prevIndex = i
	}
	return true
}

func removeItem(lines []int, item int) []int {
	index := -1
	for i, line := range lines {
		if line == item {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println(item)
	}
	lines[index] = lines[len(lines)-1]
	return lines[:len(lines)-1]
}
