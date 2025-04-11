package main

import (
	"fmt"
)

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func double[T number](num T) T {
	return 2 * num
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintableInt int
type PrintableFloat float64

func (num PrintableInt) String() string {
	return fmt.Sprintf("printing... %d", num)
}

func (num PrintableFloat) String() string {
	return fmt.Sprintf("printing... %f", num)
}

func print[T Printable](in T) {
	fmt.Println(in)
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (ll *LinkedList[T]) Add(val T) {
	n := &Node[T]{Value: val}

	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}
}

func (ll *LinkedList[T]) Insert(val T, i int) {
	n := &Node[T]{Value: val}

	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
		return
	}

	var prev *Node[T]
	for j := range i {
		if j == 0 {
			prev = ll.Head
			continue
		}

		if prev.Next == nil {
			prev.Next = n
			ll.Tail = n
			return
		}

		prev = prev.Next
	}

	n.Next = prev.Next
	prev.Next = n
}

func (ll *LinkedList[T]) Index(val T) int {
	i := 0
	for cur := ll.Head; cur != nil; cur = cur.Next {
		if cur.Value == val {
			return i
		}

		i++
	}

	return -1
}

func Ch8() {
	fmt.Println(double(1))                         // func double(num int) int
	fmt.Println(double(1_000_000_000_000_000_000)) // int literals default to type int, which is int32 or int64 depending on CPU
	var n uint8 = 120
	fmt.Println(double(n)) // func double(num uint8) uint8
	fmt.Println(double(212.32542))

	var x PrintableInt = 2
	print(x)
	var y PrintableFloat = 2.001
	print(y)

	ll := &LinkedList[int]{}
	ll.Insert(99, 3)
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Insert(88, 1)
	ll.Insert(77, 6)

	fmt.Println(ll.Index(ll.Head.Value))
	fmt.Println(ll.Index(99))
	fmt.Println(ll.Index(88))
	fmt.Println(ll.Index(1))
	fmt.Println(ll.Index(2))
	fmt.Println(ll.Index(3))
	fmt.Println(ll.Index(77))
	fmt.Println(ll.Index(ll.Tail.Value))
}
