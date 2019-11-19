package basic

import "math/big"

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}

func GoGCD(a, b int) int64 {
	gcd := big.Int{}

	res := gcd.GCD(big.NewInt(0), big.NewInt(0), big.NewInt(int64(a)), big.NewInt(int64(b)))

	return res.Int64()
}
