package main

import "fmt"

func main() {
	// test
	input1 := "babad"
	input2 := "cbbd"

	output1 := longestPalindrome(input1)
	output2 := longestPalindrome(input2)

	expected1 := []string{"bab", "aba"}
	expected2 := "bb"

	if expected1[0] == output1 || expected1[1] == output1 {
		fmt.Println("Success input 1 ", output1)
	} else {
		fmt.Println("Failed input 1 ", output1)
	}

	if expected2 == output2 {
		fmt.Println("Success input 2 ", output2)
	} else {
		fmt.Println("Failed input 2 ", output2)
	}

}

// Given a string s, return the longest palindromic substring in s.

// A string is called a palindrome string if the reverse of that string is the same as the original string.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

func longestPalindrome(s string) string {

	if len(s) == 0 || len(s) == 1 {
		return s
	}
	longest_pl := ""
	// iterate each character in string
	for i := 0; i < len(s); i++ {

		// from every character left and right charcter will be same
		// So expand through from centre left and right

		case1_pl := checkFromCentre(i-1, i+1, s)

		// Some cases may have duplciate charcters adjacent
		// so start left adjacent to right
		case2_pl := checkFromCentre(i, i+1, s)

		if len(case1_pl) > len(longest_pl) {
			longest_pl = case1_pl
		}

		if len(case2_pl) > len(longest_pl) {
			longest_pl = case2_pl
		}

	}
	// if no palindrome then return first character of string
	if longest_pl == "" {
		return string(s[0])
	}
	return longest_pl
}

func checkFromCentre(l int, r int, s string) string {
	pl_str := ""
	for l >= 0 && r < len(s) && s[l] == s[r] {
		pl_str = s[l : r+1]
		l--
		r++
	}
	return pl_str
}
