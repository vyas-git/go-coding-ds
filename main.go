package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("===== Linked List ============")

	new_list := List{}

	new_list.insertNode(1)
	new_list.insertNode(2)
	new_list.insertNode(3)

	new_list.printList()

	fmt.Println("\n\n\n\n\n")

	fmt.Println("===== Quick Sort ============")

	fmt.Println(quickSort([]int{3, 2, 5, 1, 10, 11, 13}))

	fmt.Println("\n\n\n\n\n")
	fmt.Println("===== Binary Search ============")

	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(items, 63))
}

/** ------Linked List ------- **/
type Node struct {
	data int
	next *Node
}
type List struct {
	head *Node
}

func (l *List) insertNode(val int) {
	new_node := &Node{data: val}

	if l.head == nil {
		l.head = new_node
		return
	}

	new_node.next = l.head
	l.head = new_node

}

func (l *List) printList() {
	currentList := l.head
	for currentList != nil {
		fmt.Println(currentList.data)
		currentList = currentList.next
	}

}

func quickSort(a []int) []int {
	left, right := 0, len(a)-1

	if len(a) < 2 {
		return a
	}

	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

func binarySearch(a []int, element int) int {
	left, right := 0, len(a)-1

	for left <= right {
		median := (left + right) / 2

		if element == a[median] {
			return median
		} else if element > a[median] {
			left = median + 1
		} else if element < a[median] {
			right = median - 1
		}
	}
	return -1
}
