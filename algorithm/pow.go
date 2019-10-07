package main

import "fmt"

/**
Ref: https://www.geeksforgeeks.org/write-a-c-program-to-calculate-powxn/
*/
func main() {
	fmt.Println("2 power 5 is ", power(2, 5))
}

func power(x, y int) int {
	if y == 0 {
		return 1
	}

	temp := power(x, y/2)

	if y%2 == 0 {
		return temp * temp
	}

	return x * temp * temp
}
