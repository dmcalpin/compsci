package data_structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type linkedListSuite struct {
	suite.Suite

	ll *LinkedList
}

func (s *linkedListSuite) SetupTest() {
	s.ll = NewLinkedList()

	s.ll.Insert("item 1")
	s.ll.Insert("item 2")
	s.ll.Insert("item 3")
	s.ll.Insert("item 4")
	s.ll.Insert("item 5")
}

func (s *linkedListSuite) TestNewLinkedList() {
	s.Equal(&LinkedList{}, NewLinkedList())
}

func (s *linkedListSuite) TestInsert() {
	ll := NewLinkedList()

	ll.Insert("item 1")
	ll.Insert("item 2")

	s.Equal(" -> item 2 -> item 1", ll.String())
	s.Equal("item 2", ll.head.value)
	s.Equal(2, ll.Size())
}

func (s *linkedListSuite) TestDeleteMiddle() {
	s.ll.Delete("item 3")
	s.Equal(" -> item 5 -> item 4 -> item 2 -> item 1", s.ll.String())
	s.Equal(4, s.ll.Size())
}

func (s *linkedListSuite) TestDeleteHead() {
	s.ll.Delete("item 5")
	s.Equal(" -> item 4 -> item 3 -> item 2 -> item 1", s.ll.String())
	s.Equal(4, s.ll.Size())
}

func (s *linkedListSuite) TestDeleteTail() {
	s.ll.Delete("item 1")
	s.Equal(" -> item 5 -> item 4 -> item 3 -> item 2", s.ll.String())
	s.Equal(4, s.ll.Size())
}

func TestLinkedListSuite(t *testing.T) {
	suite.Run(t, new(linkedListSuite))
}
