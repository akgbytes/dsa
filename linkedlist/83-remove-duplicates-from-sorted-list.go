package linkedlist

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list

Problem: Remove Duplicates from Sorted List

----------------------------------------
Approach (Single Pointer Traversal):
- Traverse the linked list using a pointer
- Compare current node with next node
- If both values are equal, skip the next node
- Otherwise move forward

Time Complexity: O(n)
Space Complexity: O(1)
*/

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
