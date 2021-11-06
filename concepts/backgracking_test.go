package concepts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SolveNQueenSuite struct {
	suite.Suite
}

func (s *SolveNQueenSuite) Test4By4() {
	solver := NewNQueenSolver(4)
	board := solver.Solve()
	s.Equal([][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 0, 1, 0},
	}, board)
}

func TestSolveNQueenSuite(t *testing.T) {
	suite.Run(t, new(SolveNQueenSuite))
}
