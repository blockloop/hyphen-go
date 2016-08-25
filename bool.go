package hyphen

// HasBool checks if a bool slice has a given bool
func HasBool(items []bool, value bool) bool {
	i, _ := FindBool(items, value)
	return i > -1
}

// FindBoolIndex returns the index of the given bool found in the given slice
// or -1 if none was found
func FindBoolIndex(items []bool, value bool) int {
	i, _ := FindBool(items, value)
	return i
}

// FindBool returns the index and the given bool found in the slice
// or -1, false if none was found
func FindBool(items []bool, value bool) (i int, val bool) {
	for i, val := range items {
		if val == value {
			return i, val
		}
	}
	return -1, false
}
