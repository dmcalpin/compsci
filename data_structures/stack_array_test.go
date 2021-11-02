package data_structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
	smallItems  []interface{}
	mediumItems []interface{}
}

func (s *StackSuite) SetupTest() {
	s.smallItems = []interface{}{1, 2, 3}
	s.mediumItems = []interface{}{1, 2, 3, 4, 5}
}

func (s *StackSuite) TestNewStack() {
	stack := NewStack(nil)
	s.Equal(&Stack{}, stack)

	stack = NewStack(s.smallItems)
	s.Equal(&Stack{
		items: s.smallItems,
	}, stack)
}

func (s *StackSuite) TestPush() {
	stack := NewStack(nil)

	stack.Push(1)
	s.Equal(&Stack{
		items: []interface{}{1},
	}, stack)

	stack.Push(2)
	s.Equal(&Stack{
		items: []interface{}{1, 2},
	}, stack)

	stack.Push(3)
	s.Equal(&Stack{
		items: []interface{}{1, 2, 3},
	}, stack)
}

func (s *StackSuite) TestPop() {
	stack := NewStack(s.smallItems)

	item := stack.Pop()
	s.Equal(3, item)

	item = stack.Pop()
	s.Equal(2, item)

	item = stack.Pop()
	s.Equal(1, item)

	item = stack.Pop()
	s.Equal(nil, item)
}

func (s *StackSuite) TestPeek() {
	stack := NewStack(s.smallItems)

	item := stack.Peek()
	s.Equal(3, item)

	item = stack.Peek()
	s.Equal(3, item)

	stack.Pop()

	item = stack.Peek()
	s.Equal(2, item)
}

func (s *StackSuite) TestIsEmpty() {
	stack := NewStack([]interface{}{1})
	s.False(stack.IsEmpty())

	stack.Pop()
	s.True(stack.IsEmpty())
}

func BenchmarkStackPush(t *testing.B) {
	stack := NewStack(nil)

	for i := 0; i < t.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(t *testing.B) {
	stack := NewStack(nil)
	for i := 0; i < 100000; i++ {
		stack.Push(i)
	}

	for i := 0; i < t.N; i++ {
		stack.Pop()
	}
}

func BenchmarkStackPeek(t *testing.B) {
	stack := NewStack(nil)
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < t.N; i++ {
		stack.Peek()
	}
}

func BenchmarkStackIsEmpty(t *testing.B) {
	stack := NewStack(nil)
	stack.Push(1)

	for i := 0; i < t.N; i++ {
		stack.IsEmpty()
	}
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}
