package main

import (
	"fmt"
	"strings"
)

// Morse Code is delivered in a series signals which are referred to as dashes (-) or dots (.). To keep things simple for the purposes of this challenge we'll only decode letters with a maximum length of three signals.

// Morse Code Graph

// Here is the Morse Code dichotomic search table courtesy of Wikipedia

// Morse Code Examples
// -.- translates to K
// ... translates to S
// .- translates to A
// -- translates to M
// . translates to E

// Background
// You've started work as morse code translator. Unfortunately some of the signals aren't as distinguishable as others and there are times where a . seems indistinguishable from -. In these cases you write down a ? so that you can figure out what all the posibilities of that letter for that word are later.

// Task
// Write a function possibilities that will take a string signals and return an array of possible characters that the Morse code signals could represent.

// Specification
// Challenge.possibilities(signals)
// Parameters
// signals: String - The Morse code signals that needs to be parsed. The may contain one or more unknown signals (?).

// Return Value
// List - The list of possible letters for the given group of signals. Letters will always be ordered from how they appear on the chart, from left to right.

// Constraints
// There will be no more than 3 characters within signals.

// 0 - 3 unknown signals may be given.

// Examples
// signals Return Value
// "." ["E"]
// "-" ["T"]
// "-." ["N"]
// "..." ["S"]
// "..-" ["U"]
// "?" ["E","T"]
// ".?" ["I","A"]
// "?-?" ["R","W","G","O"]
func main() {
	fmt.Println(possiblites("?-?"))
}
func possiblites(signals string) []string {
	if strings.Contains(signals, "?") {
		return checkWildCard(signals)
	} else {
		return []string{moreseMap[signals]}
	}
}
func checkWildCard(signals string) []string {
	length := len(signals)
	destinations := []string{}
	if length == 1 {
		destinations = []string{"E", "T"}
	} else {

		indexes := []int{}
		var chars = strings.Split(signals, "")

		// Save letter positions except "?"
		for index, letter := range chars {
			if letter != "?" {
				indexes = append(indexes, index)
			}
		}

		// Check in morse map
		for morseCode, _ := range moreseMap {
			// Try till singal path length == morse code path length
			if len(morseCode) != len(signals) {
				continue
			}

			// consider all morse code which has same signal length
			// & matching characters(dots or dashes) except "?"
			valid := false
			for _, v := range indexes {
				if signals[v] == morseCode[v] {
					valid = true
				} else {
					valid = false
					break
				}

			}
			if valid == true {
				destinations = append(destinations, moreseMap[morseCode])
				fmt.Println(destinations)
			} else {
				continue
			}
			//	fmt.Println(destinations)
		}
	}
	return destinations
}

var moreseMap = map[string]string{
	".":   "E",
	"-":   "T",
	"..":  "I",
	".-":  "A",
	"-.":  "N",
	"--":  "M",
	"...": "S",
	"..-": "U",
	".-.": "R",
	".--": "W",
	"-..": "D",
	"-.-": "K",
	"--.": "G",
	"---": "O",
}
