// https://leetcode.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
    counts := [26]int{}

    // count all characters
    for _, v := range s {
        counts[v - 'a']++
    }

    for i, v := range s {
        if counts[v - 'a'] == 1 {
            return i
        }
    }
    return -1
}
