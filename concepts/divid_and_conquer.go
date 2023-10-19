package concepts

// Divide and Conquer algorithms split up
// the work that needs to be done
// and then processes the results after
//
// See ../algorithms/merge_sort.go

func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + Sum(nums[1:])
}
