package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ruleMatrix = make([][]bool, 100)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rules [][]int
	var updates [][]int

	scanner := bufio.NewScanner(file)
	rulesFlag := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			rulesFlag = false
			continue
		}
		if rulesFlag {
			// add to rules
			pages := strings.Split(line, "|")
			rule := make([]int, 2)

			rule[0], err = strconv.Atoi(pages[0])
			if err != nil {
				fmt.Println(err)
			}

			rule[1], err = strconv.Atoi(pages[1])
			if err != nil {
				fmt.Println(err)
			}

			rules = append(rules, rule)
		} else {
			// add to updates
			pages := strings.Split(line, ",")
			var pageOrder []int
			for _, page := range pages {
				pageInt, err := strconv.Atoi(page)
				if err != nil {
					fmt.Println(err)
				}
				pageOrder = append(pageOrder, pageInt)
			}
			updates = append(updates, pageOrder)
		}
	}
	generateRulesMatrix(rules)
	fmt.Println("Part One: ", partOne(updates))
	fmt.Println("Part Two: ", partTwo(updates))
}

func partOne(updates [][]int) int {
	var sum int

	for _, u := range updates {
		// add the middle element to sum
		if isValidUpdate(u) {
			midIndex := (len(u) - 1) / 2
			sum += u[midIndex]
		}
	}
	return sum
}

func partTwo(updates [][]int) int {
	var sum int

	for _, u := range updates {
		if !isValidUpdate(u) {
			update, err := fixUpdate(u)
			if err == nil {
				midIndex := (len(update) - 1) / 2
				sum += update[midIndex]
			}
		}
	}

	return sum
}

func generateRulesMatrix(rules [][]int) {
	for i := range ruleMatrix {
		ruleMatrix[i] = make([]bool, 100)
	}
	// populate rule matrix with invalid points
	for _, r := range rules {
		ruleMatrix[r[1]][r[0]] = true
	}
}

func isValidUpdate(update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			first := update[i]
			second := update[j]
			if ruleMatrix[first][second] {
				return false
			}
		}
	}
	return true
}

func fixUpdate(update []int) ([]int, error) {
	newUpdate := make([]int, len(update))
	copy(newUpdate, update)
	for {
		swaps := 0
		for i := 0; i < len(update)-1; i++ {
			for j := i + 1; j < len(update); j++ {
				first := newUpdate[i]
				second := newUpdate[j]
				if ruleMatrix[first][second] {
					newUpdate[i], newUpdate[j] = newUpdate[j], newUpdate[i]
					swaps += 1
				}
			}
		}
		if swaps == 0 {
			break
		}
	}
	return newUpdate, nil
}
