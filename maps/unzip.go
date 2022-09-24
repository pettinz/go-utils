package maps

func Unzip[K comparable, T any](d map[K]T) ([]K, []T) {
	keys := []K{}
	values := []T{}

	for key, value := range d {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}
