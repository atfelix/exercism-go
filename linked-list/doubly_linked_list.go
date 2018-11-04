package linkedlist

import (
	"fmt"
)

var ErrEmptyList = fmt.Errorf("Underflow error:  Doubly-Linked List is empty")

type node struct {
	Val interface{}
	prev *node
	next *node
}

func (n *node) Next() *node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *node) Prev() *node {
	if n == nil {
		return nil
	}
	return n.prev
}

type List struct {
	head *node
	tail *node
}

func NewList(values ...interface{}) *List {
	list := &List{}
	for _, value := range values {
		list.PushBack(value)
	}
	return list
}

func (list *List) First() *node {
	return list.head
}

func (list *List) Last() *node {
	return list.tail
}

func (list *List) PopFront() (interface{}, error) {
	if list.head == nil {
		return interface{}(0), ErrEmptyList
	}
	value := list.head.Val
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}
	return value, nil
}

func (list *List) PopBack() (interface{}, error) {
	if list.tail == nil {
		return interface{}(0), ErrEmptyList
	}
	value := list.tail.Val
	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}
	return value, nil
}

func (list *List) PushFront(element interface{}) {
	if list.head == nil {
		list.head = &node { Val: element }
		list.tail = list.head
	} else {
		newHead := &node { Val: element, next: list.head }
		list.head.prev = newHead
		list.head = newHead
	}
}

func (list *List) PushBack(element interface{}) {
	if list.tail == nil {
		list.tail = &node { Val: element }
		list.head = list.tail
	} else {
		newTail := &node { Val: element, prev: list.tail }
		list.tail.next = newTail
		list.tail = newTail
	}
}

func (list *List) Reverse() {
	for current := list.head; current != nil; current = current.prev {
		current.prev, current.next = current.next, current.prev
	}

	list.head, list.tail = list.tail, list.head
}
