package concepts

// Hashing is the process of taking a value, and
// computing a smaller value which represents
// the original.
//
// In a Hash Table, the hash function computes
// an index at which data is stored and
// subsequently retrieved or deleted.

type HashTable struct {
	size   int
	values []interface{}
}

// NewHashTable creates a new HashTable
// of the given size n. n should be large
// enough to reduce the chance of collisions.
func NewHashTable(n int) *HashTable {
	return &HashTable{
		size:   n,
		values: make([]interface{}, n),
	}
}

// Insert will hash the key and store both
// the key and value at the computed index
// Efficiency: O(1)
func (h *HashTable) Insert(key string, value interface{}) {
	i := h.simpleHash(key)
	h.values[i] = value
}

// Get will hash the key and retrieve the
// value at the computed index
// Efficiency: O(1)
func (h *HashTable) Get(key string) interface{} {
	i := h.simpleHash(key)
	return h.values[i]
}

// Delete will hash the key and set
// the value at the computed index to
// nil
// Efficiency: O(1)
func (h *HashTable) Delete(key string) {
	i := h.simpleHash(key)
	h.values[i] = nil
}

// This is a basic hash which sums the int
// values of each rune and then uses mod to
// fit it into the table
func (h *HashTable) simpleHash(s string) int {
	runes := []rune(s)

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(runes[i])
	}

	index := sum % h.size
	return index
}
