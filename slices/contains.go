package slices

func Contains[T comparable](objs []T, v T) bool {
	for _, t := range objs {
		if t == v {
			return true
		}
	}

	return false
}
