package main

/*

An array A consisting of N integers is given. The elements of array A together represent a chain, and each element represents the strength of each link in the chain. We want to divide this chain into three smaller chains. All we can do is break the chain in exactly two non-adjacent positions. More precisely, we should break links P,Q (0 < P < Q < N - 1, Q - P > 1), resulting in three chains [0, P - 1], [P 1, Q - 1], [Q 1, N - 1]. The total cost of this operation is equal to A[P] A[Q].

For example, consider an array such that:

A[0] = 5
A[1] = 2
A[2] = 4
A[3] = 6
A[4] = 3
A[5] = 7
We can choose to break the following links:

(1-3): total cost is 2 6 = 8
(1-4): total cost is 2 3 = 5
(2-4): total cost is 4 3 = 7
Write a function
class Solution {
public int solution(int[] A);
}
that, given an array A of N integers, returns the minimal cost of dividing the chain into three pieces. Given the example above, the function should return 5.
*/
import "math"

func solution(A []int) int {
	N := len(A)
	minCost := math.MaxInt32

	for P := 1; P < N-1; P++ {
		for Q := P + 2; Q < N-1; Q++ { // Adjusted loop range for Q
			cost := A[P] + A[Q]
			if cost < minCost {
				minCost = cost
			}
		}
	}

	return minCost
}

func main() {
	A := []int{5, 2, 4, 6, 3, 7}
	result := solution(A)
	println(result) // Output: 5
}
