package linkedlist

/*
https://leetcode.com/problems/intersection-of-two-linked-lists

Problem: Intersection of Two Linked Lists

----------------------------------------
Naive Approach (Using HashMap):
- Traverse first linked list and store all nodes in a map
- Traverse second list and check if any node exists in map
- First matching node is the intersection

Time Complexity: O(m + n)
Space Complexity: O(m)
*/

func GetIntersectionNodeNaive(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})

	for curr := headA; curr != nil; curr = curr.Next {
		visited[curr] = struct{}{}
	}

	for curr := headB; curr != nil; curr = curr.Next {
		if _, exists := visited[curr]; exists {
			return curr
		}
	}

	return nil
}

/*
----------------------------------------
Optimal Approach (Two Pointers / Pointer Switching):
- Use two pointers a and b starting at headA and headB
- Traverse both lists
- When a reaches end, redirect it to headB
- When b reaches end, redirect it to headA
- Eventually both pointers will meet at intersection or nil

Time Complexity: O(m + n)
Space Complexity: O(1)
*/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
