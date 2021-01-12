package exercise

import (
	"strings"
)

// FindASubstring [a, b, c, d] in string tbcacbadcaik
// return -1 if not found return position if found
func FindASubstring(a []string, m string) int {
	ma := strings.Split(m, "")
	tempPos := -1
	tempLen := 0
	for i := range ma {
		tempA := flushData(a)

		for j := i; j < len(ma); j++ {
			if tempLen == len(a) {
				return tempPos
			}

			if v, ok := tempA[ma[j]]; ok {
				if v > 0 {
					tempPos = -1
					tempLen = 0
					break
				}

				if tempPos < 0 {
					tempPos = j
				}

				tempA[ma[j]] = 1
				tempLen++
			} else {
				tempPos = -1
				tempLen = 0
				break
			}
		}

		if tempLen == len(a) {
			return tempPos
		}

		tempPos = -1
		tempLen = 0
	}

	return tempPos
}

func flushData(a []string) map[string]int {
	tempA := make(map[string]int)

	for _, v := range a {
		tempA[v] = 0
	}

	return tempA
}
