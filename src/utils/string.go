package utils

import (
	"strconv"
)

// ConvertStringToFloat - convert a string to float by checking error
func ConvertStringToFloat(str string) float64 {
	s, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(err)
	}
	return s
}

// ConvertStringToInt - convert a string to int by checking error
func ConvertStringToInt(str string) int {
	s, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
	return s
}