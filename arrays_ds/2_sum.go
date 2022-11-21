package main

func main() {

}

func twoSum(nums []int, target int) []int {
	var mapIndixesNums = make(map[int]int)
	var result []int
	for key, n := range nums {
		n1 := target - n
		if _, ok := mapIndixesNums[n1]; ok {
			result = append(result, mapIndixesNums[n1], key)
			return result
		}
		mapIndixesNums[n] = key
	}
	return result

}
