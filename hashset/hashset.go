/*
Hash set implementation.
Performs Insertion, Deletion and Lookup operations in O(1) time.
*/
package hashset

import "fmt"

// Hash set data structure
type HashSet[T comparable] struct {
	set map[T]struct{}
}

// Creates an empty hash set
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		set: make(map[T]struct{}),
	}
}

// Adds an element to the set
func (s *HashSet[T]) Add(value T) {
	s.set[value] = struct{}{}
}

// Returns true if the set contains the element.
// Returns false otherwise.
func (s *HashSet[T]) Contains(value T) bool {
	_, exists := s.set[value]
	return exists
}

// Removes the element from the set and returns true.
// Returns false if the element doesn't exist int the set.
func (s *HashSet[T]) Remove(value T) bool {
	if _, exists := s.set[value]; !exists {
		return false
	}

	delete(s.set, value)
	
	return true
}

// Returns the size of the set
func (s *HashSet[T]) Size() int {
	return len(s.set)
}

// Prints the elements of the set
func (s *HashSet[T]) Display() {
	if s.Size() == 0 {
		fmt.Println("Empty Set")
		return
	}

	for k, _ := range s.set {
		fmt.Println(k)
	}
}