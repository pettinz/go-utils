package slices

func Reduce[T any, R any](objs []T, f func(int, R, T) R, initial R) R {
	for i, o := range objs {
		initial = f(i, initial, o)
	}

	return initial
}
