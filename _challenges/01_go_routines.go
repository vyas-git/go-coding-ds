package main

import "fmt"

// Assignment 1: Goroutine with Channel Problem: Write a Go program that calculates the sum of numbers from 1 to N concurrently using goroutines and channels.
// The program should take the value of N as input from the user.

func main2() {
	var N int
	fmt.Println("Enter a number ")

	fmt.Scanln(&N)
	ch := make(chan int)
	go calculcateSum(0, N/2, ch)
	go calculcateSum(N/2+1, N, ch)

	s1 := <-ch
	s2 := <-ch
	fmt.Println(s1 + s2)

}

func calculcateSum(start, end int, ch chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
}
