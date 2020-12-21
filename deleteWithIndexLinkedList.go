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

func (l linkedList) printData() {
	toPrint := l.headnode
	for l.length != 0 {
		fmt.Printf("%d\n", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Println("***************")
}

// Function to delete a node by given index
func (l *linkedList) deleteWith(index int) {
	if l.length == 0 {
		return
	}
	if index == 0 {
		l.headnode = l.headnode.nextnodepointer
		l.length--
		return
	}
	previousToDelete := l.headnode
	for i := 1; i <= l.length; i++ {
		if i != index {
			previousToDelete = previousToDelete.nextnodepointer
		} else {
			break
		}
	}
	previousToDelete.nextnodepointer = previousToDelete.nextnodepointer.nextnodepointer
	l.length--
}

func main() {
	myList := linkedList{}
	myList.insertAtTheEnd(10)
	myList.insertAtTheEnd(20)
	myList.insertAtTheEnd(30)
	myList.insertAtTheEnd(40)
	myList.printData() // 10 20 30 40
	myList.deleteWith(3)
	myList.printData() // 10 20 30
}
