package main

import (
	"fmt"
)

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	k := 5
	l := &List{}

	for i := 0; i < 5; i++ { // Here we are just building a linked list of 15 elements
		l.Push(i)
	}

	fmt.Printf("Before removing %v from the tail :\n", k)
	l.Print()
	l.removekthFromEnd(k)
	fmt.Println("---------------------")
	fmt.Println("After removing : ")
	l.Print()

}

/*
	removekthFromEnd will remove the kth element from the end of a singly linked list
	for example, if we pass 1 and the list has 10 elements, the tail element will be removed
*/
func (l *List) removekthFromEnd(k int) {
	p1 := l.head
	var p2 *Node
	cnt := 0

	for {
		if cnt >= k {
			if cnt == k {
				p2 = l.head
			} else {
				p2 = p2.Next()
			}
		}

		if p1 == l.tail {
			break
		}
		p1 = p1.Next()
		cnt++
	}

	if cnt == k-1 { // we are removing the head
		l.head = l.head.next // reset the head
		return
	}

	if k == 1 { // we are removing the tail
		p2.next = nil
		l.tail = p2
	} else {
		p2.next = p2.Next().Next()
	}
}

func (l *List) kthFromTail(k int) {
	fmt.Println("Welcome to kthFromTail(), l=")
	l.Print()
	fmt.Println("\n")

	p1 := l.head
	cnt := 0
	var p2 *Node

	for {
		if cnt >= k-1 {
			if cnt == k-1 {
				p2 = l.head
			} else {
				p2 = p2.Next()
			}
		}

		if p1 == l.tail {
			break
		}
		p1 = p1.Next()
		cnt++
	}

	fmt.Printf("%d from the tail = %v\n", k, p2)

}

func (l *List) Print() {
	var p *Node
	p = l.head
	fmt.Printf("Head: %v  Tail: %v\n", l.head, l.tail)

	for {
		fmt.Println(p)
		if p == l.tail {
			break
		}
		p = p.Next()
	}

}
