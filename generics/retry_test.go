package generics

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type RetrySuite struct {
	suite.Suite
}

func (s *RetrySuite) TestSuccess() {
	tries := 0

	err := Retry(func() error {
		if tries < 3 {
			fmt.Println("dave", tries)
			tries++
			return fmt.Errorf("even number!: %d", tries)
		}
		return nil
	}, 5, 1*time.Millisecond)

	s.Require().NoError(err)
	s.Require().Equal(3, tries)
}

func (s *RetrySuite) TestErrorsOut() {
	err := Retry(func() error {
		return errors.New("will fail")
	}, 3, 1*time.Millisecond)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "will fail")
}

func (s *RetrySuite) TestClosureRetries() {
	str := ""
	err := Retry(func() error {
		if len(str) < 100 {
			str += "trying "
			return errors.New("not long enough")
		}

		return nil
	}, 3, 1*time.Second)

	s.Require().Error(err)
	s.Require().Equal("trying trying trying ", str)
}

func TestRetrySuite(t *testing.T) {
	suite.Run(t, new(RetrySuite))
}
