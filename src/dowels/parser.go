package dowels

import (
	"fmt"
	"os"
	"utils"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBePositiveInteger = "must be a positive integer.\n"
	sumMustBe100 = "The sum of the arguments must be strictly equal to 100.\n"
	maxArg = 9
	minArg = 9
)

func printError(errorMessage string) {
	fmt.Printf("Error: %s\n", errorMessage)
}

func printErrorWithValue(valueName string, errorMessage string) {
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

func getIntegerPositiveValue(valueName string, arg string) (bool, int) {
	if !utils.IsInteger(arg) {
		printErrorWithValue(valueName, mustBePositiveInteger)
		return false, -1
	}
	integer := utils.ConvertStringToInt(arg)
	if integer < 0 {
		printErrorWithValue(valueName, mustBePositiveInteger)
		return false, -1
	}
	return true, integer
}

func checkSum() bool {
	sum := 0

	for _, value := range O {
		sum += value
	}
	if sum != 100 {
		printError(sumMustBe100)
		return false
	}
	return true
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < minArg {
		printError(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > maxArg {
		printError(tooManyArgs)
		return false
	}
	valueNames := [9]string{"O0", "O1", "O2", "O3", "O4", "O5", "O6", "O7", "O8"}
	for i, arg := range argsWithoutProg {
		valueName := valueNames[i]

		status, integer := getIntegerPositiveValue(valueName, arg)
		if !status {
			return false
		}
		O = append(O, integer)
	}
	if !checkSum() {
		return false
	}
	return true
}