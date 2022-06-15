package medium_test

import (
	"github.com/scrocrix/leetcode/medium"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type reverseIntegerTest struct {
	suite.Suite
}

func (unit *reverseIntegerTest) TestReverseReturnRegularInteger() {
	sut := &medium.ReverseInteger{}

	assert.Equal(unit.T(), 321, sut.Reverse(123))
	assert.Equal(unit.T(), 0, sut.Reverse(1534236469))
}

func (unit *reverseIntegerTest) TestReverseReturnNegativeReverse() {
	sut := &medium.ReverseInteger{}

	assert.Equal(unit.T(), -321, sut.Reverse(-123))
}

func (unit *reverseIntegerTest) TestReverseReturnNumberIgnoringZeros() {
	sut := &medium.ReverseInteger{}

	assert.Equal(unit.T(), 21, sut.Reverse(120))
}

func TestReverseInteger(t *testing.T) {
	suite.Run(t, new(reverseIntegerTest))
}
