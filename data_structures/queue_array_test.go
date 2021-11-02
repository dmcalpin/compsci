package data_structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type QueueSuite struct {
	suite.Suite
	smallItems  []interface{}
	mediumItems []interface{}
}

func (s *QueueSuite) SetupTest() {
	s.smallItems = []interface{}{1, 2, 3}
	s.mediumItems = []interface{}{1, 2, 3, 4, 5}
}

func (s *QueueSuite) TestNewQueue() {
	queue := NewQueue(nil)
	s.Equal(&Queue{}, queue)

	queue = NewQueue(s.smallItems)
	s.Equal(&Queue{
		items: s.smallItems,
	}, queue)
}

func (s *QueueSuite) TestPush() {
	queue := NewQueue(nil)

	queue.Push(1)
	s.Equal(&Queue{
		items: []interface{}{1},
	}, queue)

	queue.Push(2)
	s.Equal(&Queue{
		items: []interface{}{1, 2},
	}, queue)

	queue.Push(3)
	s.Equal(&Queue{
		items: []interface{}{1, 2, 3},
	}, queue)
}

func (s *QueueSuite) TestPop() {
	queue := NewQueue(s.smallItems)

	item := queue.Pop()
	s.Equal(1, item)

	item = queue.Pop()
	s.Equal(2, item)

	item = queue.Pop()
	s.Equal(3, item)

	item = queue.Pop()
	s.Equal(nil, item)
}

func (s *QueueSuite) TestPeek() {
	queue := NewQueue(s.smallItems)

	item := queue.Peek()
	s.Equal(1, item)

	item = queue.Peek()
	s.Equal(1, item)

	queue.Pop()

	item = queue.Peek()
	s.Equal(2, item)
}

func (s *QueueSuite) TestIsEmpty() {
	queue := NewQueue([]interface{}{1})
	s.False(queue.IsEmpty())

	queue.Pop()
	s.True(queue.IsEmpty())
}

func BenchmarkQueuePush(t *testing.B) {
	queue := NewQueue(nil)

	for i := 0; i < t.N; i++ {
		queue.Push(i)
	}
}

func BenchmarkQueuePop(t *testing.B) {
	queue := NewQueue(nil)
	for i := 0; i < 100000; i++ {
		queue.Push(i)
	}

	for i := 0; i < t.N; i++ {
		queue.Pop()
	}
}

func BenchmarkQueuePeek(t *testing.B) {
	queue := NewQueue(nil)
	for i := 0; i < 1000; i++ {
		queue.Push(i)
	}

	for i := 0; i < t.N; i++ {
		queue.Peek()
	}
}

func BenchmarkQueueIsEmpty(t *testing.B) {
	queue := NewQueue(nil)
	queue.Push(1)

	for i := 0; i < t.N; i++ {
		queue.IsEmpty()
	}
}

func TestQueueSuite(t *testing.T) {
	suite.Run(t, new(QueueSuite))
}
