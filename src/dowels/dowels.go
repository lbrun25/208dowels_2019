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
	matrix         [][]float64
	chiSquared     float64
	sTx            []float64
	probability    float64
	freedomDegrees int
}

func displayResults() {
	PrintTab()
	fmt.Printf("Distribution:           B(100, %.4f)\n", Dowels.probability)
	fmt.Printf("Chi-squared:            %.3f\n", Dowels.chiSquared)
	fmt.Printf("Degrees of freedom:     %d\n", Dowels.freedomDegrees)
	fmt.Printf("Fit validity:           %s\n", GetFitValidity())
}

// Main - Dowels' main
func Main() {
	Dowels.matrix = GetMatrixSquare()
	Dowels.probability = GetProbability()
	Dowels.sTx = GetTxValues()
	Dowels.chiSquared = GetChiSquared()
	Dowels.freedomDegrees = GetFreedomDegrees()

	displayResults()
}