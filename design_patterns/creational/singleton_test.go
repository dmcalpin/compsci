package creational_test

// Note that in this test file I use the _test
// package. This is so I can test the
// singleton as someone importing the package
// would.

import (
	"testing"

	"github.com/dmcalpin/compsci/design_patterns/creational"
	"github.com/stretchr/testify/suite"
)

type singletonSuite struct {
	suite.Suite
}

func (s *singletonSuite) TestAlwaysReturnsTheSameObject() {
	obj := creational.GetSingletonInstance()
	obj.Foo = "i already exist"

	obj2 := creational.GetSingletonInstance()

	s.Same(obj, obj2)
	s.Equal("i already exist", obj2.Foo)

	obj2.Bar = "set via obj2"

	obj3 := creational.GetSingletonInstance()

	s.Equal("set via obj2", obj.Bar)
	s.Equal("set via obj2", obj2.Bar)
	s.Equal("set via obj2", obj3.Bar)
}

func TestSingletonSuite(t *testing.T) {
	suite.Run(t, new(singletonSuite))
}
