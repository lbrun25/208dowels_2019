package neutrinos

import (
	"fmt"
	"os"
	"utils"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBePositiveInteger = "must be a positive integer.\n"
	mustBePositiveFloat = "must be a positive float.\n"
	mustBeGreatherThanZero = "must be greater than zero.\n"
	wrongNextValueFormat = "Wrong next value's format which must be a positive integer."
	maxArg = 4
	minArg = 4
)

func printError(valueName string, errorMessage string) {
	fmt.Printf("Error: '%s' %s\n", valueName, errorMessage)
}

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if arg == "-h" {
			return true
		}
	}
	return false
}

// CheckNextValueFormat - check next value that is entered in the input
func CheckNextValueFormat(input string) (bool, int) {
	if !utils.IsInteger(input) {
		fmt.Println(wrongNextValueFormat)
		return false, -1
	}
	resInt := utils.ConvertStringToInt(input)
	if resInt <= 0 {
		fmt.Printf("The next value %s", mustBeGreatherThanZero)
		return false, -1
	}
	return true, resInt
}

func getIntegerPositiveValueGreaterThanZero(valueName string, arg string) (bool, int) {
	if !utils.IsInteger(arg) {
		printError(valueName, mustBePositiveInteger)
		return false, -1
	}
	integer := utils.ConvertStringToInt(arg)
	if integer <= 0 {
		printError(valueName, mustBeGreatherThanZero)
		return false, -1
	}
	return true, integer
}

func getFloatPositiveValueGreaterThanZero(valueName string, arg string) (bool, float64) {
	if !utils.IsFloat(arg) {
		printError(valueName, mustBePositiveFloat)
		return false, -1
	}
	float := utils.ConvertStringToFloat(arg)
	if float <= 0 {
		printError(valueName, mustBeGreatherThanZero)
		return false, -1
	}
	return true, float
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < minArg {
		fmt.Println(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > maxArg {
		fmt.Println(tooManyArgs)
		return false
	}
	valueNames := [4]string{"n", "a", "h", "sd"}
	for i, arg := range argsWithoutProg {
		valueName := valueNames[i]

		// Check and assign n
		if valueName == "n" {
			status, integer := getIntegerPositiveValueGreaterThanZero(valueName, arg)
			if !status {
				return false
			}
			V.numberValues = integer
		}

		// Check and assign a
		if valueName == "a" {
			status, float := getFloatPositiveValueGreaterThanZero(valueName, arg)
			if !status {
				return false
			}
			V.arithmeticMean = float
		}

		// Check and assign h
		if valueName == "h" {
			status, float := getFloatPositiveValueGreaterThanZero(valueName, arg)
			if !status {
				return false
			}
			V.harmonicMean = float
		}

		// Check and assign sd
		if valueName == "sd" {
			status, float := getFloatPositiveValueGreaterThanZero(valueName, arg)
			if !status {
				return false
			}
			V.standardDeviation = float
		}
	}
	return true
}