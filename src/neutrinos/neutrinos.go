package neutrinos

import (
	"fmt"
	"os"
)

const (
)

// A - Struct which holds arguments
var A = Args{}

type Args struct {
	numberValues int
	arithmeticMean float64
	harmonicMean float64
	standardDeviation float64
}

func (a Args) NumberValues() float64 {
	return float64(a.numberValues)
}

// NextValue - next value given in the input
var NextValue = 1.0

func displayResults() {
	fmt.Printf("\tNumber of values:   %d\n", A.numberValues)
	fmt.Printf("\tStandard deviation: %.2f\n", A.standardDeviation)
	fmt.Printf("\tArithmetic mean:    %.2f\n", A.arithmeticMean)
	fmt.Printf("\tRoot mean square:   %.2f\n", 0.0)
	fmt.Printf("\tHarmonic means:     %.2f\n\n", A.harmonicMean)
}

func updateValues() {
	GetNumberRecordedValues()
	GetRootMeanSquare()
	GetStandardDeviation()
	GetArithmeticMean()
	GetRootMeanSquare()
	GetHarmonicMean()
}

// Neutrinos - main
// TODO: Delete this line when I will finish the project
func Neutrinos() {
	var input string

	for ;; {
		fmt.Printf("Enter next value: ")
		_, _ = fmt.Scanln(&input)
		fmt.Println("")

		if input == "END" {
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