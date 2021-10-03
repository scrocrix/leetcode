package medium_test

import (
	"testing"

	"github.com/scrocrix/leetcode/medium"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type addTwoNumbers struct {
	suite.Suite
}

func (this *addTwoNumbers) TestAddTwoNumbers() {
	this.Run("Example 1", func() {
		result := medium.AddTwoNumbers(&medium.ListNode{
			Val: 2,
			Next: &medium.ListNode{
				Val: 4,
				Next: &medium.ListNode{
					Val: 3,
				},
			},
		}, &medium.ListNode{
			Val: 5,
			Next: &medium.ListNode{
				Val: 6,
				Next: &medium.ListNode{
					Val: 4,
				},
			},
		})

		assert.Equal(this.T(), 7, result.Val)
		assert.Equal(this.T(), 0, result.Next.Val)
		assert.Equal(this.T(), 8, result.Next.Next.Val)
	})
}

func TestAddTwoNumbers(t *testing.T) {
	suite.Run(t, new(addTwoNumbers))
}
