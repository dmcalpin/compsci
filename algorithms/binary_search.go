package compsci

// Binary search works by cutting
// the search in half each iteration
// The collection must be sorted
// first.
//
// The algorithm is as follows:
// 0. left = 0, right = len(collection) - 1
// Loop ->
// 1. Start in the center of left/right
// 2. If search > iter, left = i+1
//    continue
// 3. If search < iter, right = i
//    continue
// 4. must have match, return i
//
// It has an efficiency of O(log n)

// Returns index of found item, or -1 if not found
func BinarySearch(collection []int, item int) int {
	left := 0
	right := len(collection) - 1

	for left <= right {
		middle := (left + right) / 2
		mid := collection[middle]

		if item < mid {
			right = middle - 1
		} else if item > mid {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
