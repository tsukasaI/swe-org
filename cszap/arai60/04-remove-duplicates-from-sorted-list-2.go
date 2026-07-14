// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{Val: -111, Next: head}

    cur := dummy
    for cur != nil {
        next := cur.Next
        shouldRemove := false
        for next != nil && next.Next != nil && next.Val == next.Next.Val {
            shouldRemove = true
            next = next.Next
        }
        if shouldRemove {
            if next != nil {
                cur.Next = next.Next
            } else {
                cur.Next = nil
            }
        } else {
            cur.Next = next
            cur = next
        }
    }

    return dummy.Next
}
