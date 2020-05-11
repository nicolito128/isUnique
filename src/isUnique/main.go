package isunique

// Int ... check if a []int is unique
func Int(arr []int) bool {
	exists := make(map[int]bool)

	for i := range arr {
		el := arr[i]
		if exists[el] {
			return false
		} else {
			exists[el] = true
		}
	}

	return true
}

// Str ... check if a []string is unique
func Str(arr []string) bool {
	exists := make(map[string]bool)

	for i := range arr {
		el := arr[i]
		if exists[el] {
			return false
		} else {
			exists[el] = true
		}
	}

	return true
}
