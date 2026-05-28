package strings

/*
https://leetcode.com/problems/rotate-string

Problem: Rotate String

----------------------------------------
Approach (Concatenation + Substring Check):
- If lengths differ, rotation is impossible
- Concatenate string with itself
- Any valid rotation of s must exist as a substring in s+s
- Check all substrings of length len(goal)

Time Complexity: O(n²)
Space Complexity: O(n)
*/

func RotateString(s string, goal string) bool {
	l := len(goal)

	// rotation impossible if lengths differ
	if l != len(s) {
		return false
	}

	// duplicate string
	s = s + s

	// check every substring of length l
	for i := range s {
		if i+l > len(s) {
			break
		}

		if s[i:i+l] == goal {
			return true
		}
	}

	return false
}
