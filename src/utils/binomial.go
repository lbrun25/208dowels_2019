package utils

import (
	"fmt"
	"math/big"
	"os"
)

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

// GetBinomial - get binomial
func GetBinomial(n int64, k int64, p float64) float64 {
	res := big.NewFloat(0.0).Mul(
		big.NewFloat(0.0).SetInt(getBinomialCoefficient(big.NewInt(0.0).SetInt64(n), big.NewInt(0.0).SetInt64(k))),
		BigPow(big.NewFloat(0.0).SetFloat64(p), k))
	res.Mul(res, BigPow(big.NewFloat(0.0).Sub(big.NewFloat(1.0), big.NewFloat(0.0).SetFloat64(p)), n - k))

	s := fmt.Sprintf("%f", res)
	resConverted := ConvertStringToFloat(s)
	return resConverted
}