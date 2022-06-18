package medium

import "math"

type DivideTwoIntegers struct{}

// Divide is a less efficient solution where where it runs O(n) time.
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

// DivideOptimal bases the amount of iteration on the logarithm of 32-byte integers
// making it O(log n)
func (dti *DivideTwoIntegers) DivideOptimal(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	var dividendSign = 1
	var divisorSign = 1
	if dividend < 0 {
		dividend = -dividend
		dividendSign = 0
	}

	if divisor < 0 {
		divisor = -divisor
		divisorSign = 0
	}

	res := 0
	for i := 31; i >= 0; i-- {
		if dividend >= (divisor << i) {
			res += 1 << i
			dividend -= (divisor << i)
		}
	}

	if dividendSign^divisorSign == 1 {
		res = -res
	}

	return res
}
