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

func (gs *GSet[T]) Enumerate() []T {
	elems := []T{}

	for key := range gs.val {
		elems = append(elems, key)
	}

	return elems
}

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

func (gs *GSet[T]) Add(elem T) {
	gs.val[elem] = true
}

func (gs *GSet[T]) Union(otherSet *GSet[T]) *GSet[T] {
	union := NewGSet(gs.Enumerate()...)
	for elem := range otherSet.val {
		union.Add(elem)
	}
	return union
}
