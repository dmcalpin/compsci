package generics

type GMap[T comparable, T2 any] map[T]T2

func (m GMap[T, T2]) ForEach(f func(T, T2)) GMap[T, T2] {
	for k, v := range m {
		f(k, v)
	}

	return m
}

func (m GMap[T, T2]) Map(f func(T, T2) T2) GMap[T, T2] {
	for k, v := range m {
		m[k] = f(k, v)
	}

	return m
}

func (m GMap[T, T2]) Set(k T, v T2) GMap[T, T2] {
	m[k] = v

	return m
}

func (m GMap[T, T2]) Get(k T) (T2, bool) {
	val, ok := m[k]
	return val, ok
}

func (m GMap[T, T2]) Delete(k T) GMap[T, T2] {
	delete(m, k)

	return m
}
