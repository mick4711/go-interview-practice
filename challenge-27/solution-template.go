package generics

import (
	"errors"
	"slices"
)

// ErrEmptyCollection is returned when an operation cannot be performed on an empty collection
var ErrEmptyCollection = errors.New("collection is empty")

//
// 1. Generic Pair
//

// Pair represents a generic pair of values of potentially different types
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates a new pair with the given values
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		First:  first,
		Second: second,
	}
}

// Swap returns a new pair with the elements swapped
func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{
		First:  p.Second,
		Second: p.First,
	}
}

//
// 2. Generic Stack
//

// Stack is a generic Last-In-First-Out (LIFO) data structure
type Stack[T any] struct {
	Values []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Values: []T{},
	}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.Values = append(s.Values, value)
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, ErrEmptyCollection
	}

	l := len(s.Values)
	v := s.Values[l-1]

	s.Values = slices.Delete(s.Values, l-1, l)

	return v, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, ErrEmptyCollection
	}

	return s.Values[len(s.Values)-1], nil
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.Values)
}

// IsEmpty returns true if the stack contains no elements
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Values) == 0
}

//
// 3. Generic Queue
//

// Queue is a generic First-In-First-Out (FIFO) data structure
type Queue[T any] struct {
	// TODO: Add necessary fields
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	// TODO: Implement this function
	return nil
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	// TODO: Implement this method
}

// Dequeue removes and returns the front element from the queue
// Returns an error if the queue is empty
func (q *Queue[T]) Dequeue() (T, error) {
	// TODO: Implement this method
	var zero T
	return zero, nil
}

// Front returns the front element without removing it
// Returns an error if the queue is empty
func (q *Queue[T]) Front() (T, error) {
	// TODO: Implement this method
	var zero T
	return zero, nil
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	// TODO: Implement this method
	return 0
}

// IsEmpty returns true if the queue contains no elements
func (q *Queue[T]) IsEmpty() bool {
	// TODO: Implement this method
	return true
}

//
// 4. Generic Set
//

// Set is a generic collection of unique elements
type Set[T comparable] struct {
	// TODO: Add necessary fields
}

// NewSet creates a new empty set
func NewSet[T comparable]() *Set[T] {
	// TODO: Implement this function
	return nil
}

// Add adds an element to the set if it's not already present
func (s *Set[T]) Add(value T) {
	// TODO: Implement this method
}

// Remove removes an element from the set if it exists
func (s *Set[T]) Remove(value T) {
	// TODO: Implement this method
}

// Contains returns true if the set contains the given element
func (s *Set[T]) Contains(value T) bool {
	// TODO: Implement this method
	return false
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	// TODO: Implement this method
	return 0
}

// Elements returns a slice containing all elements in the set
func (s *Set[T]) Elements() []T {
	// TODO: Implement this method
	return nil
}

// Union returns a new set containing all elements from both sets
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	// TODO: Implement this function
	return nil
}

// Intersection returns a new set containing only elements that exist in both sets
func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	// TODO: Implement this function
	return nil
}

// Difference returns a new set with elements in s1 that are not in s2
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	// TODO: Implement this function
	return nil
}

//
// 5. Generic Utility Functions
//

// Filter returns a new slice containing only the elements for which the predicate returns true
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// TODO: Implement this function
	return nil
}

// Map applies a function to each element in a slice and returns a new slice with the results
func Map[T, U any](slice []T, mapper func(T) U) []U {
	// TODO: Implement this function
	return nil
}

// Reduce reduces a slice to a single value by applying a function to each element
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	// TODO: Implement this function
	return initial
}

// Contains returns true if the slice contains the given element
func Contains[T comparable](slice []T, element T) bool {
	// TODO: Implement this function
	return false
}

// FindIndex returns the index of the first occurrence of the given element or -1 if not found
func FindIndex[T comparable](slice []T, element T) int {
	// TODO: Implement this function
	return -1
}

// RemoveDuplicates returns a new slice with duplicate elements removed, preserving order
func RemoveDuplicates[T comparable](slice []T) []T {
	// TODO: Implement this function
	return nil
}
