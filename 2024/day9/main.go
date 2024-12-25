package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("Part One: ", partOne(text))
	fmt.Println("Part Two: ", partTwo(text))
}

func partOne(input string) int {
	id := 0
	var files []string

	for i := 0; i < len(input); i += 2 {
		occupied, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return 0
		}
		free := 0
		if i < len(input)-1 {
			free, err = strconv.Atoi(string(input[i+1]))
			if err != nil {
				return 0
			}
		}

		for temp := 0; temp < occupied; temp++ {
			files = append(files, strconv.Itoa(id))
		}
		for temp := 0; temp < free; temp++ {
			files = append(files, ".")
		}
		id++
	}
	reduced := reduce(files)
	checksum := generateChecksum(reduced)

	return checksum
}

func partTwo(input string) int {
	id := 0
	var files []string

	for i := 0; i < len(input); i += 2 {
		occupied, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return 0
		}
		free := 0
		if i < len(input)-1 {
			free, err = strconv.Atoi(string(input[i+1]))
			if err != nil {
				return 0
			}
		}

		for temp := 0; temp < occupied; temp++ {
			files = append(files, strconv.Itoa(id))
		}
		for temp := 0; temp < free; temp++ {
			files = append(files, ".")
		}
		id++
	}
	adjusted := adjust(files)
	checksum := generateChecksum(adjusted)

	return checksum
}

func adjust(input []string) []string {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	currentId := inputCopy[len(inputCopy)-1]
	idCount := 1

	for i := len(inputCopy) - 2; i >= 0; i-- {
		// if current id is different add the last black
		if inputCopy[i] != "." && currentId == "." {
			currentId = inputCopy[i]
			idCount = 0
		}
		if inputCopy[i] != currentId {
			freeSpace := 0
			// search for spot
			for j := 0; j < i+freeSpace; j++ {
				if freeSpace == idCount {
					// add the new free space
					for k := 0; k < freeSpace; k++ {
						inputCopy[j-freeSpace+k], inputCopy[i+idCount-k] = inputCopy[i+idCount-k], inputCopy[j-freeSpace+k]
					}

					break
				}
				if inputCopy[j] == "." {
					freeSpace++
				} else {
					freeSpace = 0
				}
			}
			// search for next
			currentId = inputCopy[i]
			idCount = 1
		} else {
			idCount++
		}
	}

	return inputCopy
}

func reduce(input []string) []string {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)
	j := len(inputCopy) - 1
	for i, id := range inputCopy {
		if id == "." {
			// swap with the rearmost id
			for ; j > i; j-- {
				if inputCopy[j] != "." {
					inputCopy[i], inputCopy[j] = inputCopy[j], inputCopy[i]
					break
				}
			}
		}
	}
	return inputCopy
}

func generateChecksum(input []string) int {
	var checksum int
	for i := 0; i < len(input); i++ {
		if input[i] != "." {
			num, err := strconv.Atoi(input[i])
			if err != nil {
				return 0
			}
			checksum += i * num
		}
	}
	return checksum
}
