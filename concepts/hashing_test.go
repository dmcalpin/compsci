package concepts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HashingSuite struct {
	suite.Suite
}

func (s *HashingSuite) TestSimpleHash() {
	s.Equal(123, SimpleHash("Howdy", 200))
	s.Equal(189, SimpleHash("Alaska", 200))
}

func TestHashingSuite(t *testing.T) {
	suite.Run(t, new(HashingSuite))
}
