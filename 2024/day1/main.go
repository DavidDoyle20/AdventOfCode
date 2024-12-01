package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var list1, list2 []int64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		locations := strings.Fields(line)
		if len(locations) != 2 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		location1, err1 := strconv.ParseInt(locations[0], 10, 64)
		location2, err2 := strconv.ParseInt(locations[1], 10, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Skipping line with invalid numbers:", line)
		}

		list1 = append(list1, location1)
		list2 = append(list2, location2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	fmt.Println(partOne(list1, list2))
	fmt.Println(partTwo(list1, list2))
}

func partOne(list1 []int64, list2 []int64) int {
	var sum int
	for i := range list1 {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return(sum)
}

func partTwo(list1 []int64, list2 []int64) int {
	var score int64
	var prev int64
	var prevCount int64
	var pivot int64

	for _, num := range list1 {
		if num != prev {
			pivot += prevCount
			prevCount = 0
			for _, num2 := range list2[pivot:] {
				if num == num2 {
					prevCount++
				}
			}
		}
		score += num * prevCount
		prev = num
	}
	return int(score)
}