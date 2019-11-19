package main

import "fmt"

func main() {
	fmt.Println("Factorial 5 is", factorial(5))
	fmt.Println("Factorial 7 is", factorial(7))
	fmt.Println("Factorial 10 is", factorial(10))
	fmt.Println("Factorial 16 is", factorial(16))
	fmt.Println("Factorial 16 is", factorial(20))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
