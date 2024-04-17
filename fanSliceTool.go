package fanTool

func ValueExists[V comparable](slice []V, value V) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
