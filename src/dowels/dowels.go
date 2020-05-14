package dowels

import "fmt"

const (
)

// O - Slice of the size of the observed class
var O []int

func displayResults() {
	fmt.Printf("Distribution:           B(100, %.4f)\n", GetProbability())
	fmt.Printf("Chi-squared:            %.3f\n", GetChiSquared())
	//fmt.Printf("Degrees of freedom:     %d\n", GetFreedomDegrees())
	//fmt.Printf("Fit validity:           %.2f\n", 0.0)
}

// Dowels - main
func Dowels() {
	displayResults()
}