package main

import "fmt"

type node struct {
	data            int
	nextnodepointer *node
}

type linkedList struct {
	headnode *node
	length   int
}

func (l *linkedList) insertAtTheEnd(value int) {
	n := node{value, nil}
	if l.length == 0 {
		l.headnode = &n
		l.length++
		return
	}

	ptr := l.headnode
	for i := 0; i < l.length; i++ {
		if ptr.nextnodepointer == nil {
			ptr.nextnodepointer = &n
			l.length++
			return
		}
		ptr = ptr.nextnodepointer
	}
}

// Function to return the node of the given index in the linkedList
func (l linkedList) getNth(index int) int {
	pointer := l.headnode
	for i := 0; i < l.length; i++ {
		if i == index {
			return pointer.data
		}
		pointer = pointer.nextnodepointer
	}
	return -1 // Return -1 if no node present at the given index
}

func main() {
	myList := linkedList{}
	myList.insertAtTheEnd(10)
	myList.insertAtTheEnd(20)
	myList.insertAtTheEnd(30)
	myList.insertAtTheEnd(40)
	fmt.Println(myList.getNth(2)) // 30
}
