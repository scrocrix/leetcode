package easy

import "sort"

func twoSum(nums []int, target int) []int {
	var o []int
	o = append(o, nums...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	right := len(nums) - 1
	left := 0

	var i []int

	for {
		r := nums[left] + nums[right]
		if r > target {
			right = right - 1
		} else if r < target {
			left = left + 1
		} else if r == target {
			for index, num := range o {
				if num == nums[left] || num == nums[right] {
					i = append(i, index)
				}
			}
			break
		}
	}

	return i
}
