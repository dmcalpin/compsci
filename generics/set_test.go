package generics

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GSetSuite struct {
	suite.Suite
}

func (s *GSetSuite) TestInitialization() {
	gSet := NewGSet(1, 2, 3)
	elems := gSet.Enumerate()
	sort.Ints(elems)

	s.Require().Equal([]int{1, 2, 3}, elems)
}

func (s *GSetSuite) TestExists() {
	gSet := NewGSet("a", "b", "c")

	s.Require().True(gSet.Exists("a"))
	s.Require().True(gSet.Exists("b"))
	s.Require().True(gSet.Exists("c"))
	s.Require().False(gSet.Exists("d"))
}

func (s *GSetSuite) TestIsEmpty() {
	gSet := NewGSet("a", "b", "c")
	gSet2 := NewGSet[int]()

	s.Require().False(gSet.IsEmpty())
	s.Require().True(gSet2.IsEmpty())
}

func (s *GSetSuite) TestSize() {
	gSet := NewGSet("a", "b", "c")
	gSet2 := NewGSet[int]()

	s.Require().Equal(3, gSet.Size())
	s.Require().Equal(0, gSet2.Size())
}

func (s *GSetSuite) TestUnion() {
	gSet := NewGSet("a", "b", "c")
	gSet2 := NewGSet("d", "e", "f")

	union := gSet.Union(gSet2).Enumerate()
	sort.Strings(union)

	s.Require().Equal([]string{"a", "b", "c", "d", "e", "f"}, union)
}

func (s *GSetSuite) TestAdd() {
	gSet := NewGSet("a", "b", "c")
	gSet.Add("c")
	gSet.Add("d")

	updatedSet := gSet.Enumerate()
	sort.Strings(updatedSet)

	s.Require().Equal([]string{"a", "b", "c", "d"}, updatedSet)
}

func TestGSetSuite(t *testing.T) {
	suite.Run(t, new(GSetSuite))
}
