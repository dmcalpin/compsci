package compsci

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		search int
		result int
	}{
		{
			"empty returns -1",
			[]int{},
			3,
			-1,
		},
		{
			"single item with correct index",
			[]int{3},
			3,
			0,
		},
		{
			"single item with incorrect index",
			[]int{3},
			99,
			-1,
		},
		{
			"two items with first index",
			[]int{1, 3},
			1,
			0,
		},
		{
			"two items with first index",
			[]int{1, 3},
			3,
			1,
		},
		{
			"two items with no index",
			[]int{1, 3},
			99,
			-1,
		},
		{
			"3 items with first index",
			[]int{1, 2, 3},
			1,
			0,
		},
		{
			"3 items with second index",
			[]int{1, 2, 3},
			2,
			1,
		},
		{
			"3 items with second index",
			[]int{1, 2, 3},
			3,
			2,
		},
		{
			"3 items with no index",
			[]int{1, 2, 3},
			-1,
			-1,
		},
		{
			"many items, high index",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 50, 60, 100, 200, 300, 400},
			400,
			17,
		},
		{
			"many items, no index",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 50, 60, 100, 200, 300, 400},
			100000,
			-1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.result, BinarySearch(tc.input, tc.search))
		})
	}
}

func BenchmarkBinarySearch1000(t *testing.B) {
	numItems := 1000
	items := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = i
	}

	for i := 0; i < t.N; i++ {
		search := rand.Intn(numItems)
		BinarySearch(items, search)
	}
}

func BenchmarkBinarySearch1000000(t *testing.B) {
	numItems := 1000000
	items := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = i
	}

	for i := 0; i < t.N; i++ {
		search := rand.Intn(numItems)
		BinarySearch(items, search)
	}
}

func BenchmarkBinarySearch50000000(t *testing.B) {
	numItems := 500000000
	items := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = i
	}

	for i := 0; i < t.N; i++ {
		search := rand.Intn(numItems)
		BinarySearch(items, search)
	}
}
