package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var stones []string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stones = strings.Split(scanner.Text(), " ")

	var ll LinkedList
	for _, s := range stones {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ll.Append(num)
	}
	fmt.Println("Part One: ", partOne(ll, 25))
	fmt.Println("Part Two: ", partTwo(ll, 75))
}

func (ll *LinkedList) Length() int {
	count := 0
	for current := ll.head; current != nil; current = current.next {
		count++
	}
	return count
}

func (ll *LinkedList) Copy() *LinkedList {
	if ll.head == nil {
		return &LinkedList{}
	}

	// new empty linked list
	newList := &LinkedList{}

	// keep track of the current position in the original list
	current := ll.head

	// create a new node with the data of the original node
	newList.head = &Node{data: current.data, next: nil}

	// keeps track of the position to append to in the new list
	newCurrent := newList.head

	// traverse the original list by checking if the current node is the last
	for current.next != nil {
		// iterate
		current = current.next
		newNode := &Node{data: current.data, next: nil}

		// link the new node to the new liset
		newCurrent.next = newNode
		newCurrent = newNode
	}
	//fmt.Println("Done Copying")
	return newList
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// for testing
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// split one node into two
// data of node = front, next = new node with back and its tail being the tail of the original
func (node *Node) splitNode() {
	digits := countDigit(node.data)
	if digits%2 != 0 {
		return
	}
	halfDigits := digits / 2
	divisor := int(math.Pow10(halfDigits))


	left := node.data / divisor
	right := node.data % divisor

	node.data = left
	nextNode := node.next
	node.next = &Node{data: right, next: nextNode}
}

func countDigit(n int) int {
	return int(math.Floor(math.Log10((float64(n)))+1))
}

func partOne(stones LinkedList, blinks int) int {
	llist := stones.Copy()

	for i := 0; i < blinks; i++ {
		current := llist.head
		for current != nil {
			if current.data == 0 {
				current.data = 1
			} else if countDigit(current.data)%2 == 0 {
				current.splitNode()
				current = current.next
			} else {
				current.data *= 2024
			}
			current = current.next
		}
	}
	return llist.Length()
}

// calculate for each item in the list and free after that
func partTwo(stones LinkedList, blinks int) int {
	count := 0
	for current := stones.head; current != nil; current = current.next {
		//fmt.Println("Processing", int(current.data))
		count += temp(current.data, blinks)
	}
	return count
}

// returns the size a list would be starting from the number n
func temp(n int, blinks int) int {
	//fmt.Println(n, blinks)
	if blinks == 0 {
		return 1
	}
	key := fmt.Sprintf("%d_%d", n, blinks)
	if val, ok := cache[key]; ok {
		return val
	}

	count := 0
	if blinks != 0 {
		blinks--
		if n == 0 {
			count += temp(1, blinks)
		} else if countDigit(n)%2 == 0 {
			left, right := splitInt(n)
			count += temp(left, blinks) + temp(right, blinks)
		} else {
			count += temp(n*2024, blinks)
		}
	}

	cache[key] = count
	return count
}

func splitInt(n int) (int, int) {
	digits := countDigit(n)
	if digits%2 != 0 {
		return 0, 0
	}
	halfDigits := digits / 2
	divisor := int(math.Pow10(halfDigits))


	left := n / divisor
	right := n % divisor

	return left, right
}





