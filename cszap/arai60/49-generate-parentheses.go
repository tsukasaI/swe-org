package arai60

import "strings"

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	current := make([]string, 0)
	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(current, ""))
			return
		}

		if open < n {
			current = append(current, "(")
			backtrack(open+1, close)
			current = current[:len(current)-1]
		}
		if close < open {
			current = append(current, ")")
			backtrack(open, close+1)
			current = current[:len(current)-1]

		}
	}
	backtrack(0, 0)
	return result
}
