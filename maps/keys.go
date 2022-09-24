package maps

func Keys[K comparable, T any](d map[K]T) []K {
	res := []K{}

	for k := range d {
		res = append(res, k)
	}

	return res
}
