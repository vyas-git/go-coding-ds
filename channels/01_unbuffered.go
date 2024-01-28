package main

import "fmt"

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {

			v := <-ch
			fmt.Println("goroutine A : ", v)

			if v == 10 {
				close(ch)
				quit <- true
				return
			}
			v++
			ch <- v
		}

	}()

	go func() {
		for {

			v := <-ch
			fmt.Println("goroutine B : ", v)
			if v == 10 {
				close(ch)
				quit <- true
				return
			}
			v++
			ch <- v
		}
	}()

	// start channel
	ch <- 0
	<-quit
}
