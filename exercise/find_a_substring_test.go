package exercise_test

import (
	"testing"

	"github.com/dynastymasra/algo/exercise"
	"github.com/stretchr/testify/assert"
)

func Test_FindASubstring(t *testing.T) {
	res := exercise.FindASubstring([]string{"a", "b", "c", "d"}, "tbcacbadcaik")
	assert.Equal(t, 4, res)

	res = exercise.FindASubstring([]string{"a", "b", "c", "d"}, "tbcacbkdcaik")
	assert.Equal(t, -1, res)

	res = exercise.FindASubstring([]string{"a", "b", "c"}, "tbcacbkdcaik")
	assert.Equal(t, 1, res)

	res = exercise.FindASubstring([]string{"a", "b", "c"}, "abcacbkdcaik")
	assert.Equal(t, 0, res)

	res = exercise.FindASubstring([]string{"a", "b", "c"}, "aabkdcaab")
	assert.Equal(t, -1, res)

	res = exercise.FindASubstring([]string{"a", "b", "c"}, "aabkdcaabnacb")
	assert.Equal(t, 10, res)

	res = exercise.FindASubstring([]string{"a", "b", "c"}, "aabkdcaacbnac")
	assert.Equal(t, 7, res)
}
