package main

// HasInt32 checks if an int32 slice has a given int
func HasInt32(items []int, value int) bool {
	i, _ := FindInt32(items, value)
	return i > -1
}

// FindInt32Index returns the index of the given Int32 found in the given slice
// or -1 if none was found
func FindInt32Index(items []int, value int) int {
	i, _ := FindInt32(items, value)
	return i
}

// FindInt32 returns the index and the given Int32 found in the slice
// or -1, -1 if none was found
func FindInt32(items []int, value int) (i int, val int) {
	for i, val := range items {
		if val == value {
			return i, val
		}
	}
	return -1, -1
}
