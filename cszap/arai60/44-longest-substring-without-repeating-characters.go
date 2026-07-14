package arai60

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	set := make(map[byte]struct{})
	l := 0
	maxLen := 0
	for r := range s {
		for l < r {
			if _, ok := set[s[r]]; !ok {
				break
			}
			delete(set, s[l])
			l++
		}
		set[s[r]] = struct{}{}
		maxLen = max(maxLen, len(set))
	}
	return maxLen
}
