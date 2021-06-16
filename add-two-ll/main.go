/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func addLinkedListNodes_v2(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var mainAddress = &ListNode{Val: 0}
	var currentAddress = mainAddress
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
		} else {
			val1 = 0
		}
		if l2 != nil {
			val2 = l2.Val
		} else {
			val2 = 0
		}

		sum := carry + val1 + val2
		carry = sum / 10
		currentAddress.Next = &ListNode{Val: sum % 10}
		currentAddress = currentAddress.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	if carry > 0 {
		currentAddress.Next = &ListNode{Val: carry}

	}
	return mainAddress.Next
}

func addLinkedListNodes(l1 *ListNode, l2 *ListNode) *ListNode {
	var mainAddress = &ListNode{Val: 0}
	var sum int
	for currentAddress := mainAddress; l1 != nil || l2 != nil || sum > 0; currentAddress = currentAddress.Next {

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		currentAddress.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}

	return mainAddress.Next
}
