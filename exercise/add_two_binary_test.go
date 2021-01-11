package exercise_test

import (
	"testing"

	"github.com/dynastymasra/algo/exercise"
	"github.com/stretchr/testify/assert"
)

func Test_AddTwoBinary(t *testing.T) {
	res := exercise.AddTwoBinary("1101", "100")
	assert.Equal(t, "10001", res)

	res = exercise.AddTwoBinary("1101", "1100")
	assert.Equal(t, "11001", res)

	res = exercise.AddTwoBinary("001101", "1100")
	assert.Equal(t, "11001", res)

	res = exercise.AddTwoBinary("1101", "001100")
	assert.Equal(t, "11001", res)
}
