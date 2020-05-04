package dowels

import "fmt"

const (
)

// O - Slice of the size of the observed class
var O []int

func displayResults() {
	for i, o := range O {
		fmt.Println(i, "=", o)
	}
}

// Dowels - main
func Dowels() {
	displayResults()
}