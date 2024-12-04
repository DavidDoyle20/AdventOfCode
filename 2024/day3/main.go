package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part One: ", partOne(data))
	fmt.Println("Part Two: ", partTwo(data))
}

func partOne(data []byte) int {
	var sum int
	dataString := string(data)
	mul := "mul("

	// since we are searching for mul(x,x) we dont need to process last 7 characters
	// find starting index, and final index
	for i := range len(dataString) - 7 {
		if dataString[i:i+len(mul)] == mul {
			// find the end
			for j := i + len(mul); j < len(dataString); j++ {
				if dataString[j] == ')' {
					// pass the inner data only
					sum += calculate(dataString[i+len(mul) : j])
				}
				if dataString[j] != ',' && !unicode.IsNumber(rune(dataString[j])) {
					break
				}
			}
		}
	}
	return sum
}

func partTwo(data []byte) int {
	var sum int
	doMul := true
	dataString := string(data)
	mul := "mul("
	do := "do()"
	dont := "don't()"

	// since we are searching for mul(x,x) we dont need to process last 7 characters
	// find starting index, and final index
	for i := range len(dataString) - 7 {
		if dataString[i:i+len(do)] == do {
			doMul = true
		}
		if dataString[i:i+len(dont)] == dont {
			doMul = false
		}
		if dataString[i:i+len(mul)] == mul && doMul {
			// find the end
			for j := i + len(mul); j < len(dataString); j++ {
				if dataString[j] == ')' {
					// pass the inner data only
					sum += calculate(dataString[i+len(mul) : j])
				}
				if dataString[j] != ',' && !unicode.IsNumber(rune(dataString[j])) {
					break
				}
			}
		}
	}
	return sum
}

func calculate(data string) int {
	nums := strings.Split(data, ",")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return num1 * num2
}
