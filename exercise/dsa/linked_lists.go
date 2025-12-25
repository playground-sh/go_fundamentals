package dsa

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

func (n *Node[T]) GetData() T {
	return n.Data
}

func (n *Node[T]) AddNext(data T) {
	newNode := &Node[T]{Data: data}
	current := n

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func LinkedListDemo() {
	list := Node[int]{Data: 3}
	list.AddNext(5)
	list.AddNext(7)
	list.AddNext(11)

	current := &list
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
