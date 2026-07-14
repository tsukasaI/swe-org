// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
    hm := make(map[[26]int][]string)

    for _, v := range strs {
        slice := [26]int{}
        for _, b := range v {
            slice[b - 'a']++
        }
        hm[slice] = append(hm[slice], v)
    }
    
    result := make([][]string, 0, len(hm))
    for _, v := range hm {
        result = append(result, v)
    }

    return result
}
