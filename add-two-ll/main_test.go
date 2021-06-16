package main

import (
	"fmt"
	"testing"
)

func TestAddLinkedListNodes(t *testing.T) {
	tParam1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	} // input
	tParam2 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	//input

	expected := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	got := addLinkedListNodes(tParam1, tParam2)
	if !areIdentical(expected, got) {
		t.Errorf("Failed ! expected [8,0,1,3], got %v", returnListArr(got))
	} else {
		t.Logf("Accepted")
	}
}
func returnListArr(list *ListNode) []int {
	var arrList []int
	if list.Val == 0 {
		fmt.Println("No Nodes")
		return arrList

	}
	for list != nil {
		//fmt.Print(currentList.Val, "=>")
		arrList = append(arrList, list.Val)
		list = list.Next
	}
	return arrList

}
func areIdentical(l1 *ListNode, l2 *ListNode) bool {

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false

		}
		l1 = l1.Next
		l2 = l2.Next

	}
	return l1 == nil && l2 == nil

}
