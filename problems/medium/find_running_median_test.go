package medium_test

import (
	"testing"

	"github.com/scrocrix/leetcode/problems/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type runningMedian struct {
	suite.Suite
}

func (this *runningMedian) TestRunningMedian() {
	this.Run("Example one", func() {
		result := medium.RunningMedian([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		require.Equal(this.T(), []float64{
			1.0,
			1.5,
			2.0,
			2.5,
			3.0,
			3.5,
			4.0,
			4.5,
			5.0,
			5.5,
		}, result)
	})

	this.Run("Example two", func() {
		result := medium.RunningMedian([]int32{7, 3, 5, 2})
		require.Equal(this.T(), []float64{
			7.0,
			5.0,
			5.0,
			4.0,
		}, result)
	})
}

func TestRunningMedian(t *testing.T) {
	suite.Run(t, new(runningMedian))
}
