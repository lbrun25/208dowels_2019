package neutrinos

import (
	"fmt"
	"utils"
)

const (
)

func printNumberRecordedValues() {
	fmt.Printf("\tNumber of values:   %d\n", utils.GetNumberRecordedValues())
}

func printStandardDeviation() {
	fmt.Printf("\tStandard deviation: %.2f\n", utils.GetStandardDeviation())
}

func printArithmeticMean() {
	fmt.Printf("\tArithmetic mean:    %.2f\n", utils.GetArithmeticMean())
}

func printRootMeanSquare() {
	fmt.Printf("\tRoot mean square:   %.2f\n", utils.GetRootMeanSquare())
}

func printHarmonicMean() {
	fmt.Printf("\tHarmonic means:     %.2f\n\n", utils.GetHarmonicMean())
}

// Neutrinos - main
func Neutrinos() {
	var input string

	for ;; {
		fmt.Printf("Enter next value: ")
		fmt.Scanln(&input)
		fmt.Println("")

		if input == "END" {
			break
		}
		printNumberRecordedValues()
		printStandardDeviation()
		printArithmeticMean()
		printRootMeanSquare()
		printHarmonicMean()
	}
}