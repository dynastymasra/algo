package main

import (
	"fmt"
	"math"
)

/**
Ref:
https://www.geeksforgeeks.org/finding-sum-of-digits-of-a-number-until-sum-becomes-single-digit/
https://mathlair.allfunandgames.ca/digroots.php

Digital roots are based on the concept of numerical congruence formulated by Carl Friedrich Gauss. The largest digit in the decimal number system is 9, and so the sum of the digits of any number are always congruent modulo 9 to that number. This property can be used to show whether any number is divisible by 9. Obtaining the digital root is simply the ancient process of "casting out nines". For example, the number 3,128 has a digital root of 5 because its digits sum to 14, and the digits of 14 add to 5. It is the same as the remainder obtained when dividing the number by 9 (unless the number is divisible by 9, in which case the digital root is 9 instead of 0).
One way of quickly checking whether a sum involving large numbers is correct is to take the digital roots of the numbers, add them, reduce the answer to a digital root, and then see if it corresponds to the digital root of the answer. If they don't match, the answer is wrong. If they do match, the probability is fairly high that the answer is correct. A similar idea (except in binary instead of decimal) is used in computers for parity checking.
A knowledge of digital roots often enables shortcuts in solving otherwise difficult problems.
Example: Find the smallest natural number that is composed entirely of 1's and 0's and is evenly divisible by 225.
Answer: Since the digits in 225 have a digital root of 9, you know at once that the answer must also have a digital root of 9. The smallest such number is 111,111,111. The problem is to increase 111,111,111 by the smallest amount that will make it divisible by 225. Since 225 is a multiple of 25, the number we seek must also be a multiple of 25, and so the number must end in 00, 25, 50, or 75. Since the last three pairs can't be used, we attach 00 to 111,111,111 to obtain the answer: 11,111,111,100.
*/
func main() {
	fmt.Println("Digit root of -170 is", digSum1(-170), digSum2(-170))
	fmt.Println("Digit root of -111 is", digSum1(-111), digSum2(-111))
	fmt.Println("Digit root of -80 is", digSum1(-80), digSum2(-80))
	fmt.Println("Digit root of -52 is", digSum1(-52), digSum2(-52))
	fmt.Println("Digit root of -31 is", digSum1(-31), digSum2(-31))
	fmt.Println("Digit root of -20 is", digSum1(-20), digSum2(-20))
	fmt.Println("Digit root of -18 is", digSum1(-18), digSum2(-18))
	fmt.Println("Digit root of -11 is", digSum1(-11), digSum2(-11))
	fmt.Println("Digit root of -9 is", digSum1(-9), digSum2(-9))
	fmt.Println("Digit root of -3 is", digSum1(-3), digSum2(-3))
	fmt.Println("Digit root of 0 is", digSum1(0), digSum2(0))
	fmt.Println("Digit root of 3 is", digSum1(3), digSum2(3))
	fmt.Println("Digit root of 9 is", digSum1(9), digSum2(9))
	fmt.Println("Digit root of 11 is", digSum1(11), digSum2(11))
	fmt.Println("Digit root of 18 is", digSum1(18), digSum2(18))
	fmt.Println("Digit root of 20 is", digSum1(20), digSum2(20))
	fmt.Println("Digit root of 31 is", digSum1(31), digSum2(31))
	fmt.Println("Digit root of 52 is", digSum1(52), digSum2(52))
	fmt.Println("Digit root of 80 is", digSum1(80), digSum2(80))
	fmt.Println("Digit root of 111 is", digSum1(111), digSum2(111))
	fmt.Println("Digit root of 170 is", digSum1(170), digSum2(170))
}

func digSum1(x int) int {
	if x == 0 {
		return 0
	}

	if x%9 == 0 {
		if x < 0 {
			return -(9)
		}
		return 9
	}

	return x % 9
}

func digSum2(x int) int {
	if x > 9 || x < -9 {
		digRoot := 0

		for math.Abs(float64(x)) > 0 {
			digRoot += x % 10
			x = x / 10
		}
		return digRoot
	}

	return x
}
