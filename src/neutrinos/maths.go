package neutrinos

import (
	"math"
)

// GetNumberRecordedValues - Get number of recorded values
func GetNumberRecordedValues() int {
	V.numberValues += 1
	return V.numberValues
}

// GetStandardDeviation - Get standard deviation
func GetStandardDeviation() float64 {
	lhs := ((math.Pow(V.standardDeviation, 2) + math.Pow(V.arithmeticMean, 2)) *
		(V.NumberValues() - 1) + math.Pow(NextValue, 2)) / V.NumberValues()
	rhs := math.Pow(((V.arithmeticMean * (V.NumberValues() - 1)) + NextValue) / V.NumberValues(), 2)
	res := math.Sqrt(lhs - rhs)
	V.standardDeviation = res
	return res
}

// GetArithmeticMean - Get arithmetic mean
func GetArithmeticMean() float64 {
	res := V.arithmeticMean - (V.arithmeticMean / V.NumberValues()) + (NextValue / V.NumberValues())
	V.arithmeticMean = res
	return res
}

// GetRootMeanSquare - Get root mean square
func GetRootMeanSquare() float64 {
	numerator := math.Pow(V.standardDeviation, 2) * V.NumberValues() - math.Pow(V.standardDeviation, 2) +
		math.Pow(V.arithmeticMean, 2) * V.NumberValues() - math.Pow(V.arithmeticMean, 2) +
		math.Pow(NextValue, 2)
	denominator := V.NumberValues()
	res := math.Sqrt(numerator / denominator)
	V.rootMeanSquare = res
	return res
}

// GetHarmonicMean - Get harmonic mean
func GetHarmonicMean() float64 {
	numerator := NextValue * V.NumberValues() * V.harmonicMean
	denominator := NextValue * (V.NumberValues() - 1) + V.harmonicMean
	res := numerator / denominator
	V.harmonicMean = res
	return res
}