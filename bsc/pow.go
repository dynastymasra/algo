package main

import "fmt"

/**
Ref: https://www.geeksforgeeks.org/write-a-c-program-to-calculate-powxn/
*/
func main() {
	fmt.Println("2 power 5 is ", power(2, 5))
	fmt.Println("2 power -3 is ", power(2, -3))
}

func power(x float32, y int) float32 {
	if y == 0 {
		return 1
	}

	temp := power(x, y/2)

	if y%2 == 0 {
		return temp * temp
	}

	if y > 0 {
		return x * temp * temp
	}

	// if power < 0
	return (temp * temp) / x
}
