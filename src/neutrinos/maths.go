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
	res := A.arithmeticMean - (A.arithmeticMean / A.NumberValues()) + (NextValue / A.NumberValues())
	A.arithmeticMean = res
	return res
}

// GetRootMeanSquare - Get root mean square
func GetRootMeanSquare() float64 {
	numerator := math.Pow(A.standardDeviation, 2) * A.NumberValues() - math.Pow(A.standardDeviation, 2) +
		math.Pow(A.arithmeticMean, 2) * A.NumberValues() - math.Pow(A.arithmeticMean, 2) +
		math.Pow(NextValue, 2)
	denominator := A.NumberValues()
	res := math.Sqrt(numerator / denominator)
	return res
}

// GetHarmonicMean - Get harmonic mean
func GetHarmonicMean() float64 {
	numerator := NextValue * A.NumberValues() * A.harmonicMean
	denominator := NextValue * (A.NumberValues() - 1) + A.harmonicMean
	res := numerator / denominator
	A.harmonicMean = res
	return res
}