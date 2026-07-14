// https://leetcode.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
    num1Set := make(map[int]struct{})
    for _, v := range nums1 {
        num1Set[v] = struct{}{}
    }

    result := make([]int, 0, len(num1Set))
    for _, v := range nums2 {
        if _, ok := num1Set[v]; ok {
            result = append(result, v)
            delete(num1Set, v)
        }
    }
    return result
}
