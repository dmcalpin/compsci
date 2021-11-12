package creational

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type builerSuite struct {
	suite.Suite
}

func (s *builerSuite) TestCreatesAPC() {
	cb := &ComputerBuilder{}
	pc := cb.makePC()

	s.Equal(
		`- Intel i7: 4.50
- Coolmax Pro: Full
- MSI Gamer Pro: Mini ITX
- Crucial 8x4: 32GB
- Samsung Evo 970: 500GB
`,
		pc.ListParts(),
	)

	s.Equal("959.05", fmt.Sprintf("%.02f", pc.GetTotalPrice()))
}

func (s *builerSuite) TestCreatesAMac() {
	cb := &ComputerBuilder{}
	mac := cb.makeMac()

	s.Equal(
		`- 8-Core Processor: 4.50
- Macbook Pro: 14"
- Apple Memory: 16GB
- Apple SSD: 500GB
`,
		mac.ListParts(),
	)

	s.Equal("1700.00", fmt.Sprintf("%.02f", mac.GetTotalPrice()))
}

func TestBuilerSuite(t *testing.T) {
	suite.Run(t, new(builerSuite))
}
