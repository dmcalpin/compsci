package compsci

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		result []int
	}{
		{
			"emptys",
			[]int{},
			[]int{},
		},
		{
			"single item",
			[]int{3},
			[]int{3},
		},
		{
			"two items",
			[]int{3, 1},
			[]int{1, 3},
		},
		{
			"duplicate items",
			[]int{2, 1, 2, 1},
			[]int{1, 1, 2, 2},
		},
		{
			"many random items",
			[]int{2, 1, 30, 3, 5, 100, 6, 7, 20, 300, 8, 9, 4, 10, 400, 50, 60, 200},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 50, 60, 100, 200, 300, 400},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.result, SelectionSort(tc.input))
		})
	}
}

func randomSlice(n int) []int {
	numItems := n
	items := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = rand.Intn(numItems)
	}
	return items
}

func BenchmarkSelectionSort100(t *testing.B) {
	items := randomSlice(100)

	for i := 0; i < t.N; i++ {
		SelectionSort(items)
	}
}

func BenchmarkSelectionSort1000(t *testing.B) {
	items := randomSlice(1000)

	for i := 0; i < t.N; i++ {
		SelectionSort(items)
	}
}

func BenchmarkSelectionSort10000(t *testing.B) {
	items := randomSlice(10000)

	for i := 0; i < t.N; i++ {
		SelectionSort(items)
	}
}
