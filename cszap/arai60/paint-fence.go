package arai60

func paintFence(n, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	same := []int{0}
	diff := []int{k}

	for i := 1; i < n; i++ {
		same = append(same, diff[i-1])
		diff = append(diff, (k-1)*(same[i-1]+diff[i-1]))
	}
	return same[n-1] + diff[n-1]
}
