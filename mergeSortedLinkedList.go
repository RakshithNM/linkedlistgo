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

func (l *linkedList) prepend(n *node) {
	second := l.headpointer
	l.headpointer = n
	l.headpointer.nextnodepointer = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.headpointer
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Printf("\n")
}

func merge(ll1 *linkedList, ll2 *linkedList, headNodeSorting *node) *node {
	if ll1.headpointer == nil {
		return ll2.headpointer
	}
	if ll2.headpointer == nil {
		return ll1.headpointer
	}
	if ll1.headpointer != nil && ll2.headpointer != nil {
		if ll1.headpointer.data <= ll2.headpointer.data {
			headNodeSorting = ll1.headpointer
			ll1.headpointer = headNodeSorting.nextnodepointer
		} else {
			headNodeSorting = ll2.headpointer
			ll2.headpointer = headNodeSorting.nextnodepointer
		}
	}
	sortedHead := headNodeSorting
	for ll1.headpointer != nil && ll2.headpointer != nil {
		if ll1.headpointer.data <= ll2.headpointer.data {
			headNodeSorting.nextnodepointer = ll1.headpointer
			headNodeSorting = ll1.headpointer
			ll1.headpointer = headNodeSorting.nextnodepointer
		} else {
			headNodeSorting.nextnodepointer = ll2.headpointer
			headNodeSorting = ll2.headpointer
			ll2.headpointer = headNodeSorting.nextnodepointer
		}
	}

	if ll1.headpointer == nil {
		headNodeSorting.nextnodepointer = ll2.headpointer
	}

	if ll2.headpointer == nil {
		headNodeSorting.nextnodepointer = ll1.headpointer
	}
	return sortedHead
}

func main() {
	myListP := &linkedList{}
	//myListP.headpointer = nil
	node1 := &node{data: 100, nextnodepointer: nil}
	node2 := &node{data: 80, nextnodepointer: node1}
	node4 := &node{data: 40, nextnodepointer: node2}
	node5 := &node{data: 20, nextnodepointer: node4}

	myListQ := &linkedList{}
	node6 := &node{data: 90, nextnodepointer: nil}
	node7 := &node{data: 70, nextnodepointer: node6}
	node9 := &node{data: 30, nextnodepointer: node7}
	node10 := &node{data: 10, nextnodepointer: node9}

	myListP.prepend(node1)
	myListP.prepend(node2)
	myListP.prepend(node4)
	myListP.prepend(node5)
	fmt.Printf("First sorted linked list: \t")
	myListP.printListData()

	myListQ.prepend(node6)
	myListQ.prepend(node7)
	myListQ.prepend(node9)
	myListQ.prepend(node10)
	fmt.Printf("Second sorted linked list: \t")
	myListQ.printListData()

	sortedHead := merge(myListP, myListQ, nil)
	fmt.Printf("Final merged sorted linked list: \t")
	for sortedHead != nil {
		fmt.Printf("%d ", sortedHead.data)
		sortedHead = sortedHead.nextnodepointer
	}
}
