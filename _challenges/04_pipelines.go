package main

import "fmt"

func main() {
	nums := generator()
	s_nums := square(nums)
	print_n(s_nums)
}

func generator() <-chan int {
	out := make(chan int)
	go func() {
		for i := range 5 {
			out <- i
		}
		close(out)

	}()
	return out
}

func square(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range ch {
			out <- v * v

		}
		close(out)

	}()

	return out
}

func print_n(ch <-chan int) {
	for v := range ch {
		fmt.Println("square numbers", v)
	}
}
