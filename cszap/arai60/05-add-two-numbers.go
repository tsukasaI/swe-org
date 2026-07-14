/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    dummy := &ListNode{}
    current := dummy
    for l1 != nil || l2 != nil || carry == 1 {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        sum := v1 + v2 + carry
        current.Next = &ListNode{Val: sum%10}
        current = current.Next
        

        if sum > 9 {
            carry = 1
        } else {
            carry = 0
        }
    }
    return dummy.Next
}
