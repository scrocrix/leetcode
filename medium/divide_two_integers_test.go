package medium_test

import (
	"github.com/scrocrix/leetcode/medium"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type divideTwoIntegersTest struct {
	suite.Suite
}

func (unit *divideTwoIntegersTest) TestDivideReturnTruncatedInteger() {
	sut := &medium.DivideTwoIntegers{}

	assert.Equal(unit.T(), 3, sut.Divide(10, 3))
	assert.Equal(unit.T(), -2, sut.Divide(7, -3))
	assert.Equal(unit.T(), -1, sut.Divide(-1, 1))
	assert.Equal(unit.T(), 1, sut.Divide(-1, -1))
	assert.Equal(unit.T(), 2147483647, sut.Divide(-2147483648, -1))
	assert.Equal(unit.T(), -2147483648, sut.Divide(-2147483648, 1))
	assert.Equal(unit.T(), -1073741824, sut.Divide(-2147483648, 2))
}

func TestDivideTwoInteger(t *testing.T) {
	suite.Run(t, new(divideTwoIntegersTest))
}
