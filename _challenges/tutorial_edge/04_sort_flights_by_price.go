package main

// In this challenge, you are going to be attempting to sort a list of Flights based on their price from highest to lowest.

// You will have to implement the SortByPrice function that takes in a slice of type Flight and returns the sorted list of Flights.

// In order to help you see what is going on, you have been provided a very quick printFlights function which you can use to print the flights out.

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	// Origin      string
	// Destination string
	Price int
}
type ByPrice []Flight

func (f ByPrice) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByPrice) Less(i, j int) bool { return f[i].Price > f[j].Price }
func (f ByPrice) Len() int           { return len(f) }

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	//sort.Sort(ByPrice(flights))
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price > flights[j].Price
	})
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Price: %d", flight.Price)
	}
}

func main() {
	// an empty slice of flights
	flights := []Flight{
		{300},
		{150},
		{200},
		{250},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
