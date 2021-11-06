package data_structures

import (
	"fmt"
)

// A linked list is a simple data structure
// which creates list by having each
// node in the list maintain a reference
// to the next one in the list

type LinkedList struct {
	// how many items are in this list
	size int
	// the most-recently added item
	// in the list
	head *Node
}

type Node struct {
	value interface{}
	next  *Node
}

// Returns a pointer to a LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Inserts a new item at the head of the list
// e.g. 2 -> 1, Insert(3), 3 -> 2 -> 1
func (l *LinkedList) Insert(item interface{}) {
	// create a new node who's "next" is the
	// current head
	n := &Node{
		value: item,
		next:  l.head,
	}

	// then replace the current head with the
	// new node
	l.head = n

	// increase size
	l.size++
}

// Removes the first matching item from the
// list
func (l *LinkedList) Delete(item interface{}) {
	node := l.head
	if node.value == item {
		l.head = node.next
		l.size--
		return
	}

	for node.next != nil {
		if node.next.value == item {
			node.next = node.next.next
			l.size--
			return
		}
		node = node.next
	}
}

// Returns the number of items in the list
func (l *LinkedList) Size() int {
	return l.size
}

// Traverses the list and prints a string
// representation of its contents
func (l *LinkedList) String() string {
	repr := ""
	node := l.head
	for node != nil {
		repr = fmt.Sprintf("%s -> %s", repr, node.value)
		node = node.next
	}
	return repr
}
