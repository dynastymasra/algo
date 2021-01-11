package exercise

import (
	"fmt"
	"math"
)

func CashRegister(change float64) string {
	res := ""

	for change >= 0.01 {
		if change >= 100 {
			res = fmt.Sprintf("%s,%s", res, "ONE HUNDRED")
			change -= 100
		} else if change >= 50 {
			res = fmt.Sprintf("%s,%s", res, "FIFTY")
			change -= 50
		} else if change >= 20 {
			res = fmt.Sprintf("%s,%s", res, "TWENTY")
			change -= 20
		} else if change >= 10 {
			res = fmt.Sprintf("%s,%s", res, "TEN")
			change -= 10
		} else if change >= 5 {
			res = fmt.Sprintf("%s,%s", res, "FIVE")
			change -= 5
		} else if change >= 2 {
			res = fmt.Sprintf("%s,%s", res, "TWO")
			change -= 2
		} else if change >= 1 {
			res = fmt.Sprintf("%s,%s", res, "ONE")
			change -= 1
		} else if change >= 0.50 {
			res = fmt.Sprintf("%s,%s", res, "HALF DOLLAR")
			change -= .50
		} else if change >= 0.25 {
			res = fmt.Sprintf("%s,%s", res, "QUARTER")
			change -= .25
		} else if change >= 0.10 {
			res = fmt.Sprintf("%s,%s", res, "DIME")
			change -= .10
		} else if change >= 0.05 {
			res = fmt.Sprintf("%s,%s", res, "NICKEL")
			change -= .05
		} else {
			res = fmt.Sprintf("%s,%s", res, "PENNY")
			change -= .01
		}
		change = math.Round(change*100) / 100
	}

	return res[1:]
}
