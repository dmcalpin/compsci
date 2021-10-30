package concepts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveSum(t *testing.T) {
	assert.Equal(t, 0, RecursiveSum())
	assert.Equal(t, 15, RecursiveSum(1, 2, 3, 4, 5))
	assert.Equal(t, 28416, RecursiveSum(312, 324, 6390, 21390))
}

func BenchmarkHRecursiveSum(b *testing.B) {
	numSets := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986},
		{132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986, 54, 65, 76, 876, 97, 14, 74, 986, 54, 6},
		{132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986, 54, 65, 76, 876, 97, 14, 74, 986, 54, 6,
			132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986, 54, 65, 76, 876, 97, 14, 74, 986, 54, 6},
	}
	for i := 0; i < b.N; i++ {
		RecursiveSum(numSets[i%4]...)
	}
}

func BenchmarkTraditionalSum(b *testing.B) {
	numSets := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986},
		{132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986, 54, 65, 76, 876, 97, 14, 74, 986, 54, 6},
		{132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986, 54, 65, 76, 876, 97, 14, 74, 986, 54, 6,
			132, 432, 54, 65, 76, 876, 978, 76, 96, 14, 74, 986, 54, 65, 76, 876, 97, 14, 74, 986, 54, 6},
	}
	for i := 0; i < b.N; i++ {
		numSet := numSets[i%4]
		total := 0
		for _, num := range numSet {
			total += num
		}
	}
}

func TestRecursiveTreeTraersal(t *testing.T) {
	a := &BinaryTreeNode{
		left: &BinaryTreeNode{
			left: &BinaryTreeNode{
				value: "C",
			},
			value: "B",
			right: &BinaryTreeNode{
				value: "D",
			},
		},
		value: "A",
		right: &BinaryTreeNode{
			left: &BinaryTreeNode{
				value: "F",
			},
			value: "E",
			right: &BinaryTreeNode{
				value: "G",
			},
		},
	}

	strSum := RecursiveTreeTraversal(a)
	assert.Equal(t, "C->B->D->A->F->E->G->", strSum)
}

func BenchmarkRecursiveTreeTraersal(t *testing.B) {
	a := &BinaryTreeNode{
		value: "A",
		left: &BinaryTreeNode{
			value: "B",
			left: &BinaryTreeNode{
				value: "C",
			},
			right: &BinaryTreeNode{
				value: "D",
			},
		},
		right: &BinaryTreeNode{
			value: "E",
			left: &BinaryTreeNode{
				value: "F",
			},
			right: &BinaryTreeNode{
				value: "G",
			},
		},
	}

	for i := 0; i < t.N; i++ {
		RecursiveTreeTraversal(a)
	}
}
