package slices

func Flatten[T any](objs [][]T) []T {
	r := []T{}

	for _, o := range objs {
		r = append(r, o...)
	}

	return r
}
