package linkedlist

/*
https://leetcode.com/problems/reverse-linked-list

Problem: Reverse Linked List

----------------------------------------
Naive Approach (Using Extra Space):
- Traverse the linked list and store values in an array
- Create a new linked list by inserting elements from array in reverse order

Time Complexity: O(n)
Space Complexity: O(n)
*/

func ReverseListNaive(head *ListNode) *ListNode {
	// convert to array
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	var ll *ListNode
	for _, v := range arr {
		n := &ListNode{Val: v}
		n.Next = ll
		ll = n
	}

	return ll
}

/*
----------------------------------------
Optimal Approach (In-place Reversal):
- Use three pointers: prev, curr, next
- Reverse the direction of each node's pointer
- Move through the list once

Time Complexity: O(n)
Space Complexity: O(1)
*/

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // store next node
		curr.Next = prev  // reverse link
		prev = curr       // move prev forward
		curr = next       // move curr forward
	}

	return prev
}
