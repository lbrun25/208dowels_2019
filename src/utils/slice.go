package utils

func GetMatrixCopyOf(src [][]float64) [][]float64 {
	dest := make([][]float64, len(src))
	for i := range src {
		dest[i] = make([]float64, len(src[i]))
	}
	return dest
}

// GetSum - get sum of slice
func GetSum(slice []float64) float64 {
	sum := 0.0

	for _, e := range slice {
		sum += e
	}
	return sum
}