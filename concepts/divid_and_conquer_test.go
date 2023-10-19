package concepts

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type sumSuite struct {
	suite.Suite
}

func (s *sumSuite) TestEmptySlice() {
	s.Require().Equal(0, Sum([]int{}))
}

func (s *sumSuite) TestOneElement() {
	s.Require().Equal(7, Sum([]int{7}))
}

func (s *sumSuite) TestTwoElement() {
	s.Require().Equal(16, Sum([]int{7, 9}))
}

func (s *sumSuite) TestManyElements() {
	s.Require().Equal(1265, Sum([]int{-3, 25, 999, 13, -83, 4, 1, 2, 3, 4, 26, 53, 87, 134}))
}

// BenchmarkSum100-16    	 5134149	       199.6 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSum100(t *testing.B) {
	s := randomSlice(100)
	for i := 0; i < t.N; i++ {
		Sum(s)
	}
}

// BenchmarkSum1000-16    	  614184	      1919 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSum1000(t *testing.B) {
	s := randomSlice(1000)
	for i := 0; i < t.N; i++ {
		Sum(s)
	}
}

// BenchmarkSum10000-16    	   60247	     20495 ns/op	       1 B/op	       0 allocs/op
func BenchmarkSum10000(t *testing.B) {
	s := randomSlice(10000)
	for i := 0; i < t.N; i++ {
		Sum(s)
	}
}

func randomSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}
	return s
}

func TestSumSuite(t *testing.T) {
	suite.Run(t, new(sumSuite))
}
