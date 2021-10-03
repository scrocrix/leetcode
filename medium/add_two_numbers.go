package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var numbers *ListNode

	for l1 != nil || l2 != nil {
		var vl1, vl2 int

		if l1 != nil {
			vl1 = l1.Val
		} else {
			vl1 = 0
		}

		if l2 != nil {
			vl2 = l2.Val
		} else {
			vl2 = 0
		}

		// calculates the number
		sum := vl1 + vl2 + carry
		carry = sum / 10
		digit := sum % 10

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}

		// append result of sum to linked list
		if numbers == nil {
			numbers = &ListNode{
				Val: digit,
			}
		} else {
			cn := numbers
			for cn.Next != nil {
				cn = cn.Next
			}
			cn.Next = &ListNode{
				Val: digit,
			}
		}

		// any rest of the calculation gets carried on
		if l1 == nil && l2 == nil && carry > 0 {
			cn := numbers
			for cn.Next != nil {
				cn = cn.Next
			}
			cn.Next = &ListNode{
				Val: carry,
			}
		}
	}

	return numbers
}
