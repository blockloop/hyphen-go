package main

// HasFloat64 checks if a Float64 slice has a given Float64
func HasFloat64(items []float64, value float64) bool {
	i, _ := FindFloat64(items, value)
	return i > -1
}

// FindFloat64Index returns the index of the given Float64 found in the given slice
// or -1 if none was found
func FindFloat64Index(items []float64, value float64) int {
	i, _ := FindFloat64(items, value)
	return i
}

// FindFloat64 returns the index and the given Float64 found in the slice
// or -1, -1 if none was found
func FindFloat64(items []float64, value float64) (i int, val float64) {
	for i, val := range items {
		if val == value {
			return i, val
		}
	}
	return -1, -1
}
