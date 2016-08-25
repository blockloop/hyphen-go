package main

// HasString checks if a string slice has a given string
func HasString(items []string, value string) bool {
	i, _ := FindString(items, value)
	return i > -1
}

// FindStringIndex returns the index of the given String found in the given slice
// or -1 if none was found
func FindStringIndex(items []string, value string) int {
	i, _ := FindString(items, value)
	return i
}

// FindString returns the index and the given String found in the slice
// or -1, "" if none was found
func FindString(items []string, value string) (i int, val string) {
	for i, val := range items {
		if val == value {
			return i, val
		}
	}
	return -1, ""
}
