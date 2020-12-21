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

func (l *linkedList) insertAtTheEnd(val int) {
	n := &node{data: val, nextnodepointer: nil}
	if l.headnode == nil {
		l.headnode = n
		l.length++
		return
	}
	toInsert := l.headnode
	for i := 0; i < l.length; i++ {
		if toInsert.nextnodepointer == nil {
			toInsert.nextnodepointer = n
			l.length++
			return
		}
		toInsert = toInsert.nextnodepointer
	}
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

func (l linkedList) getThirdFromEnd() int {
	ptr := l.headnode
	thirdIndexFromLast := l.length - 3
	var thirdNodeFromEnd *node
	for i := 0; i <= thirdIndexFromLast; i++ {
		thirdNodeFromEnd = ptr
		ptr = ptr.nextnodepointer
	}
	return thirdNodeFromEnd.data
}

func main() {
	fmt.Println("vim-go")
	myList := linkedList{}
	myList.insertAtTheEnd(10)
	myList.insertAtTheEnd(20)
	myList.insertAtTheEnd(30)
	myList.insertAtTheEnd(40)
	myList.insertAtTheEnd(50)
	myList.insertAtTheEnd(60)

	myList.printData() // 10 20 30

	fmt.Println("The third node data from the end is: ", myList.getThirdFromEnd())
}
