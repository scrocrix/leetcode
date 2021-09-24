package medium

import (
	"math/big"
	"sort"
	"strconv"
)

// O(nˆ2)
func RunningMedian(a []int32) []float64 {
	var items []int32
	var medians []float64

	for _, num := range a {
		items = append(items, num)

		sort.Slice(items, func(i, j int) bool {
			return items[i] < items[j]
		})

		if len(items) > 0 && len(items)%2 != 0 {
			m := items[len(items)/2]
			n, _ := strconv.ParseFloat(big.NewFloat(float64(m)).Text('f', 1), 1)
			medians = append(medians, n)
		} else if len(items) > 0 && len(items)%2 == 0 {
			m1 := items[(len(items)/2)-1]
			m2 := items[len(items)/2]
			m := (float64(m1) + float64(m2)) / float64(2)
			medians = append(medians, m)
		}
	}

	return medians
}
