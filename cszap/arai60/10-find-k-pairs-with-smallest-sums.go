// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
type Node struct {
    pair []int
    sum, i, j int
}

type MinHeap []Node

func (m MinHeap) Len() int {
    return len(m)
}

func (m MinHeap) Less(i, j int) bool {
    return m[i].sum < m[j].sum
}

func (m MinHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
    n := x.(Node)
    *m = append(*m, n)
}

func (m *MinHeap) Pop() any {
    old := *m
    length := len(old) - 1
    v := old[length]
    *m = old[:length]
    return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    pq := make(MinHeap, 0)

    result := make([][]int, 0, len(pq))
    for i := range min(len(nums1), k) {
        heap.Push(&pq, Node{pair: []int{nums1[i], nums2[0]}, sum: nums1[i] + nums2[0], i: i, j: 0})
    }
    for len(result) < k {
        popped := heap.Pop(&pq).(Node)
            result = append(result, popped.pair)
        if popped.j + 1 < len(nums2) {
            popedI := popped.i
            popedJ := popped.j
            heap.Push(&pq, Node{pair: []int{nums1[popedI], nums2[popedJ + 1]}, sum: nums1[popedI] + nums2[popedJ + 1], i: popedI, j: popedJ + 1})
        }
    }
    return result
}
