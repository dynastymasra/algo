package basic

func LCM(a, b int, values ...int) int {
	result := a * b / GCD(a, b)

	for _, val := range values {
		result = LCM(result, val)
	}

	return result
}
