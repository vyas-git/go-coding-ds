package main

import (
	"testing"
)

func TestLongestSubStr(t *testing.T) {
	tParam1 := "abba"
	expected := 2
	got := lengthOfLongestSubstring(tParam1)
	if got != expected {
		t.Errorf("Failed ! expected 2, got %v", got)
	} else {
		t.Logf("Accepted")
	}
}
