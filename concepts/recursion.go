package concepts

// Recursion is the process of a function
// calling iteself, and after some
// condition is satisfied, it returns
// up the chain of calls.
// f() -> f() -> f() -> f() -|
// ret <- ret <- ret <- ret <|

// RecursiveSum will sum a set of provided
// numbers by adding the first number
// to the sum of the subsequent numbers,
// which is determined by calling itself
func RecursiveSum(nums ...int) int {
	// this condition is when
	// recursion stops and the chain
	// of retuns starts
	if len(nums) == 0 {
		return 0
	}

	// For each number this funciton
	// calls itself with the next
	// set of numbers
	return nums[0] + RecursiveSum(nums[1:]...)
}

// A BinaryTreeNode has a value, a left branch and
// a right branch, and can look similar to
// this:
//     A
//    / \
//   B   C
//  / \
// D   E
type BinaryTreeNode struct {
	value string
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// RecursiveTreeTraversal does a in-order,
// depth-first traversal of the provided tree
// node.
// This means it reads the left, then the value,
// then the right branch.
// Efficiency: O(n), n being the number of
// nodes
func RecursiveTreeTraversal(tree *BinaryTreeNode) string {
	if tree == nil {
		return ""
	}

	strSum := RecursiveTreeTraversal(tree.left)
	strSum += tree.value + "->"
	strSum += RecursiveTreeTraversal(tree.right)

	return strSum
}
