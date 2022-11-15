package main

import "fmt"

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(9, items))
}
func binarySearch(element int, list []int) (index int) {
	left := 0
	right := len(list) - 1
	index = -1
	for left <= right {
		median := (left + right) / 2
		fmt.Println(left, right, median, list[median])
		if element == list[median] {
			index = median
			return index
		} else if element > list[median] {
			left = median + 1
		} else if element < list[median] {
			right = median - 1
		}

	}
	return index
}
