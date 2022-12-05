package slices

func Unique[T comparable](objs []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range objs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
