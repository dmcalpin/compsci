package concepts

// Hashing is the process of taking a value, and
// computing a smaller value which represents
// the original.

// This is a basic hash which sums the int
// values of each rune and then uses mod to
// fit it into the table
func SimpleHash(s string, size int) int {
	runes := []rune(s)

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(runes[i])
	}

	index := sum % size
	return index
}
