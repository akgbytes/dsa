package linkedlist

/*
https://leetcode.com/problems/merge-two-sorted-lists

Problem: Merge Two Sorted Lists

----------------------------------------
Approach (Dummy Node + Two Pointers):
- Use a dummy node to build merged list
- Compare nodes from both lists
- Attach smaller node to merged list
- Append remaining nodes at the end

Time Complexity: O(n + m)
Space Complexity: O(1)
*/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ll := &ListNode{}
	llHead := ll

	// merge while both lists exist
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			ll.Next = list2
			list2 = list2.Next
		} else {
			ll.Next = list1
			list1 = list1.Next
		}

		ll = ll.Next
	}

	// attach remaining list1 nodes
	for list1 != nil {
		ll.Next = list1
		ll = ll.Next
		list1 = list1.Next
	}

	// attach remaining list2 nodes
	for list2 != nil {
		ll.Next = list2
		ll = ll.Next
		list2 = list2.Next
	}

	return llHead.Next
}
