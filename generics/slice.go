package generics

type GSlice[T any] []T

func (m GSlice[T]) ForEach(f func(T)) GSlice[T] {
	for _, v := range m {
		f(v)
	}

	return m
}

func (m GSlice[T]) Map(f func(T) T) GSlice[T] {
	for k, v := range m {
		m[k] = f(v)
	}

	return m
}
