package exercise_test

import (
	"testing"

	"github.com/dynastymasra/algo/exercise"
	"github.com/stretchr/testify/assert"
)

func Test_CashRegister(t *testing.T) {
	res := exercise.CashRegister(.06)
	assert.Equal(t, "NICKEL,PENNY", res)

	res = exercise.CashRegister(100.50)
	assert.Equal(t, "ONE HUNDRED,HALF DOLLAR", res)

	res = exercise.CashRegister(51.10)
	assert.Equal(t, "FIFTY,ONE,DIME", res)

	res = exercise.CashRegister(25.25)
	assert.Equal(t, "TWENTY,FIVE,QUARTER", res)

	res = exercise.CashRegister(12.12)
	assert.Equal(t, "TEN,TWO,DIME,PENNY,PENNY", res)
}
