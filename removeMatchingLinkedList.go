package main

import "fmt"

type node struct {
	data            int
	nextnodepointer *node
}

type linkedList struct {
	headpointer *node
	length      int
}

func (l linkedList) printListData() {
	toPrint := l.headpointer
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) deleteAllValuesOf(inValue int) *node {
	current := l.headpointer
	for l.headpointer.data == inValue {
		l.headpointer = l.headpointer.nextnodepointer
	}
	for current.nextnodepointer != nil {
		if current.nextnodepointer.data == inValue {
			current.nextnodepointer = current.nextnodepointer.nextnodepointer
		} else {
			current = current.nextnodepointer
		}
	}
	return l.headpointer
}

func main() {
	myList := &linkedList{}
	node1 := &node{data: 100, nextnodepointer: nil}
	node2 := &node{data: 100, nextnodepointer: node1}
	node3 := &node{data: 80, nextnodepointer: node2}
	node4 := &node{data: 60, nextnodepointer: node3}
	node5 := &node{data: 100, nextnodepointer: node4}
	myList.headpointer = node5
	myList.length = 5
	myList.printListData()
	final := myList.deleteAllValuesOf(100)
	for final != nil {
		fmt.Printf("%d ", final.data)
		final = final.nextnodepointer
	}
}
