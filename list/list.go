package list

/*
 * list is a generic list implementation in Go by Vitor Marcos.
 * It provides methods to add and retrieve elements from the list.
 * The list is implemented as a slice of type T, where T is a type parameter.
 */
type list[T any] struct {
	values []T
}

// newList creates a new List instance.
func newList[T any]() *list[T] {
	return &list[T]{values: make([]T, 0)}
}

// newList create a new List instance with starting values.
func newListWithValues[T any](values []T) *list[T] {
	return &list[T]{values: values}
}

// Add appends a new element to the list.
// the parameter value is the element to be added.
func (l *list[T]) Add(value T) {
	l.values = append(l.values, value)
}

// Return one element if it exists, otherwise return a zero value and false.
// the parameter index is the index of the element to be retrieved.
func (l *list[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(l.values) {
		var zero T
		return zero, false
	}
	return l.values[index], true
}

// GetAll returns all elements in the list.
func (l *list[T]) GetAll() []T {
	return l.values
}

// Remove deletes an element from the list by index and return true when it been removed and false when not found.
// the parameter index is the index of the element to be removed.
func (l *list[T]) Remove(index int) bool {
	if index < 0 || index >= len(l.values) {
		return false
	}
	l.values = append(l.values[:index], l.values[index+1:]...)
	return true
}

// Clear removes all elements from the list.
func (l *list[T]) Clear() {
	l.values = make([]T, 0)
}

// Size returns the number of elements in the list.
func (l *list[T]) Size() int {
	return len(l.values)
}

// Find searches for an element in the list and returns it and a index indicating if it was found.
// The query function is a predicate that determines if an element matches the search criteria.
// It returns the first element that matches the query and its index, or return empty value and -1 value at index if not found.
func (l *list[T]) Find(query func(T) bool) (T, int) {
	for idx, value := range l.values {
		if query(value) {
			return value, idx
		}
	}
	var zero T
	return zero, -1
}
