// Queue implemented as a singly linked list
package queue

import "fmt"

// Node in queue
type node[T any] struct {
	value T
	next *node[T]
}

// Queue data structure
type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Creates an empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Adds an element to the back of the queue
func (q *Queue[T]) Push(value T) {
	node := &node[T]{
		value: value,
		next: nil,
	}

	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size += 1
}

// Returns the first element of the queue.
// Returns zero value, false if queue is empty
func (q *Queue[T]) Front() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	return q.head.value, true
}

// Removes and returns the first element of the queue.
// Returns zero value, false if queue is empty
func (q *Queue[T]) Pop() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	temp := q.head
	q.head = temp.next
	q.size -= 1

	if q.size == 0 {
		q.tail = q.head
	}

	return temp.value, true
}

// Returns the size of the queue
func (q *Queue[T]) Size() int {
	return q.size
}

// Prints all the elements of the queue
func (q *Queue[T]) Display() {
	if q.size == 0 {
		fmt.Println("Empty Queue")
	}

	temp := q.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}