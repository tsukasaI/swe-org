package arai60

func kthGrammar(n int, k int) int {
	if n == 1 && k == 1 {
		return 0
	}

	if k%2 == 0 {
		return 1 - kthGrammar(n-1, (k+1)/2)
	} else {
		return kthGrammar(n-1, (k+1)/2)
	}
}
