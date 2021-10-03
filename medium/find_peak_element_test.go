package medium_test

import (
	"testing"

	"github.com/scrocrix/leetcode/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type findPeakElement struct {
	suite.Suite
}

func (this *findPeakElement) TestFindPeakElement() {
	this.Run("Example 1", func() {
		result := medium.FindPeakElement([]int{1, 2, 3, 1})
		require.Equal(this.T(), 2, result)
	})

	this.Run("Example 2", func() {
		result := medium.FindPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
		require.Equal(this.T(), 5, result)
	})

	this.Run("Example 3", func() {
		result := medium.FindPeakElement([]int{1, 2})
		require.Equal(this.T(), 1, result)
	})
}

func TestFindPeakElement(t *testing.T) {
	suite.Run(t, new(findPeakElement))
}
