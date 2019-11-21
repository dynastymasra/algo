package basic_test

import (
	"testing"

	"github.com/dynastymasra/algo/basic"
	"github.com/stretchr/testify/assert"
)

func Test_IsPowerLoop_2(t *testing.T) {
	res1 := basic.IsPowerLoop(2, 2)
	res2 := basic.IsPowerLoop(2, 0)
	res3 := basic.IsPowerLoop(2, 1)
	res4 := basic.IsPowerLoop(2, 10)
	res5 := basic.IsPowerLoop(2, 25)
	res6 := basic.IsPowerLoop(2, 256)
	res7 := basic.IsPowerLoop(2, 1000000)
	res8 := basic.IsPowerLoop(2, 20)
	res9 := basic.IsPowerLoop(2, 4194304)
	res10 := basic.IsPowerLoop(2, 8192)

	assert.True(t, res1)
	assert.False(t, res2)
	assert.True(t, res3)
	assert.False(t, res4)
	assert.False(t, res5)
	assert.True(t, res6)
	assert.False(t, res7)
	assert.False(t, res8)
	assert.True(t, res9)
	assert.True(t, res10)
}

func Test_PowerOfTwo(t *testing.T) {
	res1 := basic.IsPowerLoop(2, 2)
	res2 := basic.IsPowerLoop(2, 0)
	res3 := basic.IsPowerLoop(2, 1)
	res4 := basic.IsPowerLoop(2, 10)
	res5 := basic.IsPowerLoop(2, 25)
	res6 := basic.IsPowerLoop(2, 256)
	res7 := basic.IsPowerLoop(2, 1000000)
	res8 := basic.IsPowerLoop(2, 20)
	res9 := basic.IsPowerLoop(2, 4194304)
	res10 := basic.IsPowerLoop(2, 8192)

	assert.True(t, res1)
	assert.False(t, res2)
	assert.True(t, res3)
	assert.False(t, res4)
	assert.False(t, res5)
	assert.True(t, res6)
	assert.False(t, res7)
	assert.False(t, res8)
	assert.True(t, res9)
	assert.True(t, res10)
}

func Test_IsPowerLoop_4(t *testing.T) {
	res1 := basic.IsPowerLoop(4, 4)
	res2 := basic.IsPowerLoop(4, 0)
	res3 := basic.IsPowerLoop(4, 1)
	res4 := basic.IsPowerLoop(4, 16)
	res5 := basic.IsPowerLoop(4, 40)
	res6 := basic.IsPowerLoop(4, 65536)
	res7 := basic.IsPowerLoop(4, 1000000)
	res8 := basic.IsPowerLoop(4, 20)
	res9 := basic.IsPowerLoop(4, 1048576)
	res10 := basic.IsPowerLoop(4, 4096)

	assert.True(t, res1)
	assert.False(t, res2)
	assert.True(t, res3)
	assert.True(t, res4)
	assert.False(t, res5)
	assert.True(t, res6)
	assert.False(t, res7)
	assert.False(t, res8)
	assert.True(t, res9)
	assert.True(t, res10)
}

func Test_IsPowerLoop_5(t *testing.T) {
	res1 := basic.IsPowerLoop(5, 5)
	res2 := basic.IsPowerLoop(5, 0)
	res3 := basic.IsPowerLoop(5, 1)
	res4 := basic.IsPowerLoop(5, 25)
	res5 := basic.IsPowerLoop(5, 50)
	res6 := basic.IsPowerLoop(5, 3125)
	res7 := basic.IsPowerLoop(5, 1000000)
	res8 := basic.IsPowerLoop(5, 20)
	res9 := basic.IsPowerLoop(5, 1953125)
	res10 := basic.IsPowerLoop(5, 625)

	assert.True(t, res1)
	assert.False(t, res2)
	assert.True(t, res3)
	assert.True(t, res4)
	assert.False(t, res5)
	assert.True(t, res6)
	assert.False(t, res7)
	assert.False(t, res8)
	assert.True(t, res9)
	assert.True(t, res10)
}
