package maps

func Values[K comparable, T any](d map[K]T) []T {
	res := []T{}

	for _, v := range d {
		res = append(res, v)
	}

	return res
}
