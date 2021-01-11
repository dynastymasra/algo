package exercise

import (
	"fmt"
	"strings"
)

func AddTwoBinary(a, b string) string {
	result := ""
	s := 0

	bA := strings.Split(a, "")
	bB := strings.Split(b, "")

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || s == 1 {
		if i >= 0 && bA[i] == "1" {
			s++
		}

		if j >= 0 && bB[j] == "1" {
			s++
		}

		if s == 1 {
			result = fmt.Sprintf("%s%s", "1", result)
			s = 0
		} else if s == 2 {
			result = fmt.Sprintf("%s%s", "0", result)
			s = 1
		} else if s == 3 {
			result = fmt.Sprintf("%s%s", "1", result)
			s = 1
		} else {
			result = fmt.Sprintf("%s%s", "0", result)
			s = 0
		}

		i--
		j--
	}

	res := result
	for _, s := range strings.Split(res, "") {
		if s == "0" {
			result = result[1:]
		} else {
			break
		}
	}

	return result
}
