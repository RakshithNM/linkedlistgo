package main

import "fmt"

type node struct {
	data            int
	nextheadpointer *node
}

type linkedList struct {
	headpointer *node
	length      int
}

func (l *linkedList) prepend(n *node) {
	second := l.headpointer
	l.headpointer = n
	l.headpointer.nextheadpointer = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.headpointer
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.nextheadpointer
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) deleteDuplicates() {
	pointer1 := l.headpointer
	for pointer1 != nil && pointer1.nextheadpointer != nil {
		pointer2 := pointer1
		for pointer2.nextheadpointer != nil {
			if pointer1.data == pointer2.nextheadpointer.data {
				pointer2.nextheadpointer = pointer2.nextheadpointer.nextheadpointer
				l.length--
			} else {
				pointer2 = pointer2.nextheadpointer
			}
		}
		pointer1 = pointer1.nextheadpointer
	}
}

func main() {
	myList := &linkedList{}
	node1 := &node{data: 90, nextheadpointer: nil}
	node2 := &node{data: 80, nextheadpointer: node1}
	node3 := &node{data: 80, nextheadpointer: node2}
	node4 := &node{data: 60, nextheadpointer: node3}
	node5 := &node{data: 50, nextheadpointer: node4}
	node6 := &node{data: 60, nextheadpointer: node5}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.printListData()
	myList.deleteDuplicates()
	myList.printListData()
}
