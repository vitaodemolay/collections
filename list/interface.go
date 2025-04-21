package list

import "errors"

/*
 * list is a generic list pack implementation in Go by Vitor Marcos.
 * It provides methods to add and retrieve elements from the list.
 * The list is implemented as a slice of type T, where T is a type parameter.
 * The ReadOnlyList interface provides methods to access the list without modifying it.
 * The List interface extends ReadOnlyList and adds methods to modify the list.
 * The CreateReadOnlyList function creates a new ReadOnlyList instance with the given values.
 * The CreateList function creates a new List instance.
 * The CreateListWithValues function creates a new List instance with the given values.
 */

// ReadOnlyList is an interface that provides methods to access a list without modifying it.
type ReadOnlyList[T any] interface {
	// Get returns the element at the specified index and a boolean indicating if it exists.
	// The parameter index is the index of the element to be retrieved.
	// If the index is out of bounds, it returns a zero value and false.
	Get(index int) (T, bool)

	// GetAll returns all elements in the list.
	// It returns a slice of type T containing all elements.
	GetAll() []T

	// Size returns the number of elements in the list.
	// It returns an integer representing the size of the list.
	Size() int

	// Find returns the first element that matches the given query function and its index.
	// The query function is a predicate that determines if an element matches the search criteria.
	// It returns the first element that matches the query and its index,
	// or return empty value and -1 value at index if not found.
	Find(query func(T) bool) (T, int)
}

// CreateReadOnlyList creates a new ReadOnlyList instance with the given values.
// It returns a ReadOnlyList of type T.
// The parameter values is a slice of type T containing the initial values.
// The function returns a ReadOnlyList[T] and an error if any occurs during the creation.
func CreateReadOnlyList[T any](values []T) (ReadOnlyList[T], error) {
	if values == nil {
		return nil, errors.New("values cannot be empty")
	}
	return newListWithValues(values), nil
}

// List is an interface that extends ReadOnlyList and adds methods to modify the list.
// It provides methods to add, remove, and clear elements from the list.
type List[T any] interface {
	ReadOnlyList[T]

	// Add appends a new element to the list.
	// The parameter value is the element to be added.
	Add(value T)

	// Remove deletes an element from the list by index and returns true if it was removed,
	// and false if not found.
	// The parameter index is the index of the element to be removed.
	Remove(index int) bool

	// Clear removes all elements from the list.
	Clear()
}

// CreateList creates a new List instance.
// It returns a new List of type T empty.
func CreateList[T any]() List[T] {
	return newList[T]()
}

// CreateListWithValues creates a new List instance with the given values.
// It returns a new List of type T with the specified initial values.
// The parameter values is a slice of type T containing the initial values.
// The function returns a List[T] with the specified values and an error if any occurs during the creation.
func CreateListWithValues[T any](values []T) (List[T], error) {
	if values == nil {
		return nil, errors.New("values cannot be empty")
	}
	return newListWithValues(values), nil
}
