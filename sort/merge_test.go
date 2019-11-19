package sort_test

import (
	"testing"

	"github.com/dynastymasra/algo/sort"
	"github.com/stretchr/testify/assert"
)

func Test_Merge(t *testing.T) {
	items1 := []int{14, 33, 27, 10, 35, 19, 42, 44}
	expected1 := []int{10, 14, 19, 27, 33, 35, 42, 44}

	items2 := []int{38, 27, 43, 3, 9, 82, 10}
	expected2 := []int{3, 9, 10, 27, 38, 43, 82}

	items3 := []int{3, 4, 2, 1, 7, 5, 8, 9, 0, 6}
	expected3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sorted1 := sort.Merge(items1)
	sorted2 := sort.Merge(items2)
	sorted3 := sort.Merge(items3)

	assert.Equal(t, expected1, sorted1)
	assert.Equal(t, expected2, sorted2)
	assert.Equal(t, expected3, sorted3)
}
