package slices

func GroupBy[T any, R comparable](objs []T, f func(int, T) R) map[R][]T {
	r := make(map[R][]T)

	for i, o := range objs {
		k := f(i, o)
		r[k] = append(r[k], o)
	}

	return r
}
