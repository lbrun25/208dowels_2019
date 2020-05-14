package dowels

import "fmt"

const (
)

// Ox - Slice of the size of the observed class
var Ox []int

// Dowels - Holder of StructDowels
var Dowels StructDowels

// StructDowels - struct which holds the necessary values
type StructDowels struct {
	c[][]float64
	d float64
	tx []float64
	p float64
	v int
}

func displayResults() {
	PrintTab()
	fmt.Printf("Distribution:           B(100, %.4f)\n", Dowels.p)
	fmt.Printf("Chi-squared:            %.3f\n", Dowels.d)
	fmt.Printf("Degrees of freedom:     %d\n", Dowels.v)
	fmt.Printf("Fit validity:           %s\n", GetFitValidity())
}

// Main - Dowels' main
func Main() {
	Dowels.c = CreateMatrixSquare()
	Dowels.p = GetProbability()
	Dowels.tx = CreateTx()
	Dowels.d = getD()
	Dowels.v = GetFreedomDegrees()

	displayResults()
}