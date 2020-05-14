package dowels

import (
	"fmt"
	"math"
	"math/big"
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

var sumTxTmps []float64

// GetProbability - compute p
func GetProbability() float64 {
	sum := 0.0

	for i, o := range O {
		sum += float64(i * o)
	}
	res := sum / math.Pow(10, 4)
	return res
}

// Factorial - Get factorial big int
func Factorial(x *big.Int) *big.Int {
	result := big.NewInt(1)
	i := big.NewInt(2)

	if !x.IsInt64() {
		fmt.Println("The number is way too big to calculate a factorial")
		os.Exit(84)
	}
	for i.Cmp(x) != 1 {
		result.Mul(result, i)
		i = i.Add(i, big.NewInt(1))
	}
	return result
}

func getBinomialCoefficient(n *big.Int, k *big.Int) *big.Int {
	if k.Cmp(n) == 1 {
		fmt.Println("Error: k > n")
		os.Exit(84)
	}

	numerator := Factorial(n)
	subNK := big.NewInt(1).Sub(n, k)
	denominator := big.NewInt(1).Mul(Factorial(k), Factorial(subNK))
	res := big.NewInt(1).Div(numerator, denominator)
	return res
}

// BigPow - big Float
func BigPow(a *big.Float, e int64) *big.Float {
	if e == 0 {
		return big.NewFloat(1.0)
	}
	result := big.NewFloat(0.0).Copy(a)
	for i := int64(0); i < e - 1; i++ {
		result = result.Mul(result, a)
	}
	return result
}

func getBinomial(n int64, k int64, p float64) float64 {
	res := big.NewFloat(0.0).Mul(
		big.NewFloat(0.0).SetInt(getBinomialCoefficient(big.NewInt(0.0).SetInt64(n), big.NewInt(0.0).SetInt64(k))),
		BigPow(big.NewFloat(0.0).SetFloat64(p), k))
	res.Mul(res, BigPow(big.NewFloat(0.0).Sub(big.NewFloat(1.0), big.NewFloat(0.0).SetFloat64(p)), n - k))

	s := fmt.Sprintf("%f", res)
	resConverted := utils.ConvertStringToFloat(s)
	return resConverted
}

func getSumOx(c []float64) int {
	sum := 0

	for _, value := range c {
		i := int(value)
		sum += O[i]
	}
	return sum
}

func RemoveIndex(s []float64, index int) []float64 {
	return append(s[:index], s[index+1:]...)
}

func getD() float64 {
	sumRes := 0.0

	// Copy c to txTmp
	txTmp := make([][]float64, len(c))
	lhs := make([][]float64, len(c))
	for i := range c {
		txTmp[i] = make([]float64, len(c[i]))
		lhs[i] = make([]float64, len(c[i]))
	}

	// Fill
	k := 0
	for i, r := range txTmp {
		for j := range r {
			txTmp[i][j] = tx[k]
			lhs[i][j] = float64(O[k]) - tx[k]
			k++
		}
	}

	// Sum
	var sumLhs []float64
	for i, x := range c {
		sum := 0.0
		sumLhsRes := 0.0
		for j, _ := range x {
			sum += txTmp[i][j]
			sumLhsRes += lhs[i][j]
		}
		sumTxTmps = append(sumTxTmps, sum)
		sumLhs = append(sumLhs, sumLhsRes)
	}

	// SumRes
	for i := range sumTxTmps {
		lhs := 0.0
		if sumLhs != nil {
			lhs = math.Pow(sumLhs[i], 2)
		}
		sumRes +=  lhs / sumTxTmps[i]
	}
	return sumRes
}

func CreateMatrixSquare() [][]float64 {
	row := 9
	col := 1
	c := make([][]float64, row)
	rows := make([]float64, row * col)
	for i := 0; i < row; i++ {
		c[i] = rows[i * col : (i + 1) * col]
		c[i][0] = float64(i)
	}

	for i := 0; i < len(c); {
		if getSumOx(c[i]) >= 10 {
			i++
			continue
		} else if (i + 1 == len(c)) || (i > 0 && getSumOx(c[i - 1]) < getSumOx(c[i + 1])) {
			c[i - 1] = append(c[i - 1], c[i]...)
			// Delete row
			c = append(c[:i], c[i + 1:]...)
			i -= 1
		} else {
			c[i] = append(c[i], c[i + 1]...)
			// Delete row
			c = append(c[:i + 1], c[i + 2:]...)
		}
	}
	return c
}

func getSum(slice []float64) float64 {
	sum := 0.0

	for _, e := range slice {
		sum += e
	}
	return sum
}

func CreateTx() []float64 {
	var tx []float64

	for i := range O[0:8] {
		tx = append(tx, 100 * getBinomial(100, int64(i), p))
	}
	tx = append(tx, 100 - getSum(tx))
	return tx
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

func PrintTab() {
	// First row
	var xSlice []string

	for i, x := range c {
		var values []string
		var result string

		for _, y := range x {
			s := fmt.Sprintf("%d", int(y))
			values = append(values, s)
		}
		if values != nil && i == len(c) - 1 {
			result = values[0] + "+"
		} else if values != nil && len(values) > 2 {
			result = values[0] + "-" + values[len(values) - 1]
		} else {
			result = strings.Join(values, "-")
		}
		xSlice = append(xSlice, result)
	}
	printRow(xSlice, "%s", "   x\t| ", "\t| ", "Total")

	// Second row
	var sumOx []int

	for _, x := range c {
		sumOx = append(sumOx, getSumOx(x))
	}
	printRow(sumOx, "%d", "  Ox\t| ", "\t| ", "100")

	// Third row
	printRow(sumTxTmps, "%.1f","  Tx\t| ", "\t| ", "100")
}

// GetFreedomDegrees - get degrees of freedom
func GetFreedomDegrees() int {
	degrees := len(c) - 2
	if degrees < 1 {
		printError("Something went wrong, degrees of freedom must be greater than one")
		os.Exit(84)
	}
	return degrees
}

// GetFitValidity - get fits validity
func GetFitValidity() string {
	i := len(fits) - 1
	for j, l := range distributionTable[v - 1] {
		if d < l {
			i = j
			break
		}
	}
	return fits[i]
}