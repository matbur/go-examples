package main

import (
	"github.com/matbur95/ucisw/structures"
	"fmt"
)

func main() {
	ll := structures.NewLinkedList()
	ll.Print()
	ll.Append(1)
	ll.Print()
	ll.Append(2)
	ll.Print()
	ll.Append(3)
	ll.Print()
	ll.Append(4)
	ll.Print()

	fmt.Println()

	ll.Insert(0, 4)
	ll.Print()
	ll.Insert(0, 3)
	ll.Print()
	ll.Insert(0, 2)
	ll.Print()
	ll.Insert(0, 1)
	ll.Print()

	fmt.Println()

	ll.Insert(1, 0)
	ll.Print()
	ll.Insert(100, 0)
	ll.Print()

	fmt.Println(ll.Index(1))
	fmt.Println(ll.Index(0))
	fmt.Println(ll.Index(2))
	fmt.Println(ll.Index(3))
	fmt.Println(ll.Index(4))
	fmt.Println(ll.Index(5))
}
