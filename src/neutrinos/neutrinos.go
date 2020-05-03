package neutrinos

import (
	"fmt"
	"os"
)

const (
)

// Values - Struct which holds values
var V = Values{}

type Values struct {
	numberValues int
	arithmeticMean float64
	harmonicMean float64
	standardDeviation float64
	rootMeanSquare float64
}

func (v Values) NumberValues() float64 {
	return float64(v.numberValues)
}

// NextValue - next value given in the input
var NextValue = 1.0

func displayResults() {
	fmt.Printf("\tNumber of values:   %d\n", V.numberValues)
	fmt.Printf("\tStandard deviation: %.2f\n", V.standardDeviation)
	fmt.Printf("\tArithmetic mean:    %.2f\n", V.arithmeticMean)
	fmt.Printf("\tRoot mean square:   %.2f\n", V.rootMeanSquare)
	fmt.Printf("\tHarmonic mean:     %.2f\n\n", V.harmonicMean)
}

func updateValues() {
	GetNumberRecordedValues()
	GetRootMeanSquare()
	GetStandardDeviation()
	GetArithmeticMean()
	GetHarmonicMean()
}

// Neutrinos - main
func Neutrinos() {
	var input string

	for ;; {
		fmt.Printf("Enter next value: ")
		_, _ = fmt.Scanln(&input)
		if input == "END" || input == "" {
			break
		}
		status, value := CheckNextValueFormat(input)
		if !status {
			os.Exit(84)
		}
		NextValue = float64(value)
		updateValues()
		displayResults()
	}
}