package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tParam1 := []int{2, 7, 11, 15} // input
	tParam2 := 9                   //input
	got := twoSum(tParam1, tParam2)
	if tParam1[got[0]]+tParam1[got[1]] != tParam2 {
		t.Errorf("Failed ! expected [0,1], got %v", got)
	} else {
		t.Logf("Accepted")
	}
}
