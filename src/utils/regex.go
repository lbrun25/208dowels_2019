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