package generics

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GSetSuite struct {
	suite.Suite
}

func (s *GSetSuite) TestNewGSet() {
	gSet := NewGSet(1, 2, 3)
	elems := gSet.Enumerate()
	sort.Ints(elems)

	s.Require().Equal([]int{1, 2, 3}, elems)
}

func (s *GSetSuite) TestUnion() {
	gSet := NewGSet("a", "b", "c")
	gSet2 := NewGSet("d", "e", "f")

	union := gSet.Union(gSet2).Enumerate()
	sort.Strings(union)

	s.Require().Equal([]string{"a", "b", "c", "d", "e", "f"}, union)
}

func (s *GSetSuite) TestIntersection() {
	// gSet > gSet2
	gSet := NewGSet("a", "b", "c", "q", "r", "t")
	gSet2 := NewGSet("b", "c", "d")

	union := gSet.Intersection(gSet2).Enumerate()
	sort.Strings(union)

	s.Require().Equal([]string{"b", "c"}, union)

	// gSet < gSet2
	gSet = NewGSet("b", "c", "d")
	gSet2 = NewGSet("a", "b", "c", "q", "r", "t")

	union = gSet.Intersection(gSet2).Enumerate()
	sort.Strings(union)

	s.Require().Equal([]string{"b", "c"}, union)
}

func (s *GSetSuite) TestDifference() {
	// gSet > gSet2
	gSet := NewGSet("a", "b", "c", "q", "r", "t")
	gSet2 := NewGSet("b", "c", "d")

	union := gSet.Difference(gSet2).Enumerate()
	sort.Strings(union)

	s.Require().Equal([]string{"a", "q", "r", "t"}, union)

	// gSet < gSet2
	gSet = NewGSet("b", "c", "d")
	gSet2 = NewGSet("a", "b", "c", "q", "r", "t")

	union = gSet.Difference(gSet2).Enumerate()
	sort.Strings(union)

	s.Require().Equal([]string{"d"}, union)
}

func (s *GSetSuite) TestIsSubsetOf() {
	gSet := NewGSet("a", "b", "c", "d", "r", "t")
	gSet2 := NewGSet("b", "c", "r")
	gSet3 := NewGSet("a", "b", "z")

	s.Require().False(gSet.IsSubsetOf(gSet2))
	s.Require().True(gSet2.IsSubsetOf(gSet))
	s.Require().False(gSet3.IsSubsetOf(gSet))
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

func (s *GSetSuite) TestEnumerate() {
	gSet := NewGSet("a", "b", "c")

	enumeration := gSet.Enumerate()
	sort.Strings(enumeration)

	s.Require().Equal([]string{"a", "b", "c"}, enumeration)
}

func (s *GSetSuite) TestAdd() {
	gSet := NewGSet("a", "b", "c")
	gSet.Add("c")
	gSet.Add("d")

	updatedSet := gSet.Enumerate()
	sort.Strings(updatedSet)

	s.Require().Equal([]string{"a", "b", "c", "d"}, updatedSet)
}

func (s *GSetSuite) TestRemove() {
	gSet := NewGSet("a", "b", "c")
	gSet.Remove("b")
	gSet.Remove("z")

	updatedSet := gSet.Enumerate()
	sort.Strings(updatedSet)

	s.Require().Equal([]string{"a", "c"}, updatedSet)
}

func (s *GSetSuite) TestPop() {
	gSet := NewGSet("a", "b", "c")
	enumeration := gSet.Enumerate()

	popped := gSet.Pop()
	s.Require().NotNil(popped)
	s.Require().Contains(enumeration, *popped)
	s.Require().Equal(2, gSet.Size())

	popped = gSet.Pop()
	s.Require().Contains(enumeration, *popped)
	s.Require().Equal(1, gSet.Size())

	popped = gSet.Pop()
	s.Require().Contains(enumeration, *popped)
	s.Require().Equal(0, gSet.Size())

	popped = gSet.Pop()
	s.Require().Nil(popped)
	s.Require().Equal(0, gSet.Size())
}

func (s *GSetSuite) TestMap() {
	gSet := NewGSet("a", "b", "c")
	gSet.Map(func(elem string) string {
		return elem + "X"
	})

	enumeration := gSet.Enumerate()
	sort.Strings(enumeration)
	s.Require().Equal([]string{"aX", "bX", "cX"}, enumeration)
}

func (s *GSetSuite) TestFilter() {
	gSet := NewGSet("a", "b", "c", "d")
	gSet.Filter(func(elem string) bool {
		return elem == "a" || elem == "d"
	})

	enumeration := gSet.Enumerate()
	sort.Strings(enumeration)
	s.Require().Equal([]string{"b", "c"}, enumeration)
}

func (s *GSetSuite) TestClear() {
	gSet := NewGSet("a", "b", "c", "d")
	gSet.Clear()

	enumeration := gSet.Enumerate()
	s.Require().Equal([]string{}, enumeration)
}

func (s *GSetSuite) TestEqual() {
	gSet := NewGSet("a", "b", "c", "d")
	gSet2 := NewGSet("a", "b", "c")
	gSet3 := NewGSet("a", "b", "c", "d")
	gSet4 := NewGSet("b", "c", "d", "e")

	s.Require().False(gSet.Equal(gSet2))
	s.Require().True(gSet.Equal(gSet3))
	s.Require().False(gSet.Equal(gSet4))
}

func TestGSetSuite(t *testing.T) {
	suite.Run(t, new(GSetSuite))
}
