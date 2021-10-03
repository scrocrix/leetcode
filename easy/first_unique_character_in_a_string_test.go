package easy_test

import (
	"testing"

	"github.com/scrocrix/leetcode/easy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type firstUniqChar struct {
	suite.Suite
}

func (this *firstUniqChar) TestFirstUniqChar() {
	this.Run("Example 1", func() {
		result := easy.FirstUniqChar("leetcode")
		assert.Equal(this.T(), 0, result)
	})

	this.Run("Example 2", func() {
		result := easy.FirstUniqChar("loveleetcode")
		assert.Equal(this.T(), 2, result)
	})

	this.Run("Example 3", func() {
		result := easy.FirstUniqChar("aabb")
		assert.Equal(this.T(), -1, result)
	})
}

func TestFirstUniqChar(t *testing.T) {
	suite.Run(t, new(firstUniqChar))
}
