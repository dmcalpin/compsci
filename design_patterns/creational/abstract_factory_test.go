package creational

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type abstractFactorySuite struct {
	suite.Suite
}

func (s *abstractFactorySuite) TestCreatesASquare() {
	fp := FactoryProducer{}
	f := fp.GetFactory("all")
	square := f.GetShape("square")

	s.Equal("Square", square.GetName())
}

func (s *abstractFactorySuite) TestCreatesCircle() {
	fp := FactoryProducer{}
	f := fp.GetFactory("all")
	square := f.GetShape("circle")

	s.Equal("Circle", square.GetName())
}

func TestAbstractFactorySuite(t *testing.T) {
	suite.Run(t, new(abstractFactorySuite))
}
