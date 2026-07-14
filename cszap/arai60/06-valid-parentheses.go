// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
    pair := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    stack := make([]byte, 0, len(s))

    for i := range s {
	if len(stack) == 0 {
		return false
	}
        if v, ok := pair[s[i]]; ok {
            // remove the element at the top of the stack if the v is same as it
            if v == stack[len(stack) - 1] {
                stack = stack[:len(stack) - 1]
            } else {
                return false
            }
        } else {
            stack = append(stack, s[i])
        }
    }

    return len(stack) == 0
}
