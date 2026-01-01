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

func (n *Node[T]) Print() {
	current := n
	for current != nil {
		fmt.Print(current.Data)
		current = current.Next
		if current != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func (n *Node[T]) Length() int {
	count := 0
	current := n

	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func LinkedListDemo() {
	list := Node[int]{Data: 3}
	list.AddNext(5)
	list.AddNext(7)
	list.AddNext(11)

	list.Print()
	fmt.Printf("List has %d nodes in total\n", list.Length())
}
