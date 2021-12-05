package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	// animals := []string{"snail", "dog", "cow", "elephant", "chicken", "mouse"}
	// sorted_animals := customSort(animals)

	// fmt.Println("Custom Sort -- ", sorted_animals)

	// sorted_animals = inBuiltSort(animals)
	// fmt.Println("In Built Sort -- ", sorted_animals)

	fmt.Println(quickSort([]int{2, 1, 3}))
	fmt.Println(count)
}

var count = 0

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
		count++
		fmt.Println(count)
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a

}
func inBuiltSort(animals []string) []string {
	sort.Strings(animals)
	return animals
}

func customSort(animals []string) []string {
	tmp := ""
	for i := 0; i < len(animals); i++ {

		for j := 0; j < len(animals); j++ {
			//check and swap

			str_1 := animals[i]
			str_2 := animals[j]

			ascii_str_1 := []byte(str_1)
			ascii_str_2 := []byte(str_2)

			if ascii_str_1[0] < ascii_str_2[0] {

				tmp = str_2
				animals[j] = str_1
				animals[i] = tmp

			} else if ascii_str_1[0] == ascii_str_2[0] { // both first ascii values are same, then check and swap with other ascii values

				short_string_len := 0 // ex : cow,chicken take shorter len of two strs

				if len(str_1) == len(str_1) {
					short_string_len = len(str_1) // if both are same length

				} else if len(str_1) < len(str_2) {
					short_string_len = len(str_1) // if first string is short
				} else {
					short_string_len = len(str_2) // second string is short
				}
				for ii := 0; ii < short_string_len; ii++ { // loop till shorter string len
					// check and swap
					if ascii_str_1[ii] < ascii_str_2[ii] {
						tmp = str_2
						animals[j] = str_1
						animals[i] = tmp
						break
					}
				}
			}
		}
	}
	return animals

}
