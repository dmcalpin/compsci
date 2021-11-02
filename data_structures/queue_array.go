package data_structures

// Queue is a First-in, First-out (FI-FO)
// collection. This Queue uses a
// slice as its underlying structure

type Queue struct {
	items []interface{}
}

// Returns a new Queue prepopulated
// with any items passed in. The
// front of the queue is index 0
func NewQueue(items []interface{}) *Queue {
	return &Queue{
		items: items,
	}
}

// Adds an item to the back of the
// queue
func (q *Queue) Push(k interface{}) {
	q.items = append(q.items, k)
}

// Returns the first item in the queue
// and removes it from the list
func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Returns the first item in the queue
// but does not remove it
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}

// Checks if there are no items
// in the queue
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
