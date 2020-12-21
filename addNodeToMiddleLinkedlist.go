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

func (l *linkedList) prepend(n *node) {
	temp := l.headnode
	l.headnode = n
	n.nextnodepointer = temp
	l.length++
}

func (l linkedList) printData() {
	toPrint := l.headnode
	for l.length != 0 {
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) insertAtMiddle(n *node) {
	toTraverse := l.headnode
	c := 0
	for toTraverse != nil {
		if l.length%2 == 0 {
			c = l.length / 2
		} else {
			c = (l.length + 1) / 2
		}
		toTraverse = toTraverse.nextnodepointer
	}

	current := l.headnode
	var toInsertAfter *node = nil
	for i := 0; i < c; i++ {
		toInsertAfter = current
		current = current.nextnodepointer
	}

	n.nextnodepointer = toInsertAfter.nextnodepointer
	toInsertAfter.nextnodepointer = n
	l.length++
}

func main() {
	myList1 := linkedList{}
	n1 := &node{data: 10}
	n2 := &node{data: 20}
	n3 := &node{data: 30}
	myList1.prepend(n1)
	myList1.prepend(n2)
	myList1.prepend(n3)
	myList1.printData() // 30 20 10
	n5 := &node{data: 100}
	myList1.insertAtMiddle(n5)
	myList1.printData() // 30 20 100 10

	myList2 := linkedList{}
	m1 := &node{data: 1000}
	m2 := &node{data: 2000}
	m3 := &node{data: 3000}
	m4 := &node{data: 4000}
	myList2.prepend(m1)
	myList2.prepend(m2)
	myList2.prepend(m3)
	myList2.prepend(m4)
	myList2.printData() // 4000 3000 2000 1000
	m5 := &node{data: 100}
	myList2.insertAtMiddle(m5)
	myList2.printData() // 4000 3000 100 2000 1000
}
