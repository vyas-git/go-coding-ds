package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan int)
	oddStart := make(chan int)
	evenStart := make(chan int)

	var wg sync.WaitGroup

	// Function to print odd numbers
	printOdd := func(ch chan int, wg *sync.WaitGroup) {
		<-oddStart
		evenStart <- 1
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			fmt.Println("odd   ==================>", i)
			ch <- 1 // Wait for the signal
			<-ch    // Signal the even number goroutine
		}
	}

	// Function to print even numbers
	printEven := func(ch chan int, wg *sync.WaitGroup) {
		<-evenStart
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch    // Wait for the signal
			ch <- 1 // Signal the odd number goroutine
			fmt.Println("even  ==================>", i)

		}
	}

	// Add two goroutines to the WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		oddStart <- 1
	}()
	// Start the goroutines
	go printOdd(ch, &wg)
	go printEven(ch, &wg)

	wg.Wait()
}
