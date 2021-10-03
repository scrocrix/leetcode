package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	} else {
		i := (len(nums) - 1) / 2
		j := i + 1
		return float64((float64(nums[i]) + float64(nums[j])) / 2)
	}

	return -1
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}
