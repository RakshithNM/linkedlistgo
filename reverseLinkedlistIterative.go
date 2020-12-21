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
		fmt.Printf("%d\t", toPrint.data)
		l.length--
		toPrint = toPrint.nextnodepointer
	}
	fmt.Println("")
}

func (l *linkedList) reverseIterative() {
	var previous *node = nil
	var current *node = l.headnode
	var next *node = l.headnode
	for next != nil {
		next = next.nextnodepointer
		current.nextnodepointer = previous
		previous = current
		current = next
	}
	l.headnode = previous
}

func main() {
	myList := linkedList{}
	myList.insertAtTheEnd(10)
	myList.insertAtTheEnd(20)
	myList.insertAtTheEnd(30)
	myList.insertAtTheEnd(40)
	myList.insertAtTheEnd(50)
	myList.printData() // 10 20 30 40 50
	myList.reverseIterative()
	myList.printData() // 50 40 30 20 10
}
