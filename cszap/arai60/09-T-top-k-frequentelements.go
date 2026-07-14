// https://leetcode.com/problems/top-k-frequent-elements/description/

func topKFrequent(nums []int, k int) []int {
    hm := make(map[int]int)

    for _, v := range nums {
        hm[v]++
    }

    numFreqArray := make([][]int, 0, len(hm))

    for i, f := range hm {
        numFreqArray = append(numFreqArray, []int{i, f})
    }

    sort.Slice(numFreqArray, func(i, j int) bool {
        return numFreqArray[i][1] > numFreqArray[j][1]
    })

    result := make([]int, k)

    for i := 0; i < k; i++ {
        result[i] = numFreqArray[i][0]
    }
    return result
}

type Freq struct{
    v int
    freq int
}

type MinFrepHeap []Freq

func (m MinFrepHeap) Len() int {
    return len(m)
}

func (m MinFrepHeap) Less(i, j int) bool {
    return m[i].freq < m[j].freq
}

func (m MinFrepHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m *MinFrepHeap) Push(x any) {
    f, _ := x.(Freq)
    *m = append(*m, f)
}

func (m *MinFrepHeap) Pop() any {
    old := *m
    v := old[len(old) - 1]
    old = old[:len(old) - 1]
    *m = old
    return v
}

func topKFrequentHeap(nums []int, k int) []int {
    hm := make(map[int]int)

    for _, v := range nums {
        hm[v]++
    }

    pq := make(MinFrepHeap, 0, len(hm))
    for i, f := range hm {
        heap.Push(&pq, Freq{i, f})
        if pq.Len() > k {
            heap.Pop(&pq)
        }
    }

    result := make([]int, k)
    for i := range k {
        result[i] = pq[i].v
    }
    return result
}
