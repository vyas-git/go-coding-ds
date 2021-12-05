package main

import "fmt"

func main() {
	var number, reminder, result int
	fmt.Println("Enter Any Number")
	fmt.Scan(&number)

	for number > 0 {
		reminder = number % 10
		result = result*10 + reminder
		number = number / 10
	}
	fmt.Println(result)
}
