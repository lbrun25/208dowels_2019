package dowels

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
	"utils"
)

var distributionTable = [][]float64 {
	{0.00, 0.02, 0.06, 0.15, 0.27, 0.45, 0.71, 1.07, 1.64, 2.71, 3.84, 5.41, 6.63},
	{0.02, 0.21, 0.45, 0.71, 1.02, 1.39, 1.83, 2.41, 3.22, 4.61, 5.99, 7.82, 9.21},
	{0.11, 0.58, 1.01, 1.42, 1.87, 2.37, 2.95, 3.66, 4.64, 6.25, 7.81, 9.84, 11.35},
	{0.30, 1.06, 1.65, 2.19, 2.75, 3.36, 4.04, 4.88, 5.99, 7.78, 9.49, 11.67, 13.28},
	{0.55, 1.61, 2.34, 3.00, 3.66, 4.35, 5.13, 6.06, 7.29, 9.24, 11.07, 13.33, 15.01},
	{0.70, 2.20, 3.07, 3.83, 4.57, 5.35, 6.21, 7.23, 8.56, 10.64, 12.59, 15.03, 16.81},
	{1.24, 2.83, 3.82, 4.67, 5.49, 6.35, 7.28, 8.38, 9.80, 12.02, 14.07, 16.62, 18.48},
	{1.65, 3.49, 4.59, 5.53, 6.42, 7.34, 8.35, 9.52, 11.03, 13.36, 15.51, 18.17, 20.09},
	{2.09, 4.17, 5.38, 6.39, 7.36, 8.34, 9.41, 10.66, 12.24, 14.68, 16.92, 19.63, 21.67},
	{2.56, 4.87, 6.18, 7.27, 8.30, 9.34, 10.47, 11.78, 13.44, 15.99, 18.31, 21.16, 23.21}}

var fits = []string {
	"P > 99%",
	"90% < P < 99%",
	"80% < P < 90%",
	"70% < P < 80%",
	"60% < P < 70%",
	"50% < P < 60%",
	"40% < P < 50%",
	"30% < P < 40%",
	"20% < P < 30%",
	"10% < P < 20%",
	"5% < P < 10%",
	"2% < P < 5%",
	"1% < P < 2%",
	"P < 1%",
}

var sSumTx []float64
var sSumOX []int

// GetProbability - get probability of the distribution
func GetProbability() float64 {
	sum := 0.0

	for i, o := range Ox {
		sum += float64(i * o)
	}
	res := sum / math.Pow(10, 4)
	return res
}

// GetChiSquared - get Chi-squared value
func GetChiSquared() float64 {
	res := 0.0
	matrix := Dowels.matrix
	sTx := Dowels.sTx
	ssTx := utils.GetMatrixCopyOf(matrix)
	ssLhs := utils.GetMatrixCopyOf(matrix)

	//Fill 2D slices
	k := 0
	for i, r := range ssTx {
		for j := range r {
			ssTx[i][j] = sTx[k]
			ssLhs[i][j] = float64(Ox[k]) - sTx[k]
			k++
		}
	}

	// Get sum slices
	var sumLhs []float64
	for i, x := range matrix {
		sum := 0.0
		sumLhsRes := 0.0
		for j, _ := range x {
			sum += ssTx[i][j]
			sumLhsRes += ssLhs[i][j]
		}
		sSumTx = append(sSumTx, sum)
		sumLhs = append(sumLhs, sumLhsRes)
	}

	// Get the result
	for i := range sSumTx {
		lhs := 0.0
		if sumLhs != nil {
			lhs = math.Pow(sumLhs[i], 2)
		}
		res += lhs / sSumTx[i]
	}
	return res
}

func getSumOx(s []float64) int {
	sum := 0

	for _, value := range s {
		i := int(value)
		sum += Ox[i]
	}
	return sum
}

// GetMatrixSquare - get matrix square
func GetMatrixSquare() [][]float64 {
	row := 9
	col := 1
	matrix := make([][]float64, row)
	rows := make([]float64, row * col)
	for i := 0; i < row; i++ {
		matrix[i] = rows[i * col : (i + 1) * col]
		matrix[i][0] = float64(i)
	}

	for i := 0; i < len(matrix); {
		if getSumOx(matrix[i]) >= 10 {
			i++
			continue
		} else if (i + 1 == len(matrix)) || (i > 0 && getSumOx(matrix[i - 1]) < getSumOx(matrix[i + 1])) {
			matrix[i - 1] = append(matrix[i - 1], matrix[i]...)
			// Delete row
			matrix = append(matrix[:i], matrix[i + 1:]...)
			i -= 1
		} else {
			matrix[i] = append(matrix[i], matrix[i + 1]...)
			// Delete row
			matrix = append(matrix[:i + 1], matrix[i + 2:]...)
		}
	}
	return matrix
}

// GetTxValues - get slice of Tx values
func GetTxValues() []float64 {
	var s []float64

	for i := range Ox[0:8] {
		s = append(s, 100 * utils.GetBinomial(100, int64(i), Dowels.probability))
	}
	s = append(s, 100 - utils.GetSum(s))
	return s
}

func printRow(slice interface{}, format string, start string, delimiter string, end string) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	fmt.Printf("%s", start)
	for i := 0; i < s.Len(); i++ {
		s := fmt.Sprintf(format, s.Index(i).Interface())
		s += delimiter
		fmt.Printf("%s", s)
	}
	fmt.Println(end)
}

func computeSumOxSlice() {
	for _, x := range Dowels.matrix {
		sSumOX = append(sSumOX, getSumOx(x))
	}
}

// PrintTab - print tab of x, Ox, Tx
func PrintTab() {
	matrix := Dowels.matrix
	var xSlice []string

	computeSumOxSlice()
	for i, x := range matrix {
		var values []string
		var result string

		for _, y := range x {
			s := fmt.Sprintf("%d", int(y))
			values = append(values, s)
		}
		if values != nil && i == len(matrix) - 1 {
			result = values[0] + "+"
		} else if values != nil && len(values) > 2 {
			result = values[0] + "-" + values[len(values) - 1]
		} else {
			result = strings.Join(values, "-")
		}
		xSlice = append(xSlice, result)
	}
	printRow(xSlice, "%s", "   x\t| ", "\t| ", "Total")
	printRow(sSumOX, "%d", "  Ox\t| ", "\t| ", "100")
	printRow(sSumTx, "%.1f","  Tx\t| ", "\t| ", "100")
}

// GetFreedomDegrees - get degrees of freedom
func GetFreedomDegrees() int {
	degrees := len(Dowels.matrix) - 2

	if degrees < 1 {
		printError("Something went wrong, degrees of freedom must be greater than one")
		os.Exit(84)
	}
	return degrees
}

// GetFitValidity - get fits validity
func GetFitValidity() string {
	i := len(fits) - 1

	for j, value := range distributionTable[Dowels.freedomDegrees- 1] {
		if Dowels.chiSquared < value {
			i = j
			break
		}
	}
	return fits[i]
}