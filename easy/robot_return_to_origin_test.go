package easy_test

import (
	"testing"

	"github.com/scrocrix/leetcode/easy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type judgeCircle struct {
	suite.Suite
}

func (this *judgeCircle) TestJudgeCircle() {
	this.Run("Example 1", func() {
		assert.True(this.T(), easy.JudgeCircle("UD"))
	})

	this.Run("Example 2", func() {
		assert.False(this.T(), easy.JudgeCircle("LL"))
	})

	this.Run("Example 3", func() {
		assert.False(this.T(), easy.JudgeCircle("RRDD"))
	})

	this.Run("Example 4", func() {
		assert.False(this.T(), easy.JudgeCircle("LDRRLRUULR"))
	})
}

func TestJudgeCircle(t *testing.T) {
	suite.Run(t, new(judgeCircle))
}
