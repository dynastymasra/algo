package exercise_test

import (
	"testing"

	"github.com/dynastymasra/algo/exercise"
	"github.com/stretchr/testify/assert"
)

func Test_PingPongMatch_1(t *testing.T) {
	t.Parallel()

	time := exercise.PingPongTime{
		Min: 9,
		Max: 18,
	}
	slot := []exercise.PingPongSlot{
		{
			Start: 9,
			End:   10,
		}, {
			Start: 9,
			End:   12,
		},
	}

	single, double := exercise.PingPongMatch(time, slot)

	assert.Equal(t, 1, single)
	assert.Equal(t, 0, double)
}

func Test_PingPongMatch_2(t *testing.T) {
	t.Parallel()

	time := exercise.PingPongTime{
		Min: 9,
		Max: 18,
	}
	slot := []exercise.PingPongSlot{
		{
			Start: 9,
			End:   12,
		}, {
			Start: 13,
			End:   15,
		}, {
			Start: 10,
			End:   18,
		},
	}

	single, double := exercise.PingPongMatch(time, slot)

	assert.Equal(t, 4, single)
	assert.Equal(t, 0, double)
}

func Test_PingPongMatch_3(t *testing.T) {
	t.Parallel()

	time := exercise.PingPongTime{
		Min: 9,
		Max: 18,
	}
	slot := []exercise.PingPongSlot{
		{
			Start: 9,
			End:   11,
		}, {
			Start: 12,
			End:   14,
		}, {
			Start: 13,
			End:   15,
		}, {
			Start: 11,
			End:   18,
		}, {
			Start: 14,
			End:   18,
		}, {
			Start: 10,
			End:   14,
		}, {
			Start: 15,
			End:   18,
		},
	}

	single, double := exercise.PingPongMatch(time, slot)

	assert.Equal(t, 7, single)
	assert.Equal(t, 1, double)
}

func Test_PingPongMatch_4(t *testing.T) {
	t.Parallel()

	time := exercise.PingPongTime{
		Min: 9,
		Max: 18,
	}

	slot := []exercise.PingPongSlot{
		{
			Start: 9,
			End:   10,
		}, {
			Start: 9,
			End:   12,
		}, {
			Start: 14,
			End:   15,
		}, {
			Start: 16,
			End:   18,
		}, {
			Start: 14,
			End:   18,
		}, {
			Start: 10,
			End:   13,
		}, {
			Start: 15,
			End:   18,
		}, {
			Start: 10,
			End:   14,
		}, {
			Start: 9,
			End:   13,
		},
	}

	single, double := exercise.PingPongMatch(time, slot)

	assert.Equal(t, 6, single)
	assert.Equal(t, 2, double)
}

func Test_PingPongMatch_5(t *testing.T) {
	t.Parallel()

	time := exercise.PingPongTime{
		Min: 9,
		Max: 18,
	}

	slot := []exercise.PingPongSlot{
		{
			Start: 9,
			End:   14,
		}, {
			Start: 15,
			End:   17,
		}, {
			Start: 12,
			End:   13,
		}, {
			Start: 16,
			End:   18,
		}, {
			Start: 10,
			End:   15,
		}, {
			Start: 14,
			End:   16,
		}, {
			Start: 9,
			End:   13,
		}, {
			Start: 9,
			End:   15,
		},
	}

	single, double := exercise.PingPongMatch(time, slot)

	assert.Equal(t, 5, single)
	assert.Equal(t, 3, double)
}
