package internal

func containValue(available []string, value string) bool {
	for _, item := range available {
		if item == value {
			return true
		}
	}
	return false
}
