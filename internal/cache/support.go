package cache

func contains[T comparable](items []T, search T) bool {
	for _, item := range items {
		if search == item {
			return true
		}
	}
	return false
}
