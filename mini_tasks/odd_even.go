package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan int, 1)
	oddStart := make(chan int, 1)
	evenStart := make(chan int, 1)

	var wg sync.WaitGroup

	// Function to print odd numbers
	printOdd := func(ch chan int, wg *sync.WaitGroup) {
		<-oddStart
		evenStart <- 1
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			<-ch // Wait for the signal
			fmt.Println(i)
			ch <- 1 // Signal the even number goroutine
		}
	}

	// Function to print even numbers
	printEven := func(ch chan int, wg *sync.WaitGroup) {
		<-evenStart
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch // Wait for the signal
			fmt.Println(i)
			ch <- 1 // Signal the odd number goroutine
		}
	}

	// Add two goroutines to the WaitGroup
	wg.Add(2)

	// Start the goroutines
	go printOdd(ch, &wg)
	go printEven(ch, &wg)

	// Start the sequence
	ch <- 1 // Signal the odd number goroutine to start
	oddStart <- 1
	// Wait for both goroutines to finish
	wg.Wait()
}
