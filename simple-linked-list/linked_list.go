package linkedlist

import (
	"fmt"
)

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(slice []int) *List {
	if slice == nil || len(slice) == 0 {
		return &List{}
	}
	return &List{head: NewElement(slice), size: len(slice)}
}

func NewElement(slice []int) *Element {
	if slice == nil || len(slice) == 0 {
		return nil
	}
	element := &Element{data: slice[len(slice) - 1], next: NewElement(slice[:len(slice) - 1])}
	return element
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Array() []int {
	size := list.Size()
	slice := make([]int, size)
	head := list.head

	for head != nil {
		slice[size - 1] = head.data
		size--
		head = head.next
	}

	return slice
}

func (list *List) Pop() (int, error) {
	if list.Size() == 0 {
		return -1, fmt.Errorf("Cannot Pop:  List is empty")
	}
	data := list.head.data
	*list = List{head: list.head.next, size: list.size - 1}
	return data, nil
}

func (list *List) Push(element int) {
	if list.Size() == 0 {
		*list = List{head: &Element{data: element, next: nil}, size: 1}
	} else {
		*list = List{head: &Element{data: element, next: list.head}, size: list.size + 1}
	}
}

func (list *List) Reverse() *List {
	reverse := New(nil)

	for list.head != nil {
		data, err := list.Pop()
		if err != nil {
			panic("The head is non-empty:  We should never get here")
		}
		reverse.Push(data)
	}

	return reverse
}