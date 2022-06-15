package medium

import (
	"strconv"
	"strings"
)

type ReverseInteger struct{}

func (ri *ReverseInteger) Reverse(x int) int {
	var s = strings.Split(strconv.Itoa(x), "")

	if x < 0 {
		s = s[1:]
	}

	var l = 0
	var r = len(s) - 1
	for l < r {
		var e = s[l]
		var f = s[r]
		s[l] = f
		s[r] = e
		l++
		r--
	}

	var i int
	if x < 0 {
		i, _ = strconv.Atoi("-" + strings.Join(s, ""))
	}
	if x > 0 {
		i, _ = strconv.Atoi(strings.Join(s, ""))
	}

	if i >= 2147483648-1 || i <= -2147483648 {
		return 0
	}

	return i
}
