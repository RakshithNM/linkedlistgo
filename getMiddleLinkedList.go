package main

import "fmt"

type node struct {
	data            int
	nextnodepointer *node
}

type linkedList struct {
	headnodepointer *node
	length          int
}

// Method to prepend a node to the linked list
func (l *linkedList) prepend(n *node) {
	second := l.headnodepointer
	l.headnodepointer = n
	l.headnodepointer.nextnodepointer = second
	l.length++
}

// Method to get the middle node of the linked list
func (l *linkedList) getMiddle() int {
	slowPointer := l.headnodepointer
	fastPointer := l.headnodepointer
	for fastPointer != nil && fastPointer.nextnodepointer != nil {
		fastPointer = fastPointer.nextnodepointer.nextnodepointer
		slowPointer = slowPointer.nextnodepointer
	}
	return slowPointer.data
}

func main() {
	myList := linkedList{}
	n1 := &node{data: 10}
	n2 := &node{data: 20}
	n3 := &node{data: 30}
	n4 := &node{data: 40}
	n5 := &node{data: 50}
	n6 := &node{data: 60}
	n7 := &node{data: 70}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.prepend(n4)
	myList.prepend(n5)
	myList.prepend(n6)
	myList.prepend(n7)
	fmt.Println(myList.getMiddle()) // 40
}
