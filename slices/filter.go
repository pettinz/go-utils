package slices

func Filter[T any](objs []T, f func(int, T) bool) []T {
	r := []T{}

	for i, o := range objs {
		if f(i, o) {
			r = append(r, o)
		}
	}

	return r
}
