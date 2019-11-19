package basic_test

import (
	"testing"

	"github.com/dynastymasra/algo/basic"
	"github.com/stretchr/testify/assert"
)

func Test_GCD(t *testing.T) {
	gcd1 := basic.GCD(54, 24)
	gcd2 := basic.GCD(252, 105)

	assert.Equal(t, 6, gcd1)
	assert.Equal(t, 21, gcd2)
}

func Test_GCD_GO(t *testing.T) {
	gcd1 := basic.GoGCD(54, 24)
	gcd2 := basic.GoGCD(252, 105)
	gcd3 := basic.GoGCD(18, 27)

	assert.Equal(t, int64(6), gcd1)
	assert.Equal(t, int64(21), gcd2)
	assert.Equal(t, int64(9), gcd3)
}
