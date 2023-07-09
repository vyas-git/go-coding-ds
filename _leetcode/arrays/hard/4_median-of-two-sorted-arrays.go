package main

import "sort"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

/*
Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.


Example 2:

Input: nums1 = [1,2,5], nums2 = [3,4,6]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

func main() {

}

func medianOfTwoArrays(arr1 []int, arr2 []int) float32 {

	merged_arr := append(arr1, arr2...)
	sort.Ints(merged_arr)
	tLen := len(merged_arr)

	if tLen%2 == 0 {
		return float32(merged_arr[(tLen/2)-1] + merged_arr[(tLen/2)])
	} else {
		return float32(merged_arr[(tLen / 2)])
	}
}
