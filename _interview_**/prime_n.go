package main

import "fmt"

// return prime numbers in go

func main() {
	fmt.Println("Enter a number ")
	var N int

	fmt.Scanln(&N)
	fmt.Println("input number ", N)
	count := 0
	for i := 2; i < N; i++ {

		if isPrime(i) {
			fmt.Println("prime", i)
			count++
		}
	}
	fmt.Println(count)
}
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
