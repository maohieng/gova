package slices

func IndexOf[T comparable](slice []T, item T) int {
	for i, t := range slice {
		if t == item {
			return i
		}
	}
	return -1
}

func Contains[T comparable](slice []T, item T) bool {
	return IndexOf[T](slice, item) > -1
}
