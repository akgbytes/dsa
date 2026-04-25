package linkedlist

/*
https://leetcode.com/problems/middle-of-the-linked-list

Problem: Middle of the Linked List

----------------------------------------
Naive Approach (2-pass):
- First pass: count total number of nodes
- Second pass: go to (length / 2)-th node

Time Complexity: O(n)
Space Complexity: O(1)
*/

func MiddleNodeNaive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// count length
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	// go to middle
	mid := length / 2
	current = head
	for _ = range mid {
		current = current.Next
	}

	return current
}

/*
----------------------------------------
Optimal Approach (Slow-Fast Pointer):
- Use two pointers (slow and fast)
- Slow moves 1 step, fast moves 2 steps
- When fast reaches end, slow will be at the middle

Time Complexity: O(n)
Space Complexity: O(1)
*/

func MiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
