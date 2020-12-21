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

func (l linkedList) printListData() {
	toPrint := l.headnodepointer
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
}

// Method to prepend a node to the linked list
func (l *linkedList) prepend(n *node) {
	second := l.headnodepointer
	l.headnodepointer = n
	l.headnodepointer.nextnodepointer = second
	l.length++
}

func main() {
	myList := linkedList{}
	n1 := &node{data: 10}
	n2 := &node{data: 20}
	n3 := &node{data: 30}
	n4 := &node{data: 40}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.prepend(n4)
	myList.printListData() // 40 30 20 10
}
