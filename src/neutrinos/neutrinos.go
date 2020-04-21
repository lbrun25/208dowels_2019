package neutrinos

import (
	"fmt"
	"os"
)

const (
)

func printNumberRecordedValues() {
	fmt.Printf("\tNumber of values:   %d\n", GetNumberRecordedValues())
}

func printStandardDeviation() {
	fmt.Printf("\tStandard deviation: %.2f\n", GetStandardDeviation())
}

func printArithmeticMean() {
	fmt.Printf("\tArithmetic mean:    %.2f\n", GetArithmeticMean())
}

func printRootMeanSquare() {
	fmt.Printf("\tRoot mean square:   %.2f\n", GetRootMeanSquare())
}

func printHarmonicMean() {
	fmt.Printf("\tHarmonic means:     %.2f\n\n", GetHarmonicMean())
}

// Neutrinos - main
func Neutrinos() {
	var input string

	for ;; {
		fmt.Printf("Enter next value: ")
		fmt.Scanln(&input)
		if !CheckNextValueFormat(input) {
			os.Exit(84)
		}
		// TODO: Delete this line when I will finish the project
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