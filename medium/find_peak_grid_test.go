package medium_test

import (
	"testing"

	"github.com/scrocrix/leetcode/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type findPeakGrid struct {
	suite.Suite
}

func (this *findPeakGrid) TestFindPeakGrid() {
	this.Run("Example 1", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{1, 4},
			[]int{3, 2},
		})

		require.Equal(this.T(), []int{1, 0}, result)
	})

	this.Run("Example 2", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{10, 20, 15},
			[]int{21, 30, 14},
			[]int{7, 16, 32},
		})

		require.Equal(this.T(), []int{2, 2}, result)
	})

	this.Run("Example 3", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{20, 43, 38, 24, 31},
			[]int{36, 34, 23, 28, 48},
			[]int{22, 23, 45, 30, 18},
			[]int{20, 17, 15, 8, 47},
			[]int{13, 21, 8, 48, 35},
			[]int{49, 45, 12, 13, 14},
			[]int{48, 31, 18, 47, 38},
			[]int{49, 34, 39, 19, 7},
		})

		require.Equal(this.T(), []int{7, 2}, result)
	})

	this.Run("Example 4", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{11, 77, 6, 56, 95, 42, 78, 33, 4, 12, 30, 87, 74, 76, 3, 98},
			[]int{87, 35, 74, 6, 98, 97, 23, 61, 67, 84, 57, 36, 24, 75, 56, 82},
			[]int{63, 12, 61, 45, 53, 43, 42, 86, 14, 21, 43, 53, 89, 64, 37, 85},
			[]int{52, 98, 71, 5, 59, 12, 41, 33, 42, 15, 99, 52, 66, 1, 85, 70},
			[]int{60, 53, 7, 25, 67, 47, 80, 97, 64, 81, 57, 98, 7, 27, 35, 71},
			[]int{17, 37, 57, 90, 63, 16, 4, 81, 78, 31, 91, 21, 40, 71, 2, 53},
			[]int{12, 29, 15, 60, 16, 83, 20, 64, 92, 38, 66, 91, 35, 54, 62, 61},
			[]int{32, 59, 71, 19, 50, 25, 77, 66, 96, 18, 9, 36, 99, 44, 65, 35},
			[]int{1, 37, 93, 39, 76, 75, 85, 63, 20, 58, 69, 62, 45, 43, 71, 98},
			[]int{70, 49, 59, 42, 25, 90, 31, 74, 50, 30, 100, 53, 6, 41, 96, 61},
			[]int{91, 50, 85, 84, 20, 64, 5, 17, 65, 76, 43, 13, 57, 13, 68, 59},
			[]int{2, 93, 19, 53, 84, 20, 10, 61, 34, 44, 80, 89, 51, 84, 11, 43},
			[]int{87, 69, 49, 41, 81, 92, 98, 63, 70, 5, 9, 31, 81, 88, 37, 30},
			[]int{96, 12, 56, 13, 98, 82, 91, 13, 57, 22, 89, 39, 4, 6, 42, 8},
			[]int{38, 46, 48, 56, 19, 32, 68, 53, 73, 88, 24, 22, 23, 11, 79, 49},
			[]int{80, 88, 31, 64, 75, 19, 30, 60, 82, 33, 59, 33, 86, 2, 60, 37},
			[]int{87, 70, 21, 10, 57, 45, 76, 39, 55, 89, 33, 59, 39, 26, 50, 62},
		})

		require.Equal(this.T(), []int{16, 15}, result)
	})

	this.Run("Example 5", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{25, 37, 23, 37, 19},
			[]int{45, 19, 2, 43, 26},
			[]int{18, 1, 37, 44, 50},
		})

		require.Equal(this.T(), []int{2, 4}, result)
	})

	this.Run("Example 6", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{10, 30, 40, 50, 20},
			[]int{1, 3, 2, 500, 4},
		})

		require.Equal(this.T(), []int{1, 3}, result)
	})

	this.Run("Example 7", func() {
		result := medium.FindPeakGrid([][]int{
			[]int{10, 50, 40, 30, 20},
			[]int{1, 500, 2, 3, 4},
		})

		require.Equal(this.T(), []int{1, 1}, result)
	})
}

func TestFindPeakGrid(t *testing.T) {
	suite.Run(t, new(findPeakGrid))
}
