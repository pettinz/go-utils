package slices

func FlatMap[T any, R any](objs []T, f func(int, T) []R) []R {
	r := []R{}

	for i, o := range objs {
		r = append(r, f(i, o)...)
	}

	return r
}
