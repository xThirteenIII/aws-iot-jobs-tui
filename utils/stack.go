package utils

// Stack is a LIFO data structure
type Stack[T any] struct {
	items []T
	length int
}

// NewStack returns an empty stuck.
func NewStack[T any]() *Stack[T] {
	return  &Stack[T]{
		items: []T{},
		length: 0,
	}
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.length++
}

// Pop removes last item from the stack and returns it.
func (s *Stack[T]) Pop(item T) T{

	// If stack is empty return empty T variable.
	if s.length == 0 {
		var t T
		return t
	}

	// Remove last item
	topElement := s.items[s.length-1]
	s.items = s.items[:s.length-1]
	s.length-=1

	return topElement
}

// Peek returns the top item of the stack, without removing it.
func (s *Stack[T]) Pee(item T) T{
		
	// If stack is empty return empty T variable.
	if s.length == 0 {
		var t T
		return t
	}
	return s.items[s.length - 1]
}

func (s *Stack[T]) Len() int {
	return s.length
}

// Clear removes all items from the stack
func (s *Stack[T]) Clear() {
	s.items = nil
	s.length = 0
}

