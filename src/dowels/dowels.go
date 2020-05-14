package dowels

import "fmt"

const (
)

// O - Slice of the size of the observed class
var O []int

var c[][]float64

var d float64

var tx []float64

var p float64

var v int

func displayResults() {
	PrintTab()
	fmt.Printf("Distribution:           B(100, %.4f)\n", p)
	fmt.Printf("Chi-squared:            %.3f\n", d)
	fmt.Printf("Degrees of freedom:     %d\n", v)
	fmt.Printf("Fit validity:           %s\n", GetFitValidity())
}

// Dowels - main
func Dowels() {
	p = GetProbability()
	tx = CreateTx()
	c = CreateMatrixSquare()
	d = getD()
	v = GetFreedomDegrees()

	displayResults()
}