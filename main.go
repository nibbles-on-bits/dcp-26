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
	l := &List{}

	for i := 0; i < 15; i++ {
		l.Push(i)
	}

	n := l.First()
	println(n.value)

	n = n.Next()
	println(n.value)

	l.kthFromTail(3)
}

/*
   To remove kth from the tail, we find k-1 from the tail and repoint it to k+1 from the tail
*/
func (l *List) removekthFromTail(k int) {
	fmt.Println("Welcome to kthFromTail(), l=")
	l.Print()
	fmt.Println("\n")
	var p *Node
	p = l.head
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

		//fmt.Println(p)
		if p == l.tail {
			break
		}
		p = p.Next()
		cnt++
	}

	fmt.Printf("%d from the tail = %v\n", k, p2)

}

func (l *List) kthFromTail(k int) {
	fmt.Println("Welcome to kthFromTail(), l=")
	l.Print()
	fmt.Println("\n")
	var p *Node
	p = l.head
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

		//fmt.Println(p)
		if p == l.tail {
			break
		}
		p = p.Next()
		cnt++
	}

	fmt.Printf("%d from the tail = %v\n", k, p2)

}

func (l *List) Print() {
	var p *Node
	p = l.head

	for {
		fmt.Println(p)
		if p == l.tail {
			break
		}
		p = p.Next()
	}

}
