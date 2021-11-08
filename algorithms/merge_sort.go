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
	merged := []int{}

	l := 0
	r := 0

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			merged = append(merged, right[r])
			r++
		} else {
			merged = append(merged, left[l])
			l++
		}
	}

	for l < len(left) {
		merged = append(merged, left[l])
		l++
	}
	for r < len(right) {
		merged = append(merged, right[r])
		r++
	}

	return merged
}
