package double_linked_list

import "fmt"

type node struct {
	value      int
	next, prev *node
}

func newNode(value int) *node {
	return &node{value, nil, nil}
}

type DoubleLinkedList struct {
	head, tail *node
	length     uint
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{nil, nil, 0}
}

func (dll*DoubleLinkedList) String() string {
	s := fmt.Sprintf("%2d: [", dll.length)

	if dll.head == nil {
		return s + "]"
	}

	node := dll.head
	s += fmt.Sprint(node.value)

	for node.next != nil {
		node = node.next
		s += fmt.Sprintf(", %d", node.value)
	}
	return s + "]"
}

func (dll*DoubleLinkedList) Print() {
	fmt.Println(dll)
}

func (dll*DoubleLinkedList) AppendRight(value int) *DoubleLinkedList {
	new_node := newNode(value)
	dll.length += 1

	if dll.head == nil {
		dll.head = new_node
		dll.tail = new_node
		return dll
	}

	dll.tail.next = new_node
	new_node.prev = dll.tail
	dll.tail = new_node

	return dll
}

func (dll*DoubleLinkedList) AppendLeft(value int) *DoubleLinkedList {
	new_node := newNode(value)
	dll.length += 1

	if dll.head == nil {
		dll.head = new_node
		dll.tail = new_node
		return dll
	}

	dll.head.prev = new_node
	new_node.next = dll.head
	dll.head = new_node

	return dll
}

func (dll*DoubleLinkedList) PopRight() int {
	if dll.head == nil {
		panic("pop from empty list")
	}

	dll.length -= 1

	node := dll.tail

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return node.value
	}

	dll.tail = node.prev
	dll.tail.next = nil

	return node.value
}

func (dll*DoubleLinkedList) PopLeft() int {
	if dll.head == nil {
		panic("pop from empty list")
	}

	dll.length -= 1

	node := dll.head

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return node.value
	}

	dll.head = node.next
	dll.head.prev = nil

	return node.value
}
