/* Given a linked list and a value x, partition it such all nodes less than x come before nodes greater than or equal to x. */
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
	fmt.Println()
}

func arrange(l *linkedList, x int) *node {
	var lowerHead, lowerTail, equalHead, equalTail, greaterHead, greaterTail *node = nil, nil, nil, nil, nil, nil
	current := l.headpointer
	for current != nil {
		if current.data == x {
			if equalHead == nil {
				equalHead = current
			}
			if equalTail == nil {
				equalTail = current
			} else {
				equalTail.nextnodepointer = current
				equalTail = current
			}
		} else if current.data < x {
			if lowerHead == nil {
				lowerHead = current
			}
			if lowerTail == nil {
				lowerTail = current
			} else {
				lowerTail.nextnodepointer = current
				lowerTail = current
			}
		} else {
			if greaterHead == nil {
				greaterHead = current
			}
			if greaterTail == nil {
				greaterTail = current
			} else {
				greaterTail.nextnodepointer = current
				greaterTail = current
			}
		}
		current = current.nextnodepointer
	}
	if greaterHead != nil {
		greaterTail.nextnodepointer = nil
	}
	if lowerHead == nil {
		if equalHead == nil {
			return greaterHead
		} else {
			equalTail.nextnodepointer = greaterHead
			return equalHead
		}
	} else if equalHead == nil {
		lowerTail.nextnodepointer = greaterHead
		return lowerHead
	}
	lowerTail.nextnodepointer = equalHead
	equalTail.nextnodepointer = greaterHead
	return lowerHead
}

func main() {
	myList := &linkedList{}
	node1 := &node{data: 30, nextnodepointer: nil}
	node2 := &node{data: 50, nextnodepointer: node1}
	node3 := &node{data: 80, nextnodepointer: node2}
	node4 := &node{data: 70, nextnodepointer: node3}
	node5 := &node{data: 10, nextnodepointer: node4}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	fmt.Printf("Original unarragned linked list\n")
	myList.printListData()
	arrangedHead := arrange(myList, 50)
	fmt.Printf("Arranged linked list\n")
	for arrangedHead != nil {
		fmt.Printf("%d ", arrangedHead.data)
		arrangedHead = arrangedHead.nextnodepointer
	}
}
