package arai60

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(N)
// Space: O(N)
func hasCycle(head *ListNode) bool {
	hm := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := hm[head]; ok {
			return true
		}
		hm[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycleWithConstantMemory(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func followUpHasCycleWithConsistencyAndReturnTheStartOfCycleNode(head *ListNode) *ListNode {
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
	if fast == slow {
		return nil
	}
	
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
