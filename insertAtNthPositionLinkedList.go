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

func (l linkedList) printData() {
	toPrint := l.headnode
	for l.length != 0 {
		fmt.Printf("%d\n", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
}

func (l *linkedList) insertAtNthPosition(position int, value int) {
	//fmt.Println(position, value)
	n := node{value, nil}
	toInsertAfter := l.headnode
	//fmt.Println(toInsertAfter)
	for i := 1; i <= l.length; i++ {
		if i == position {
			n.nextnodepointer = toInsertAfter.nextnodepointer
			toInsertAfter.nextnodepointer = &n
			l.length++
		}
		toInsertAfter = toInsertAfter.nextnodepointer
	}
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

func main() {
	myList := linkedList{}
	myList.insertAtTheEnd(10)
	myList.insertAtTheEnd(20)
	myList.insertAtTheEnd(30)
	myList.insertAtTheEnd(40)
	fmt.Println("Before insertion at the 2 index")
	myList.printData()
	myList.insertAtNthPosition(2, 100)
	fmt.Println("After insertion at the 2 index")
	myList.printData()
}
