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

func (l *linkedList) deleteDuplicates() {
	current := l.headpointer
	for current.nextnodepointer != nil {
		if current.data == current.nextnodepointer.data {
			current.nextnodepointer = current.nextnodepointer.nextnodepointer
			l.length--
		}
		current = current.nextnodepointer
	}
}

func main() {
	myList := &linkedList{}
	node1 := &node{data: 100, nextnodepointer: nil}
	node2 := &node{data: 80, nextnodepointer: node1}
	node3 := &node{data: 60, nextnodepointer: node2}
	node7 := &node{data: 60, nextnodepointer: node3}
	node4 := &node{data: 40, nextnodepointer: node7}
	node5 := &node{data: 20, nextnodepointer: node4}
	node6 := &node{data: 20, nextnodepointer: node5}
	myList.headpointer = node6
	myList.length = 7
	myList.printListData()
	myList.deleteDuplicates()
	myList.printListData()
}
