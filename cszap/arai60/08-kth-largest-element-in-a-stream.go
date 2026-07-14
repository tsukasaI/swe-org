// https://leetcode.com/problems/kth-largest-element-in-a-stream/

// 1st



type KthLargest struct {
    k int
    nums []int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{
        k: k,
        nums: nums,
    }
}


func (this *KthLargest) Add(val int) int {
    this.nums = append(this.nums, val)
    sort.Ints(this.nums)
    return this.nums[len(this.nums)-this.k]
}

----

// 2nd

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j] 
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(v any) {
    *h = append(*h, v.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    v := old[len(old) - 1]
    *h = old[:len(old) - 1]
    return v
}

type KthLargest struct {
    k int
    nums MinHeap
}


func Constructor(k int, nums []int) KthLargest {
    h := MinHeap(nums)
    heap.Init(&h)
    for len(h) > k {
        heap.Pop(&h)
    }
    return KthLargest{
        k: k,
        nums: h,
    }
}


func (this *KthLargest) Add(val int) int {
    heap.Push(&(this.nums), val)
    if len(this.nums) > this.k {
        heap.Pop(&this.nums)
    }
    v := this.nums[0]
    return v
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */


