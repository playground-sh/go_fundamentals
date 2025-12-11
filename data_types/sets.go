package data_types

import (
	"fmt"
	"strings"
)

// Source - https://www.bytesizego.com/blog/set-in-golang

// Set is a collection of unique elements
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]struct{}),
	}
}

// Add inserts an element into the set
func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s *Set[T]) Contains(value T) bool {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// List returns all elements in the set as a slice
func (s *Set[T]) List() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

// Implement the fmt.Stringer interface on the generic Set type.
// String returns all elements in the set as a string representation of a slice
func (s *Set[T]) String() string {
	// Access the map through s.elements
	// 1. Get the length of the internal map
	keys := make([]string, 0, len(s.elements))

	// 2. Range over the internal map (s.elements)
	for key := range s.elements {
		// Since this specific Set only holds strings (T=string),
		// we can append the key directly without fmt.Sprintf("%v", key).
		// Use %v to safely convert the generic type T to a string
		keys = append(keys, fmt.Sprintf("%v", key))
	}

	// 3. Join the elements
	return "[" + strings.Join(keys, ", ") + "]"
}
