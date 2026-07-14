package arai60

func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := make(map[string]struct{})
	for _, v := range wordList {
		set[v] = struct{}{}
	}

	if _, ok := set[endWord]; !ok {
		return 0
	}

	queue := []string{beginWord}
	count := 1
	for len(queue) > 0 {
		size := len(queue)
		for range size {
			target := queue[0]
			queue = queue[1:]
			// 各文字をa-zに変換してsetにあるか確認
			for i := range target {
				altered := []rune(target)
				for j := range 26 {
					altered[i] = rune("a"[0]) + rune(j)
					if string(altered) == endWord {
						return count + 1
					}
					if _, ok := set[string(altered)]; ok {
						queue = append(queue, string(altered))
						delete(set, string(altered))
					}
				}
			}
		}
		count++
	}
	return 0
}
