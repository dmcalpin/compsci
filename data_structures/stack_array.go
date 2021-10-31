package data_structures

// A stack is a collection of
// data where the elements are
// operated on in a last-in, first-out
// manner. This stack uses a slice
// as it's underlying data structure.

type Stack struct {
	items []interface{}
}

// Creates a new stack pre-populated
// with items. Pass nil if you don't
// want to pre-populate.
func NewStack(items []interface{}) *Stack {
	return &Stack{
		items: items,
	}
}

// Adds a new item to the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Removes and returns the last
// item added to the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	size := len(s.items)
	lastItem := s.items[size-1]
	s.items = s.items[0 : size-1]

	return lastItem
}

// Only returns the last item
// in the stack
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	lastItem := s.items[len(s.items)-1]
	return lastItem
}

// Returns true if there are no
// itmes in the stack
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
