package generics

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GSliceSuite struct {
	suite.Suite
}

func (s *GSliceSuite) TestIsTypeSafe() {
	l1 := GSlice[int]{1, 2, 3}

	s.Require().Equal(1, l1[0])
	s.Require().Equal(2, l1[1])
	s.Require().Equal(3, l1[2])
}

func (s *GSliceSuite) TestForEach() {
	l1 := GSlice[int]{1, 2}
	vals := []int{}

	l1.ForEach(func(v int) {
		vals = append(vals, v)
	})

	s.Require().Equal([]int{1, 2}, vals)
}

func (s *GSliceSuite) TestMap() {
	l1 := GSlice[int]{1, 2}

	l1.Map(func(v int) int {
		return v * 2
	}).Map(func(v int) int {
		return v + 1
	})
	s.Require().Equal(l1, GSlice[int]{3, 5})
}

func TestGSliceSuite(t *testing.T) {
	suite.Run(t, new(GSliceSuite))
}
