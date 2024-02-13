package main

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
