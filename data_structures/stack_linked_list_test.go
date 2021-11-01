package data_structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackLLSuite struct {
	suite.Suite
}

func (s *StackLLSuite) TestNewStackLL() {
	stack := NewStackLL()
	s.Equal(&StackLL{}, stack)
}

func (s *StackLLSuite) TestPush() {
	stack := NewStackLL()

	stack.Push(1)
	s.Equal(1, stack.head.value)

	stack.Push(2)
	s.Equal(2, stack.head.value)

	stack.Push(3)
	s.Equal(3, stack.head.value)
}

func (s *StackLLSuite) TestPop() {
	stack := NewStackLL()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item := stack.Pop()
	s.Equal(3, item)

	item = stack.Pop()
	s.Equal(2, item)

	item = stack.Pop()
	s.Equal(1, item)

	item = stack.Pop()
	s.Nil(item)
}

func (s *StackLLSuite) TestPeek() {
	stack := NewStackLL()
	stack.Push(2)
	stack.Push(3)

	item := stack.Peek()
	s.Equal(3, item)

	item = stack.Peek()
	s.Equal(3, item)

	stack.Pop()

	item = stack.Peek()
	s.Equal(2, item)
}

func (s *StackLLSuite) TestIsEmpty() {
	stack := NewStackLL()
	stack.Push(1)
	s.False(stack.IsEmpty())

	stack.Pop()
	s.True(stack.IsEmpty())
}

func BenchmarkStackLLPush(t *testing.B) {
	stack := NewStackLL()

	for i := 0; i < t.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackLLPop(t *testing.B) {
	stack := NewStackLL()
	for i := 0; i < 100000; i++ {
		stack.Push(i)
	}

	for i := 0; i < t.N; i++ {
		stack.Pop()
	}
}

func BenchmarkStackLLPeek(t *testing.B) {
	stack := NewStackLL()
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < t.N; i++ {
		stack.Peek()
	}
}

func BenchmarkIsStackLLEmpty(t *testing.B) {
	stack := NewStackLL()
	stack.Push(1)

	for i := 0; i < t.N; i++ {
		stack.IsEmpty()
	}
}

func TestStackLLSuite(t *testing.T) {
	suite.Run(t, new(StackLLSuite))
}
