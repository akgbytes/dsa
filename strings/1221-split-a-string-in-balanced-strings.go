package strings

/*
https://leetcode.com/problems/split-a-string-in-balanced-strings

Problem: Split a String in Balanced Strings

----------------------------------------
Approach (Counting):
- Traverse the string once
- Count occurrences of 'L' and 'R'
- Whenever both counts become equal, one balanced substring is found
- Reset the counters and continue

Time Complexity: O(n)
Space Complexity: O(1)
*/

func BalancedStringSplit(s string) int {
	leftCount := 0
	rightCount := 0

	balancedCount := 0

	for _, ch := range s {
		if ch == 'L' {
			leftCount++
		} else {
			rightCount++
		}

		if leftCount == rightCount {
			balancedCount++
			leftCount = 0
			rightCount = 0
		}
	}

	return balancedCount
}
