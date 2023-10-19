package compsci

import "math"

// SelectionSort sorts by searching the
// slice for the smallest number, removes
// that number from the slice and appends
// it to a new slice.
//
// This runs in O(n^2) time.
func SelectionSort(nums []int) []int {
	var smallest = func(nums []int) int {
		smallest := math.MaxInt
		smallestIndex := 0

		for i, num := range nums {
			if num < smallest {
				smallest = num
				smallestIndex = i
			}
		}

		return smallestIndex
	}

	sorted := make([]int, len(nums))

	for i := range nums {
		smallestIndex := smallest(nums)
		sorted[i] = nums[smallestIndex]

		// Delete from the slice
		nums[smallestIndex] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}

	return sorted
}
