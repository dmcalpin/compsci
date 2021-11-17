package behavioral

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type observerSuite struct {
	suite.Suite
}

func (s *observerSuite) TestMultipleObservers() {
	numSubject := &NumberSubject{}

	ob1 := &PrintObserver{}
	ob2 := &RepeatingPrintObserver{}

	numSubject.AddObserver(ob1)
	numSubject.AddObserver(ob2)

	numSubject.SetState(23)

	s.Equal("Print Observer:  23\n", ob1.updatedVal)
	s.Equal("Repeating Print Observer:  23 23 23\n", ob2.updatedVal)

	numSubject.SetState(1)

	s.Equal("Print Observer:  1\n", ob1.updatedVal)
	s.Equal("Repeating Print Observer:  1 1 1\n", ob2.updatedVal)
}

func TestObserverSuite(t *testing.T) {
	suite.Run(t, new(observerSuite))
}
