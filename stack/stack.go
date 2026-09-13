/*
Stack implementation as a singly linked list.
Performs Push, Pop and Top operations in O(1) time.
*/
package stack

import "fmt"

// Node in stack
type node[T any] struct {
	value T
	next *node[T]
}

// Stack data structure
type Stack[T any] struct {
	head *node[T]
	size int
}

// Creates an empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head: nil,
		size: 0,
	}
}

// Adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	node := &node[T]{
		value: value,
		next: nil,
	}

	if s.size == 0 {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.size += 1
}

// Returns the top element of the stack.
// Returns zero-value, false if stack is empty.
func (s *Stack[T]) Top() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	return s.head.value, true
}

// Removes and returns the top element of the stack.
// Returns zero-value, false if stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	temp := s.head
	s.head = s.head.next
	s.size -= 1

	return temp.value, true
}

// Returns the size of the stack
func (s *Stack[T]) Size() int {
	return s.size
}

// Prints the elements of the stack from top to bottom
func (s *Stack[T]) Display() {
	if s.size == 0 {
		fmt.Println("Empty Stack")
		return
	}

	temp := s.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}