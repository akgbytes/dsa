package linkedlist

/*
https://leetcode.com/problems/palindrome-linked-list

Problem: Palindrome Linked List

----------------------------------------
Approach (Middle + Reverse + Compare):
- Find the middle of the linked list
- Reverse the second half of the list
- Compare first half and reversed second half
- If all values match → palindrome

Time Complexity: O(n)
Space Complexity: O(1)
*/

func IsPalindrome(head *ListNode) bool {
	// find middle
	mid := MiddleNode(head)

	// reverse second half
	mid = ReverseList(mid)

	// compare both halves
	current := head
	for mid != nil {
		if current.Val != mid.Val {
			return false
		}
		current = current.Next
		mid = mid.Next
	}

	return true
}
