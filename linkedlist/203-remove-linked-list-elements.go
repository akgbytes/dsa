package linkedlist

/*
https://leetcode.com/problems/remove-linked-list-elements

Problem: Remove Linked List Elements

----------------------------------------
Approach (Using Sentinel Node):
- Use a dummy (sentinel) node pointing to head
- Traverse the list using current pointer
- If current.Next matches target value, skip it
- Else move forward

Time Complexity: O(n)
Space Complexity: O(1)
*/

func RemoveElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{Next: head}
	current := sentinel

	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return sentinel.Next
}
