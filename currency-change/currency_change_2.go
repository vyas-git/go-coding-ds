package main

import (
	"fmt"
)

// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

// Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

// You may assume that you have an infinite number of each kind of coin.

// Example 1:

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
// Example 2:

// Input: coins = [2], amount = 3
// Output: -1
// Example 3:

// Input: coins = [1], amount = 0
// Output: 0
// Example 4:

// Input: coins = [1], amount = 1
// Output: 1
// Example 5:

// Input: coins = [1], amount = 2
// Output: 2
func run() {
	coins := []int{1, 2, 5}
	minWays := coinChange(coins, 11)
	fmt.Println(minWays)
}
func min(x, y int) int {

	if x < y {
		return x
	} else {
		return y
	}
}

func coinChange(coins []int, amount int) int {

	INF := 65535
	size := amount + 1

	// idx: money
	// value: minimal change count
	change := make([]int, size)

	// Initialized to infinity
	for idx, _ := range change {
		change[idx] = INF
	}

	// Initialization for $0
	change[0] = 0

	for value := 1; value <= amount; value += 1 {

		for _, coin := range coins {

			if coin > value {

				// coin value is to big, can not make change with current coin
				continue
			}

			// update dp_table, try to make change with coin
			change[value] = min(change[value-coin]+1, change[value])
		}
	}

	if change[amount] == INF {

		// Reject, no solution
		return -1

	} else {

		// Accept, return total count of coin change
		return change[amount]
	}

}
