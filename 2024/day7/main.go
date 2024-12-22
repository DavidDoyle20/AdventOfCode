package main

import (
	"bufio"
	"fmt"
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

	var equations [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		nums[0] = strings.Trim(nums[0], ":")

		intNums := make([]int, len(nums))
		for i, n := range nums {
			number, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err)
			}
			intNums[i] = number
		}
		equations = append(equations, intNums)
	}

	fmt.Println("Part One: ", partOne(equations))
	fmt.Println("Part Two: ", partTwo(equations))
}

func partOne(equations [][]int) int {
	var sum int
	// each equation
	for _, equation := range equations {
		targetNum := equation[0]
		num := addAndMultiply(targetNum, equation[1:])
		if num != -1 {
			sum += num
		}
	}
	return sum
}

func addAndMultiply(target int, eq []int) int {
	if len(eq) == 1 {
		return eq[0]
	}
	mul := addAndMultiply(target, append([]int{eq[0] * eq[1]}, eq[2:]...))
	add := addAndMultiply(target, append([]int{eq[0] + eq[1]}, eq[2:]...))

	if mul == target || add == target {
		return target
	} else {
		return -1
	}
}

func partTwo(equations [][]int) int {
	var sum int
	// each equation
	for _, equation := range equations {
		targetNum := equation[0]
		num := addMultiplyOr(targetNum, equation[1:])
		if num != -1 {
			sum += num
		}
	}
	return sum
}

func addMultiplyOr(target int, eq []int) int {
	if len(eq) == 1 {
		return eq[0]
	}
	mul := addMultiplyOr(target, append([]int{eq[0] * eq[1]}, eq[2:]...))
	add := addMultiplyOr(target, append([]int{eq[0] + eq[1]}, eq[2:]...))

	res, err := strconv.Atoi(fmt.Sprintf("%d%d", eq[0], eq[1]))
	if err != nil {
		fmt.Println(err)
	}
	or := addMultiplyOr(target, append([]int{res}, eq[2:]...))

	if mul == target || add == target || or == target {
		return target
	} else {
		return -1
	}
}
