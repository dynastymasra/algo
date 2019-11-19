package basic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dynastymasra/algo/basic"
)

func Test_LCM(t *testing.T) {
	lcm1 := basic.LCM(23, 40)
	lcm2 := basic.LCM(23, 40, 50)
	lcm3 := basic.LCM(23, 40, 50, 71)
	lcm4 := basic.LCM(6, 7, 21)
	lcm5 := basic.LCM(12, 30)

	assert.Equal(t, 920, lcm1)
	assert.Equal(t, 4600, lcm2)
	assert.Equal(t, 326600, lcm3)
	assert.Equal(t, 42, lcm4)
	assert.Equal(t, 60, lcm5)
}
