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

func (l linkedList) countIterative() int {
	count := 0
	headnode := l.headnode
	for headnode != nil {
		count++
		headnode = headnode.nextnodepointer
	}
	return count
}

func (l *linkedList) countRecursive(n *node) int {
	if n.nextnodepointer == nil {
		return 1
	}
	n = n.nextnodepointer
	return 1 + l.countRecursive(n)
}

func (l *linkedList) prepend(n *node) {
	second := l.headnode
	l.headnode = n
	l.headnode.nextnodepointer = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.headnode
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	myList := linkedList{}
	n1 := &node{data: 50}
	n2 := &node{data: 60}
	n3 := &node{data: 70}
	n4 := &node{data: 80}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.prepend(n4)
	myList.insertAtTheEnd(10)
	myList.insertAtTheEnd(20)
	myList.insertAtTheEnd(30)
	myList.insertAtTheEnd(40)

	myList.printListData()
	fmt.Println(myList)
	fmt.Println(myList.countIterative())   // 8
	fmt.Println(myList.countRecursive(n4)) // 8
}
