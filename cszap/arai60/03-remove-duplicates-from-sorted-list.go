package arai60

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next

	for slow != nil && fast != nil {
		if slow.Val == fast.Val {
			for fast != nil && fast.Val == slow.Val {
				fast = fast.Next
			}
		}
		slow.Next = fast

		if fast == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	return head
}


/**
 * 2026-03-14
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicate2(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    cur := head

    for cur != nil {
        next := cur.Next
        for next != nil && cur.Val == next.Val {
            next = next.Next
        }
        cur.Next = next
        cur = next
    }

    return head
}
