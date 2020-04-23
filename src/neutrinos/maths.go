package neutrinos

import (
	"math"
)

// GetNumberRecordedValues - Get number of recorded values
func GetNumberRecordedValues() int {
	A.numberValues += 1
	return A.numberValues
}

// GetStandardDeviation - Get standard deviation
func GetStandardDeviation() float64 {
	lhs := ((math.Pow(A.standardDeviation, 2) + math.Pow(A.arithmeticMean, 2)) *
		(A.NumberValues() - 1) + math.Pow(NextValue, 2)) / A.NumberValues()
	rhs := math.Pow(((A.arithmeticMean * (A.NumberValues() - 1)) + NextValue) / A.NumberValues(), 2)
	res := math.Sqrt(lhs - rhs)
	A.standardDeviation = res
	return res
}

// GetArithmeticMean - Get arithmetic mean
func GetArithmeticMean() float64 {
	sum := (A.arithmeticMean * (A.NumberValues() - 1)) + NextValue
	res := A.NumberValues() / sum
	//A.arithmeticMean = res
	return res
}

// GetRootMeanSquare - Get root mean square
func GetRootMeanSquare() float64 {
	return 0.0
}

// GetHarmonicMean - Get harmonic mean
func GetHarmonicMean() float64 {
	return 0.0
}