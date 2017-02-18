package linked_list

import "fmt"

type node struct {
	value int
	next  *node
}

func newNode(value int) *node {
	return &node{value, nil}
}

type LinkedList struct {
	head   *node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

func (ll *LinkedList) String() string {
	s := fmt.Sprintf("%2d: [", ll.length)

	if ll.head == nil {
		return s + "]"
	}

	node := ll.head
	s += fmt.Sprint(node.value)

	for node.next != nil {
		node = node.next
		s += fmt.Sprintf(", %d", node.value)
	}
	return s + "]"
}

func (ll *LinkedList) Print() {
	fmt.Println(ll.String())
}

func (ll *LinkedList) Insert(index, value int) *LinkedList {
	if index == -1 {
		index = ll.length
	} else if index <= -ll.length {
		index = 0
	} else if index < 0 {
		index = ll.length - index - 1
	}

	new_node := newNode(value)
	ll.length++

	node := ll.head

	if ll.head == nil || index == 0 {
		ll.head = new_node
		new_node.next = node
		return ll
	}

	for i := 1; i < index && node.next != nil; i++ {
		node = node.next
	}
	new_node.next = node.next
	node.next = new_node
	return ll
}

func (ll*LinkedList) Append(value int) *LinkedList {
	return ll.Insert(-1, value)
}

func (ll *LinkedList) Index(value int) int {
	const notFoundIndex = -1

	if ll.head == nil {
		return notFoundIndex
	}

	node := ll.head

	if node.value == value {
		return 0
	}

	for i := 1; node.next != nil; i++ {
		node = node.next
		if node.value == value {
			return i
		}
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
