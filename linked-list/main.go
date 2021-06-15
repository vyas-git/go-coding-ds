package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func main() {
	var list LinkedList
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	list.append(6)
	list.append(7)
	list.append(8)
	list.append(9)
	list.append(10)
	list.append(11)
	list.deleteValNode(5)
	list.printList()
}
func (l *LinkedList) printList() {

	if l.head == nil {
		fmt.Println("No Nodes")
		return

	}
	currentList := l.head
	for currentList != nil {
		fmt.Print(currentList.data, "=>")
		currentList = currentList.next
	}
	fmt.Println("")

}
func (l *LinkedList) prepend(val int) { // add node at begining
	newNode := ListNode{data: val}

	if l.head == nil {
		l.head = &newNode
		return
	}
	newNode.next = l.head
	l.head = &newNode
}
func (l *LinkedList) append(val int) { // add node at end
	newNode := ListNode{data: val}
	if l.head == nil {
		l.head = &newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &newNode
}
func (l *LinkedList) deleteValNode(val int) {
	if l.head.data == val {
		l.head = l.head.next
		return
	}
	current_list := l.head
	for current_list.next.data != val {
		if current_list.next == nil {
			return
		}
		current_list = current_list.next
	}
	current_list.next = current_list.next.next
}
