// https://leetcode.com/problems/reverse-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseListIter(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var l *ListNode
    r := head

    for r != nil {
        next := r.Next

        r.Next = l
        l = r
        r = next
    }
    return l
}


func reverseListRec(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    var reverseLinkedListRecursive func(*ListNode, *ListNode) *ListNode
    reverseLinkedListRecursive  = func(l, r *ListNode) *ListNode {
        if r == nil {
            return l
        }
        next := r.Next
        r.Next = l
        return reverseLinkedListRecursive(r, next)
    }
    var l *ListNode
    return reverseLinkedListRecursive(l, head)
}
