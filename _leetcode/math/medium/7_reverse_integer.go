package main

import (
	"fmt"
	"math"
)

/*  Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

*/

func main() {
	fmt.Println(reverseInt(123))
}
func reverseInt(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}
	return rev
}
