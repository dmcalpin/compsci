package data_structures

// A stack is a collection of
// data where the elements are
// operated on in a last-in, first-out
// manner. This stack uses a linked-list
// as it's underlying data structure.

type StackLL struct {
	size int
	head *StackNode
}

type StackNode struct {
	value interface{}
	next  *StackNode
}

// Returns a new StackLL
func NewStackLL() *StackLL {
	return &StackLL{}
}

// Adds a new item to the stack
func (s *StackLL) Push(item interface{}) {
	s.size++

	node := &StackNode{
		value: item,
		next:  s.head,
	}

	s.head = node
}

// Removes and returns the last
// item added to the stack
func (s *StackLL) Pop() interface{} {
	if s.head == nil {
		return nil
	}

	s.size--
	item := s.head
	s.head = s.head.next

	return item.value
}

// Only returns the last item
// in the stack
func (s *StackLL) Peek() interface{} {
	if s.head != nil {
		return s.head.value
	}

	return nil
}

// Returns true if there are no
// itmes in the stack
func (s *StackLL) IsEmpty() bool {
	return s.size == 0
}
