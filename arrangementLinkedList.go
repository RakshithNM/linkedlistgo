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
			// Initially assign current to the head and tail
			if equalHead == nil {
				equalHead = current
			}
			if equalTail == nil {
				equalTail = current
			} else {
				// Create a list with the equal values
				equalTail.nextnodepointer = current
				equalTail = current
			}
		} else if current.data < x {
			// Initially assign current to the head and tail
			if lowerHead == nil {
				lowerHead = current
			}
			if lowerTail == nil {
				lowerTail = current
			} else {
				// Create a list with the lower values - not sorted
				lowerTail.nextnodepointer = current
				lowerTail = current
			}
		} else {
			// Initially assign current to the head and tail
			if greaterHead == nil {
				greaterHead = current
			}
			if greaterTail == nil {
				greaterTail = current
			} else {
				// Create a list with the greater values - not sorted
				greaterTail.nextnodepointer = current
				greaterTail = current
			}
		}
		current = current.nextnodepointer
	}
	// Set the tail of the linkedList
	if greaterHead != nil {
		greaterTail.nextnodepointer = nil
	}
	// Connecting all three lists begins
	// Check if lowerList is not existent
	if lowerHead == nil {
		// If lowerList and  equalList is not present, return greaterHead
		if equalHead == nil {
			return greaterHead
		} else {
			// If lowerList not present and equalList is present, connect equalList to greaterList
			equalTail.nextnodepointer = greaterHead
			return equalHead
		}
		// If equalList is not present connect lowerList to greaterList
	} else if equalHead == nil {
		lowerTail.nextnodepointer = greaterHead
		return lowerHead
	}
	// All present, connect all
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
