package sort

func Merge(items []int) []int {
	size := len(items)

	if size <= 1 {
		return items
	}

	mid := size / 2

	left := items[:mid]
	right := items[mid:]

	return merge(Merge(left), Merge(right))
}

func merge(left, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for i := 0; i < len(left); i++ {
		result = append(result, left[i])
	}

	for i := 0; i < len(right); i++ {
		result = append(result, right[i])
	}

	return result
}
