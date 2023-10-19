package compsci

import "math/rand"

func QuickSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	// define the lower and upper bounds
	left, right := 0, n-1

	// pick a random index to pivot on
	pivot := rand.Intn(n)

	// move the pivot to the end
	swap(nums, pivot, right)

	for i := range nums {
		if nums[i] < nums[right] {
			// swap the left and the current number if
			// the current is less than the right
			swap(nums, left, i)
			left++
		}
	}

	// swap the left and right
	swap(nums, left, right)

	// Sort the smaller numbers
	QuickSort(nums[:left])

	// Sort the larger numbers
	QuickSort(nums[left+1:])

	return nums
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
