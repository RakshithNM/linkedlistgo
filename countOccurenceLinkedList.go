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

func (l linkedList) countOccurence(value int) int {
	count := 0
	headnode := l.headnodepointer
	for headnode != nil {
		if headnode.data == value {
			count++
		}
		headnode = headnode.nextnodepointer
	}
	return count
}

func main() {
	myList := linkedList{}
	n1 := &node{data: 10}
	n2 := &node{data: 0}
	n3 := &node{data: 10}
	n4 := &node{data: 10}
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
	fmt.Println(myList.countOccurence(10))
}
