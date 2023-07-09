package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/
func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(findTwoSum(nums, target))
}

func findTwoSum(a []int, target int) (int, int) {
	var hashMap = make(map[int]int)

	for i, v := range a {
		if n2, ok := hashMap[target-v]; ok {
			return n2, i
		}
		hashMap[v] = i

	}
	return -1, -1
}
