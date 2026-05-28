package linkedlist

/*
https://leetcode.com/problems/odd-even-linked-list

Problem: Odd Even Linked List

----------------------------------------
Approach (Separate Odd and Even Chains):
- Traverse the linked list
- Maintain separate odd and even chains
- Store head of even list
- Connect odd list with even list at the end

Time Complexity: O(n)
Space Complexity: O(1)
*/

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head

	even := &ListNode{}
	evenHead := even

	for odd.Next != nil {
		// attach even node
		even.Next = odd.Next
		even = even.Next

		// skip even node in odd chain
		if odd.Next.Next != nil {
			odd.Next = odd.Next.Next
			odd = odd.Next
		} else {
			break
		}
	}

	// terminate even chain
	even.Next = nil

	// attach even list after odd list
	odd.Next = evenHead.Next

	return head
}
