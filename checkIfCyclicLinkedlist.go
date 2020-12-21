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
	if l.headnode == nil {
		temp = n.nextnodepointer
	}
	l.headnode = n
	n.nextnodepointer = temp
	l.length++
}

func (l linkedList) printData() {
	toPrint := l.headnode
	for l.length != 0 {
		fmt.Printf("%v\t", toPrint)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Println()
}

func (l linkedList) detectLoop() bool {
	slow := l.headnode
	fast := l.headnode
	for l.length != 0 {
		slow = slow.nextnodepointer
		if fast.nextnodepointer != nil {
			fast = fast.nextnodepointer.nextnodepointer
		} else {
			return false
		}
		if slow == nil || fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
		l.length--
	}
	return false
}

func main() {
	var n1, n2, n3, n4, n5, n6 *node
	myList := linkedList{}
	n6 = &node{data: 60}
	n5 = &node{data: 50, nextnodepointer: n6}
	n4 = &node{data: 40, nextnodepointer: n5}
	n3 = &node{data: 30, nextnodepointer: n4}
	n2 = &node{data: 20, nextnodepointer: n3}
	n1 = &node{data: 10, nextnodepointer: n2}
	n6.nextnodepointer = n3

	myList.prepend(n6)
	myList.prepend(n5)
	myList.prepend(n4)
	myList.prepend(n3)
	myList.prepend(n2)
	myList.prepend(n1)
	myList.printData()

	fmt.Println(myList.detectLoop())
}
