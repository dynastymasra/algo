package main

import "fmt"

func main() {
	arr := []int{9, 6, 4, 2, 5, 7, 1, 8, 11, 3}

	fmt.Println("Missing number is", missingNumber(arr))
}

func missingNumber(arr []int) int {
	size := len(arr) + 1
	total := size * (size + 1) / 2

	var a int
	for _, i := range arr {
		a += i
	}

	return total - a
}
