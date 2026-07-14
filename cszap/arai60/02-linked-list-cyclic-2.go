package arai60

func detectCycle(head *ListNode) *ListNode {
	passed := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := passed[head]; ok {
			return head
		}
		passed[head] = struct{}{}
		head = head.Next
	}
	return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycleTwoPointer(head *ListNode) *ListNode {
	fast, slow := head, head
	found := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	for slow != head {
		head = head.Next
		slow = slow.Next
	}
	return head
}



/**
 * 2026/03/14
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            break
        }
    }

    if slow != fast {
        return nil
    }

    for head != slow {
        head = head.Next
        slow = slow.Next
    }
    return head
}
