// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]

// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Explanation:

// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
// Example 2:

// Input: strs = [""]

// Output: [[""]]

// Example 3:

// Input: strs = ["a"]

// Output: [["a"]]

package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortKey(s string) string {
	strs := strings.Split(s, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	group_map := make(map[string][]string)
	for _, w := range strs {
		sort_key := sortKey(w)
		group_map[sort_key] = append(group_map[sort_key], w)
	}

	result := [][]string{}

	for _, v := range group_map {
		result = append(result, v)
	}
	fmt.Println(result)
}
