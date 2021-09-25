package medium

import (
	"container/heap"
	"fmt"
)

// IntHeap
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap
type MaxHeap struct {
	IntHeap
}

func (h MaxHeap) Less(i, j int) bool {
	return h.IntHeap[i] > h.IntHeap[j]
}

func RunningMedian(a []int32) []float64 {
	min := &IntHeap{}
	heap.Init(min)

	max := &MaxHeap{}
	heap.Init(max)

	var medians []float64

	for _, number := range a {
		if max.Len() == 0 {
			heap.Push(max, int(number))
			medians = append(medians, float64(number))
			continue
		}

		heap.Push(min, int(number))

		if min.Len()-max.Len() >= 2 {
			heap.Push(max, heap.Pop(min).(int))
		}

		mn := heap.Pop(min).(int)
		mmn := heap.Pop(max).(int)
		heap.Push(min, mn)
		heap.Push(max, mmn)

		medians = append(medians, float64((mn+mmn)/2))
	}

	fmt.Println("Min:")
	fmt.Println(min)
	fmt.Println("Max:")
	fmt.Println(max)

	return medians
}
