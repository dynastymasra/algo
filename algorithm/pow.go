package main

import "fmt"

func main() {
	fmt.Println("2 power 2 is ", power(2, 4))
}

func power(x, y int) int {
	if y == 0 {
		return 1
	} else if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	} else {
		return x * power(x, y/2) * power(x, y/2)
	}
}
