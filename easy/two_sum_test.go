package easy_test

import (
	"testing"

	"github.com/scrocrix/leetcode/easy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type twoSum struct {
	suite.Suite
}

func (this *twoSum) TestTwoSum() {
	this.Run("Example 1", func() {
		result := easy.TwoSum([]int{2, 7, 11, 15}, 9)
		assert.Equal(this.T(), []int{0, 1}, result)
	})

	this.Run("Example 2", func() {
		result := easy.TwoSum([]int{3, 2, 4}, 6)
		assert.Equal(this.T(), []int{1, 2}, result)
	})

	this.Run("Example 3", func() {
		result := easy.TwoSum([]int{3, 3}, 6)
		assert.Equal(this.T(), []int{0, 1}, result)
	})
}

func TestTwoSum(t *testing.T) {
	suite.Run(t, new(twoSum))
}
