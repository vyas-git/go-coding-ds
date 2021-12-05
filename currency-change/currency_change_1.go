package main

import "fmt"

// Given array of coins and amount
// Calculate no of ways can make amount by coins
// Example amount = 5, coins = [1,2,5]
// Output 4
// Explaination : There are four ways to make amount 5 :
// 5=1+1+1+1+1, 5=1+1+1+2, 5=1+2+2, 5=5

func main() {
	possible_ways := change(5, []int{1, 2, 5})
	fmt.Println(possible_ways)
}
func change(amount int, coins []int) int {
	ways := make([]int, amount+1)
	ways[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			ways[i] += ways[i-coin]
		}
	}
	return ways[amount]
}
