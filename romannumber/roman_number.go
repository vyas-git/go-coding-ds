package main

import (
	"fmt"
)

func main() {
	roman := intToRoman(1304)
	fmt.Println(roman)
}

// Create a function taking a positive integer as its parameter and returning a string containing the Roman Numeral representation of that integer.

// Modern Roman numerals are written by expressing each digit separately starting with the left most digit and skipping any digit with a value of zero. In Roman numerals 1990 is rendered: 1000=M, 900=CM, 90=XC; resulting in MCMXC. 2008 is written as 2000=MM, 8=VIII; or MMVIII. 1666 uses each Roman symbol in descending order: MDCLXVI.

// Examples:

// Input	Output
// 1000	"M"
// 1990	"MCMXC"
// 2008	"MMVIII"
// 1666	"MDCLXVI"
// Help:

// Symbol    Value
// I          1
// V          5
// X          10
// L          50
// C          100
// D          500
// M          1,000
// Remember that there can't be more than 3 identical symbols in a row.

// More about roman numerals - http://en.wikipedia.org/wiki/Roman_numerals

// Specification
// numerals(n)
// Converts from a decimal number to Roman numerals

// Parameters
// n: int - Number to convert to Roman numerals

// Return Value
// String - Resulting Roman numeral

// Examples
// n	Return Value
// 89	"LXXXIX"
// 91	"XCI"
// 1889	"MDCCCLXXXIX"
func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}
