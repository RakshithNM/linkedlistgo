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

func (l *linkedList) deleteLastNodeLinkedlist() {
	toDelete := l.headnode
	for i := 0; i < l.length; i++ {
		if toDelete.nextnodepointer.nextnodepointer == nil {
			toDelete.nextnodepointer = nil
			l.length--
		}
		toDelete = toDelete.nextnodepointer
	}
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
	myList.printData() // 40 30 20 10
	myList.deleteLastNodeLinkedlist()
	myList.printData() // 40 30 20
}
