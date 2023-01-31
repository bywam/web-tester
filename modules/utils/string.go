package utils

func Contains(strFind string, items []string) bool {
	for _, s := range items {
		if s == strFind {
			return true
		}
	}
	return false
}
