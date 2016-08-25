package main

// HasFloat32 checks if a Float32 slice has a given Float32
func HasFloat32(items []float32, value float32) bool {
	i, _ := FindFloat32(items, value)
	return i > -1
}

// FindFloat32Index returns the index of the given Float32 found in the given slice
// or -1 if none was found
func FindFloat32Index(items []float32, value float32) int {
	i, _ := FindFloat32(items, value)
	return i
}

// FindFloat32 returns the index and the given Float32 found in the slice
// or -1, -1 if none was found
func FindFloat32(items []float32, value float32) (i int, val float32) {
	for i, val := range items {
		if val == value {
			return i, val
		}
	}
	return -1, -1
}
