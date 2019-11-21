package basic

import "fmt"

func IsPowerLoop(x, n int) bool {
	pow := n
	if pow > 0 {
		for pow%x == 0 {
			pow = pow / x
		}
	}

	if pow == 1 {
		fmt.Printf("%d is power of %d\n", n, x)
		return true
	} else {
		fmt.Printf("%d is not power of %d\n", n, x)
		return false
	}
}

/** PowerOfTwo bit manipulation
x = 4 = (100)2
x - 1 = 3 = (011)2
x & (x-1) = 4 & 3 = (100)2 & (011)2 = (000)2

x = 6 = (110)2
x - 1 = 5 = (101)2
x & (x-1) = 6 & 5 = (110)2 & (101)2 = (100)2
*/
func PowerOfTwo(n int) bool {
	return n != 0 && n&(n-1) == 0
}
