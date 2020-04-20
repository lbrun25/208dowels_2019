package utils

import (

)

// SliceFloatIndex - get index of a float slice at a specific key
func SliceFloatIndex(values []float64, key float64) int {
	for i, value := range values {
		if value == key {
			return i
		}
	}
	return -1
}