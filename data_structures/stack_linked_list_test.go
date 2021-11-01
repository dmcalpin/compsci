package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StackLLSuite struct {
	suite.Suite
}

func (s *StackLLSuite) TestNewStackLL() {
	stack := NewStackLL()
	assert.Equal(s.T(), &StackLL{}, stack)
}

func (s *StackLLSuite) TestPush() {
	stack := NewStackLL()

	stack.Push(1)
	assert.Equal(s.T(), 1, stack.head.value)

	stack.Push(2)
	assert.Equal(s.T(), 2, stack.head.value)

	stack.Push(3)
	assert.Equal(s.T(), 3, stack.head.value)
}

func (s *StackLLSuite) TestPop() {
	stack := NewStackLL()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item := stack.Pop()
	assert.Equal(s.T(), 3, item)

	item = stack.Pop()
	assert.Equal(s.T(), 2, item)

	item = stack.Pop()
	assert.Equal(s.T(), 1, item)

	item = stack.Pop()
	assert.Equal(s.T(), nil, item)
}

func (s *StackLLSuite) TestPeek() {
	stack := NewStackLL()
	stack.Push(2)
	stack.Push(3)

	item := stack.Peek()
	assert.Equal(s.T(), 3, item)

	item = stack.Peek()
	assert.Equal(s.T(), 3, item)

	stack.Pop()

	item = stack.Peek()
	assert.Equal(s.T(), 2, item)
}

func (s *StackLLSuite) TestIsEmpty() {
	stack := NewStackLL()
	stack.Push(1)
	assert.False(s.T(), stack.IsEmpty())

	stack.Pop()
	assert.True(s.T(), stack.IsEmpty())
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
