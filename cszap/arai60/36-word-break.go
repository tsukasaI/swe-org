package arai60

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	wordDictSet := make(map[string]struct{})
	for _, v := range wordDict {
		wordDictSet[v] = struct{}{}
	}
	dp[0] = true

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			_, ok := wordDictSet[s[i:j]]
			if dp[i] && ok {
				dp[j] = true
			}
		}
	}
	return dp[len(dp)-1]
}
