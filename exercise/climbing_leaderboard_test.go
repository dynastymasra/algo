package exercise_test

import (
	"testing"

	"github.com/dynastymasra/algo/exercise"
	"github.com/stretchr/testify/assert"
)

func Test_ClimbingLeaderBoard_Test_1(t *testing.T) {
	scores := []int{100, 100, 50, 40, 40, 20, 10}
	input := []int{5, 25, 50, 120}
	output := []int{6, 4, 2, 1}

	res := exercise.ClimbingLeaderBoard(scores, input)

	assert.NotEmpty(t, res)
	assert.Equal(t, output, res)
}

func Test_ClimbingLeaderBoard_Test_2(t *testing.T) {
	scores := []int{100, 90, 90, 80, 75, 60}
	input := []int{50, 65, 77, 90, 102}
	output := []int{6, 5, 4, 2, 1}

	res := exercise.ClimbingLeaderBoard(scores, input)

	assert.NotEmpty(t, res)
	assert.Equal(t, output, res)
}

func Test_ClimbingLeaderBoard_Test_3(t *testing.T) {
	scores := []int{100, 100, 50, 40, 40, 20, 10}
	input := []int{5, 10, 25, 50, 120}
	output := []int{6, 5, 4, 2, 1}

	res := exercise.ClimbingLeaderBoard(scores, input)

	assert.NotEmpty(t, res)
	assert.Equal(t, output, res)
}
