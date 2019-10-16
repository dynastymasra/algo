package main

import "fmt"

func main() {
	fmt.Println("Fibonacci 1 is", fibonacci(1), "and sum is", fibonacci(1+2)-1)
	fmt.Println("Fibonacci 2 is", fibonacci(2), "and sum is", fibonacci(2+2)-1)
	fmt.Println("Fibonacci 5 is", fibonacci(5), "and sum is", fibonacci(5+2)-1)
	fmt.Println("Fibonacci 9 is", fibonacci(9), "and sum is", fibonacci(9+2)-1)
	fmt.Println("Fibonacci 14 is", fibonacci(14), "and sum is", fibonacci(14+2)-1)
	fmt.Println("Fibonacci 20 is", fibonacci(20), "and sum is", fibonacci(20+2)-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
