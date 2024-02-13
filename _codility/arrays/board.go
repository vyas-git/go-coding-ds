package main
/*
There is a board with 2 rows and N columns, represented by a matrix M. Rows are numbered from 0 to 1 from top to bottom and columns are numbered from 0 to N-1 from
left to right. Each cell contains either 0 or a 1. You know that:
* the sum of integers in the 0-th(upper) row is equal to U.
* the sum of integers in the 1-st(lower) row is equal to L.
* the sum of integers in the K-th column is equal to C[K].
Your job is to recover M based on this information.
Write a function:
class Solution { public String solution(int M, int L, int[] C);}
that given two integers U, L and an array C of N integers, as described above, return a string describing the matrix M in the following format.
The first part of the string should be description of the upper row(N characters: 0 or 1), then there shouldbe comma(,), and finally there should be description of
the lower row(N characters: 0 or 1.) The output string should not contain any whitespace.

If there exist multiple valid Ms, your function may return any one of them. If no valid M exist, your function should return the word IMPOSSIBLE.
Example:

Given U = 3, L = 2, C = [2, 1, 1, 0, 1], your function may, for example return 11001, 10100 which described the following format.
1 1 0 0 1
1 0 1 0 0

Given U = 2, L = 3, C = [0, 0, 1, 1, 2], your function should return the word IMPOSSIBLE, because no matrix exist for such conditions.

*/
import (
	"fmt"
	"strings"
)

func solution(U, L int, C []int) string {
	totalSum := 0
	for _, sum := range C {
		totalSum += sum
	}

	// Check if the sum of U and L equals the total sum
	if U+L != totalSum {
		return "IMPOSSIBLE"
	}

	N := len(C)
	matrix := make([][]int, 2)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		if C[i] == 0 {
			matrix[0][i] = 0
			matrix[1][i] = 0
		} else if C[i] == 2 {
			matrix[0][i] = 1
			matrix[1][i] = 1
			U--
			L--
		} else if C[i] == 1 {
			if U > L {
				matrix[0][i] = 1
				matrix[1][i] = 0
				U--
			} else {
				matrix[0][i] = 0
				matrix[1][i] = 1
				L--
			}
		}
	}

	var result strings.Builder
	for i := 0; i < N; i++ {
		fmt.Fprintf(&result, "%d", matrix[0][i])
	}
	result.WriteString(",")
	for i := 0; i < N; i++ {
		fmt.Fprintf(&result, "%d", matrix[1][i])
	}

	return result.String()
}

func main() {
	fmt.Println(solution(3, 2, []int{2, 1, 1, 0, 1})) // Output: 11001,10100
	fmt.Println(solution(2, 3, []int{0, 0, 1, 1, 2})) // Output: IMPOSSIBLE
}
