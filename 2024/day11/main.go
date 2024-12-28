package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Length() int {
	var count int
	current := ll.head
	for current != nil {
		current = current.next
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
	fmt.Println("Done Copying")
	return newList
}

func (ll *LinkedList) Prepend(data string) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

func (ll *LinkedList) Append(data string) {
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

func (ll *LinkedList) AppendNode(node *Node) {
	if ll.head == nil {
		ll.head = node
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

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
	if len(node.data)%2 != 0 {
		return
	}
	mid := len(node.data) / 2
	left := node.data[:mid]
	right := node.data[mid:]
	left = strings.TrimLeft(left, "0")
	right = strings.TrimLeft(right, "0")

	if len(left) == 0 {
		left = "0"
	}
	if len(right) == 0 {
		right = "0"
	}

	node.data = left
	nextNode := node.next

	newNode := &Node{data: right, next: nextNode}
	node.next = newNode
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
		ll.Append(s)
	}
	//fmt.Println("Part One: ", partOne(ll, 25))
	fmt.Println("Part Two: ", partOne(ll, 75))
}

func partOne(stones LinkedList, blinks int) int {
	llist := stones.Copy()

	for i := 0; i < blinks; i++ {
		current := llist.head
		for current != nil {
			if current.data == "0" {
				current.data = "1"
			} else if len(current.data)%2 == 0 {
				current.splitNode()
				current = current.next
			} else {
				num, err := strconv.Atoi(current.data)
				if err != nil {
					return 0
				}
				num *= 2024
				current.data = strconv.Itoa(num)
			}
			current = current.next
		}
	}
	return llist.Length()
}
