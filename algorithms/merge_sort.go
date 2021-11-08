package compsci

func MergeSort(items []int) []int {
	n := len(items)

	// zero or one is already sorted
	if n <= 1 {
		return items
	}

	middle := n / 2
	left := items[0:middle]
	right := items[middle:n]

	sl := MergeSort(left)
	sr := MergeSort(right)

	return merge(sl, sr)
}

func merge(left, right []int) []int {
	merged := make([]int, len(left)+len(right))

	l := 0
	r := 0

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			merged[l+r] = right[r]
			r++
		} else {
			merged[l+r] = left[l]
			l++
		}
	}

	for l < len(left) {
		merged[l+r] = left[l]
		l++
	}
	for r < len(right) {
		merged[l+r] = right[r]
		r++
	}

	return merged
}
