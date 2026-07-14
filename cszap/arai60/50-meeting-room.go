package arai60

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return x
}

func minMeetingRooms(meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	h := &minHeap{}
	count := 0
	for _, meeting := range meetings {
		for len(*h) > 0 && meeting[0] >= (*h)[0] {
			_ = heap.Pop(h)
		}
		heap.Push(h, meeting[1])
		count = max(count, len(*h))
	}
	return count
}
