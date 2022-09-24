package maps

func Zip[K comparable, T any](keys []K, values []T) map[K]T {
	res := make(map[K]T)

	for i, key := range keys {
		res[key] = values[i]
	}

	return res
}
