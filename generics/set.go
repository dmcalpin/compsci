package generics

type GSet[T comparable] struct {
	val map[T]bool
}

func NewGSet[T comparable](elements ...T) *GSet[T] {
	gset := &GSet[T]{
		val: make(map[T]bool),
	}

	for _, elem := range elements {
		gset.val[elem] = true
	}

	return gset
}

// CORE OPERATIONS

func (gs *GSet[T]) Union(otherSet *GSet[T]) *GSet[T] {
	union := NewGSet(gs.Enumerate()...)
	for elem := range otherSet.val {
		union.Add(elem)
	}
	return union
}

func (gs *GSet[T]) Intersection(otherSet *GSet[T]) *GSet[T] {
	intersection := NewGSet[T]()

	var baseSet *GSet[T]
	var compSet *GSet[T]
	if gs.Size() > otherSet.Size() {
		baseSet = otherSet
		compSet = gs
	} else {
		baseSet = gs
		compSet = otherSet
	}

	for elem := range baseSet.val {
		if compSet.Exists(elem) {
			intersection.Add(elem)
		}
	}

	return intersection
}

func (gs *GSet[T]) Difference(otherSet *GSet[T]) *GSet[T] {
	difference := NewGSet[T]()

	for elem := range gs.val {
		if !otherSet.Exists(elem) {
			difference.Add(elem)
		}
	}

	return difference
}

func (gs *GSet[T]) IsSubsetOf(otherSet *GSet[T]) bool {
	if gs.Size() > otherSet.Size() {
		return false
	}

	for elem := range gs.val {
		if !otherSet.Exists(elem) {
			return false
		}
	}

	return true
}

// STATIC SET OPERATIONS

func (gs *GSet[T]) Exists(elem T) bool {
	_, ok := gs.val[elem]
	return ok
}

func (gs *GSet[T]) IsEmpty() bool {
	return gs.Size() == 0
}

func (gs *GSet[T]) Size() int {
	return len(gs.val)
}

// Iterate Goes here

func (gs *GSet[T]) Enumerate() []T {
	elems := []T{}

	for key := range gs.val {
		elems = append(elems, key)
	}

	return elems
}

// DYNAMIC SET OPERATIONS

func (gs *GSet[T]) Add(elem T) {
	gs.val[elem] = true
}

func (gs *GSet[T]) Remove(elem T) {
	delete(gs.val, elem)
}

// OTHER OPERATIONS

func (gs *GSet[T]) Pop() *T {
	for elem := range gs.val {
		retVal := elem
		delete(gs.val, elem)
		return &retVal
	}
	return nil
}

// Pick Goes here

func (gs *GSet[T]) Map(f func(elem T) T) *GSet[T] {
	for elem := range gs.val {
		delete(gs.val, elem)
		gs.Add(f(elem))
	}
	return gs
}

func (gs *GSet[T]) Filter(f func(elem T) bool) *GSet[T] {
	for elem := range gs.val {
		if f(elem) {
			delete(gs.val, elem)
		}
	}
	return gs
}

// Fold Goes Here

func (gs *GSet[T]) Clear() *GSet[T] {
	for elem := range gs.val {
		delete(gs.val, elem)
	}
	return gs
}

func (gs *GSet[T]) Equal(otherSet *GSet[T]) bool {
	if gs.Size() != otherSet.Size() {
		return false
	}

	for elem := range gs.val {
		if !otherSet.Exists(elem) {
			return false
		}
	}
	return true
}

// Hash goes here
