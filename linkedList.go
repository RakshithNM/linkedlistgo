package main

import "fmt"

type node struct {
	data    int
	pointer *node
}

type linkedList struct {
	headpointer *node
	length      int
}

func (l *linkedList) prepend(n *node) {
	second := l.headpointer
	l.headpointer = n
	l.headpointer.pointer = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.headpointer
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.pointer
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.headpointer.data == value {
		l.headpointer = l.headpointer.pointer
		l.length--
		return
	}

	previousToDelete := l.headpointer
	for previousToDelete.pointer.data != value {
		if previousToDelete.pointer.pointer == nil {
			return
		}
		previousToDelete = previousToDelete.pointer
	}
	previousToDelete.pointer = previousToDelete.pointer.pointer
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 19}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.printListData()
	myList.deleteWithValue(18)
	myList.printListData()

	myList.deleteWithValue(100)
	myList.deleteWithValue(19)
	// delete from empty list
}
