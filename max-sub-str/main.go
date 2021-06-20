package main

import (
	"fmt"
)

func main() {
	str := "abba"
	max := lengthOfLongestSubstring(str)
	fmt.Println(max)
}
func lengthOfLongestSubstring(str string) int {
	max, left, str_map := 0, 0, make(map[rune]int)
	for index, char := range str {
		if _, ok := str_map[char]; ok == true && str_map[char] >= left {
			if index-left > max {
				max = index - left
			}
			left = str_map[char] + 1

		}
		str_map[char] = index
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}
