package medium

func FindPeakElement(nums []int) int {
	p := 0

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			p = 0
		} else {
			p = 1
		}
	}

	if len(nums) > 1 {
		for i, c := range nums {
			if i == 0 {
				n2 := nums[i+1]
				if c > n2 {
					p = i
				}
				continue
			}

			if i == len(nums)-1 {
				b := nums[i-1]
				if c > b {
					p = i
				}
				continue
			}

			b := nums[i-1]
			a := nums[i+1]

			if c > b && c > a {
				p = i
			}
		}
	}

	return p
}
