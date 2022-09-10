package generics

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GMapSuite struct {
	suite.Suite
}

func (s *GMapSuite) TestIsTypeSafe() {
	m1 := GMap[string, int]{
		"a": 1,
		"b": 2,
	}

	s.Require().Equal(1, m1["a"])
	s.Require().Equal(2, m1["b"])
}

func (s *GMapSuite) TestZeroValueForMissingKeys() {
	m1 := GMap[string, int]{}
	val, ok := m1["nope"]

	s.Require().False(ok)
	s.Require().Equal(0, val)
}

func (s *GMapSuite) TestForEach() {
	m1 := GMap[string, int]{
		"a": 1,
		"b": 2,
	}
	keys := []string{}
	vals := []int{}

	m1.ForEach(func(k string, v int) {
		keys = append(keys, k)
		vals = append(vals, v)
	})

	sort.Strings(keys)
	sort.Ints(vals)
	s.Require().Equal([]string{"a", "b"}, keys)
	s.Require().Equal([]int{1, 2}, vals)
}

func (s *GMapSuite) TestMap() {
	m1 := GMap[string, int]{
		"a": 1,
		"b": 2,
	}

	m1.Map(func(k string, v int) int {
		return v * 2
	}).Map(func(k string, v int) int {
		return v + 1
	})

	s.Require().Equal(m1, GMap[string, int]{"a": 3, "b": 5})
}

func (s *GMapSuite) TestSet() {
	m1 := GMap[string, int]{
		"a": 1,
		"b": 2,
	}

	m1.Set("b", 3).
		Set("c", 4)

	s.Require().Equal(m1, GMap[string, int]{"a": 1, "b": 3, "c": 4})
}

func (s *GMapSuite) TestGet() {
	m1 := GMap[string, int]{
		"a": 1,
		"b": 2,
	}

	val, ok := m1.Get("b")
	s.Require().Equal(2, val)
	s.Require().True(ok)

	val, ok = m1.Get("nope")
	s.Require().Equal(0, val)
	s.Require().False(ok)
}

func (s *GMapSuite) TestDelete() {
	m1 := GMap[string, int]{
		"a": 1,
		"b": 2,
	}

	m1.Delete("b")
	s.Require().Equal(m1, GMap[string, int]{
		"a": 1,
	})
}

func TestGMapSuite(t *testing.T) {
	suite.Run(t, new(GMapSuite))
}
