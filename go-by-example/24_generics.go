package main

import (
	"fmt"
)

func SlicesIndex[S ~[]E, E comparable] (s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T

	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func Generics() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")


    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())

}


// =================================================== //

func sumNumbers(numbers []int) int {
	var result int

	for i := range numbers {
		result += numbers[i]
	}
	return result
}

type Number interface {
	int | int32 | int64 | float32
}

func sumNumbersGeneric[T Number] (numbers []T) T {
	var result T

	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func Generics2() {
	// number1 := []int{1, 2, 3, 4, 5, 6}
	// number2 := []int32{1, 2, 3, 4, 5, 6}
	number3 := []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	res := sumNumbersGeneric(number3)
	fmt.Println(res)
}
