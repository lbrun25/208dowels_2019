package utils

import (
	"regexp"
)

// IsInteger - check if the value is a positive integer
func IsInteger(arg string) bool {
	var re = regexp.MustCompile("[-+]?\\d+")

	match := re.FindString(arg)
	if len(arg) != len(match) {
		return false
	}
	return true
}

// IsFloat - check if the value is a positive float
func IsFloat(arg string) bool {
	var re = regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")

	match := re.FindString(arg)
	if len(arg) != len(match) {
		return false
	}
	return true
}