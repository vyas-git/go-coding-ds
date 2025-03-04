package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 1}, 3))
}
func threeSumClosest(nums []int, target int) int {
	closest := -target
	lenN := len(nums)
	sort.Ints(nums)

	if lenN == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	for i := 0; i < lenN; i++ {

		l, r := i+1, lenN-1

		for l < r && l > 0 {
			n1 := nums[i]
			n2 := nums[l]
			n3 := nums[r]

			sum := n1 + n2 + n3

			closest = max(sum, closest)

			if sum < target {
				l++
			} else if sum > target {
				r--
			} else if sum == target {
				l++
			}

		}
	}
	return closest
}
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
