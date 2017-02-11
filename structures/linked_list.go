package structures

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Print() {
	fmt.Printf("%2d: {", ll.length)
	defer fmt.Println("}")

	if ll.head == nil {
		return
	}

	node := ll.head
	fmt.Print(node.value)

	for node.next != nil {
		node = node.next
		fmt.Printf(", %d", node.value)
	}
}

func (ll *LinkedList) Append(value int) {
	new_node := NewNode(value)
	ll.length++

	if ll.head == nil {
		ll.head = new_node
		return
	}

	node := ll.head
	for node.next != nil {
		node = node.next
	}
	node.next = new_node
}

func (ll *LinkedList) Insert(index, value int) {
	new_node := NewNode(value)
	ll.length++

	node := ll.head

	if ll.head == nil || index == 0 {
		ll.head = new_node
		new_node.next = node
		return
	}

	for i := 1; i < index && node.next != nil; i++ {
		node = node.next
	}
	new_node.next = node.next
	node.next = new_node
}

func (ll *LinkedList) Index(value int) int {
	const notFoundIndex = -1

	if ll.head == nil {
		return notFoundIndex
	}

	node := ll.head
	for i := 0; node.next != nil; i++ {
		if node.value == value {
			return i
		}
		node = node.next
	}
	return notFoundIndex
}

func (ll *LinkedList) Remove(value int) int {
	const notFoundIndex = -1

	if ll.head == nil {
		return notFoundIndex
	}

	node := ll.head

	if node.value == value {
		ll.head = node.next
		ll.length--
		return 0
	}

	if node.next.value == value {
		node.next = node.next.next
		ll.length--
		return 1
	}

	for i := 1; node.next.next != nil; i++ {
		if node.next.value == value {
			node.next = node.next.next
			ll.length--
			return i
		}
		node = node.next
	}

	if node.next.value == value {
		node.next = nil
		ll.length--
		return ll.length
	}

	return notFoundIndex
}

func (ll *LinkedList) Pop(index int) int {
	node := ll.head

	if ll.head == nil {
		panic("pop from empty list")
	}

	if index >= ll.length {
		panic("pop index out of range")
	}

	if index == 0 {
		ll.head = node.next
		ll.length--
		return node.value
	}

	for i := 1; i < index; i++ {
		node = node.next
	}

	node.next = node.next.next
	ll.length--

	return node.value
}
