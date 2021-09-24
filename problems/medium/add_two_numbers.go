package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

/*func main() {
	result1 := addTwoNumbers(&ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}, &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	})

	result2 := addTwoNumbers(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	})

	for result1.Next != nil {
		fmt.Println(result1.Val)
		result1 = result1.Next
	}

	fmt.Println(result1.Val)

	for result2.Next != nil {
		fmt.Println(result2.Val)
		result2 = result2.Next
	}

	fmt.Println(result2.Val)
}*/
