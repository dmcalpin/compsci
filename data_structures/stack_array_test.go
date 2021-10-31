package data_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(s.T(), &Stack{}, stack)

	stack = NewStack(s.smallItems)
	assert.Equal(s.T(), &Stack{
		items: s.smallItems,
	}, stack)
}

func (s *StackSuite) TestPush() {
	stack := NewStack(nil)

	stack.Push(1)
	assert.Equal(s.T(), &Stack{
		items: []interface{}{1},
	}, stack)

	stack.Push(2)
	assert.Equal(s.T(), &Stack{
		items: []interface{}{1, 2},
	}, stack)

	stack.Push(3)
	assert.Equal(s.T(), &Stack{
		items: []interface{}{1, 2, 3},
	}, stack)
}

func (s *StackSuite) TestPop() {
	stack := NewStack(s.smallItems)

	item := stack.Pop()
	assert.Equal(s.T(), 3, item)

	item = stack.Pop()
	assert.Equal(s.T(), 2, item)

	item = stack.Pop()
	assert.Equal(s.T(), 1, item)

	item = stack.Pop()
	assert.Equal(s.T(), nil, item)
}

func (s *StackSuite) TestPeek() {
	stack := NewStack(s.smallItems)

	item := stack.Peek()
	assert.Equal(s.T(), 3, item)

	item = stack.Peek()
	assert.Equal(s.T(), 3, item)

	stack.Pop()

	item = stack.Peek()
	assert.Equal(s.T(), 2, item)
}

func (s *StackSuite) TestIsEmpty() {
	stack := NewStack([]interface{}{1})
	assert.False(s.T(), stack.IsEmpty())

	stack.Pop()
	assert.True(s.T(), stack.IsEmpty())
}

func BenchmarkPush(t *testing.B) {
	stack := NewStack(nil)

	for i := 0; i < t.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkPop(t *testing.B) {
	stack := NewStack(nil)
	for i := 0; i < 100000; i++ {
		stack.Push(i)
	}

	for i := 0; i < t.N; i++ {
		stack.Pop()
	}
}

func BenchmarkPeek(t *testing.B) {
	stack := NewStack(nil)
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < t.N; i++ {
		stack.Peek()
	}
}

func BenchmarkIsEmpty(t *testing.B) {
	stack := NewStack(nil)
	stack.Push(1)

	for i := 0; i < t.N; i++ {
		stack.IsEmpty()
	}
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}
