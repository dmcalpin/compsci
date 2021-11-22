package structural

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type decoratorSuite struct {
	suite.Suite
}

func (s *decoratorSuite) TestPlainBurger() {
	b := NewBurger()
	s.Equal(2.99, b.GetCost())
	s.Equal("A Hamburger", b.GetDescription())
}

func (s *decoratorSuite) TestCheeseburgerWithKatchup() {
	b := AddKatchup(AddCheese(NewBurger()))

	s.Equal(3.74, b.GetCost())
	s.Equal("A Hamburger, with Cheese, with Katchup", b.GetDescription())
}

func (s *decoratorSuite) TestPlainFries() {
	f := NewFries()
	s.Equal(1.99, f.GetCost())
	s.Equal("French Fries", f.GetDescription())
}

func (s *decoratorSuite) TestFriesWithKatchup() {
	b := AddKatchup(NewFries())

	s.Equal(2.24, b.GetCost())
	s.Equal("French Fries, with Katchup", b.GetDescription())
}

func TestDecoratorSuite(t *testing.T) {
	suite.Run(t, new(decoratorSuite))
}
