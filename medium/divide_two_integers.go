package medium

import "math"

type DivideTwoIntegers struct{}

func (dti *DivideTwoIntegers) Divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}

	var isNegative = false
	if dividend < 0 || divisor < 0 {
		if dividend > 0 || divisor > 0 {
			isNegative = true
		}
	}

	var amount = 0
	var d1 = int(math.Abs(float64(divisor)))
	var d2 = int(math.Abs(float64(dividend)))
	for d2 >= d1 {
		amount += 1
		d2 -= d1
	}

	if amount > 2147483647 {
		return 2147483647
	}

	if amount < -2147483648 {
		return 2147483647
	}

	if isNegative {
		return -amount
	}

	return amount
}
