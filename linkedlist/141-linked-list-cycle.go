package linkedlist

/*
https://leetcode.com/problems/linked-list-cycle

Problem: Linked List Cycle

----------------------------------------
Naive Approach (Using HashSet):
- Traverse the linked list
- Store each visited node in a map
- If a node is visited again, cycle exists

Time Complexity: O(n)
Space Complexity: O(n)
*/

func HasCycleNaive(head *ListNode) bool {
	set := make(map[*ListNode]struct{})

	for head != nil {
		if _, exists := set[head]; exists {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}

	return false
}

/*
----------------------------------------
Optimal Approach (Floyd’s Cycle Detection / Slow-Fast Pointer):
- Use two pointers: slow and fast
- Slow moves 1 step, fast moves 2 steps
- If there is a cycle, they will meet
- If fast reaches nil, no cycle exists

Time Complexity: O(n)
Space Complexity: O(1)
*/

func HasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
