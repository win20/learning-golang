package main

import (
	"fmt"
	"iter"
	"slices"
)

type List2[T any] struct {
	head, tail *element2[T]
}

type element2[T any] struct {
	next *element2[T]
	val T
}

func (lst *List2[T]) Push2(v T) {
	if lst.tail == nil {
		lst.head = &element2[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element2[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List2[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func RangeOverIterators() {
	lst := List2[int]{}
	lst.Push2(10)
	lst.Push2(13)
	lst.Push2(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
