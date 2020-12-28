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
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.nextnodepointer
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) prepend(n *node) {
	temp := l.headnode
	l.headnode = n
	n.nextnodepointer = temp
	l.length++
}

func (l *linkedList) swap(ptr1 *node, ptr2 *node) {
	temp := ptr1.data
	ptr1.data = ptr2.data
	ptr2.data = temp
}

func (l *linkedList) sort() {
	for i := 0; i < l.length-1; i++ {
		swapped := false
		first := l.headnode
		for j := 0; j < l.length-i-1 && first.nextnodepointer != nil; j++ {
			if first.data > first.nextnodepointer.data {
				l.swap(first, first.nextnodepointer)
				swapped = true
			}
			first = first.nextnodepointer
		}
		if swapped == false {
			break
		}
	}
}

func main() {
	myList := linkedList{}
	n0 := &node{data: 1}
	n1 := &node{data: 10}
	n2 := &node{data: 20}
	n4 := &node{data: 40}
	n3 := &node{data: 30}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n4)
	myList.prepend(n0)
	myList.prepend(n3)
	myList.printData() // 30 1 40 20 10
	myList.sort()
	myList.printData() // 1 10 20 30 40
}
