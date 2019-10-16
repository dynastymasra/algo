package main

import "fmt"

/**
Ref: https://www.geeksforgeeks.org/finding-sum-of-digits-of-a-number-until-sum-becomes-single-digit/
*/
func main() {
	fmt.Println("Digit root of 0 is", digSum(0))
	fmt.Println("Digit root of 3 is", digSum(3))
	fmt.Println("Digit root of 9 is", digSum(9))
	fmt.Println("Digit root of 11 is", digSum(11))
	fmt.Println("Digit root of 18 is", digSum(18))
	fmt.Println("Digit root of 20 is", digSum(20))
	fmt.Println("Digit root of 31 is", digSum(31))
	fmt.Println("Digit root of 52 is", digSum(52))
	fmt.Println("Digit root of 80 is", digSum(80))
	fmt.Println("Digit root of 111 is", digSum(111))
	fmt.Println("Digit root of 170 is", digSum(170))
}

func digSum(n int) int {
	if n == 0 {
		return 0
	}

	if n%9 == 0 {
		return 9
	}

	return n % 9
}
