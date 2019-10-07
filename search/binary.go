package main

import "fmt"

func main() {
	arr := []int{2, 5, 8, 10, 17, 23, 26, 40, 51, 73, 80, 83, 91, 99, 100, 101}

	fmt.Println("Number 80 is found in array", binarySearch(arr, 80))
	fmt.Println("Number 10 is found in array", binarySearch(arr, 10))
	fmt.Println("Number 17 is found in array", binarySearch(arr, 17))
	fmt.Println("Number 15 is found in array", binarySearch(arr, 15))
	fmt.Println("Number 95 is found in array", binarySearch(arr, 95))
}

func binarySearch(arr []int, x int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] == x {
			return true
		}

		if arr[median] < x {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return false
}
