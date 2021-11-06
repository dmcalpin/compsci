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

func (s *SolveNQueenSuite) Test8By8() {
	solver := NewNQueenSolver(8)
	board := solver.Solve()
	s.Equal([][]int{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
	}, board)
}

func (s *SolveNQueenSuite) Test16X16() {
	solver := NewNQueenSolver(16)
	board := solver.Solve()
	s.Equal([][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	}, board)
}

// BenchmarkNQueen4X-8   	 1215984	       960.2 ns/op	     296 B/op	       7 allocs/op
func BenchmarkNQueen4X(t *testing.B) {
	for i := 0; i < t.N; i++ {
		solver := NewNQueenSolver(4)
		solver.Solve()
	}
}

// BenchmarkNQueen8X-8   	   91275	     11438 ns/op	     872 B/op	      12 allocs/op
func BenchmarkNQueen8X(t *testing.B) {
	for i := 0; i < t.N; i++ {
		solver := NewNQueenSolver(8)
		solver.Solve()
	}
}

// BenchmarkNQueen16X-8   	     340	   3448572 ns/op	    2792 B/op	      21 allocs/op
func BenchmarkNQueen16X(t *testing.B) {
	for i := 0; i < t.N; i++ {
		solver := NewNQueenSolver(16)
		solver.Solve()
	}
}

func TestSolveNQueenSuite(t *testing.T) {
	suite.Run(t, new(SolveNQueenSuite))
}
