package linkedlist

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list

Problem: Remove Nth Node From End of List

----------------------------------------
Approach (Two Pointers + Sentinel):
- Use a dummy (sentinel) node
- Move fast pointer n steps ahead
- Move both slow and fast until fast reaches end
- Slow will be just before the target node
- Remove slow.Next

Time Complexity: O(n)
Space Complexity: O(1)
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{Next: head}
	slow, fast := sentinel, sentinel

	// move fast n steps ahead
	for range n {
		fast = fast.Next
	}

	// move both pointers
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// delete node
	slow.Next = slow.Next.Next

	return sentinel.Next
}
