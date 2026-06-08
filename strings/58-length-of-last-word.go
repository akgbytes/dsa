package strings

/*
https://leetcode.com/problems/length-of-last-word

Problem: Length of Last Word

----------------------------------------
Approach (Reverse Traversal):
- Traverse the string from right to left
- Ignore trailing spaces
- Mark the end of the last word
- Continue until a space is found
- Calculate the length using start and end indices

Time Complexity: O(n)
Space Complexity: O(1)
*/

func LengthOfLastWord(s string) int {
	endDetected := false
	validIdxEnd := -1
	validIdxStart := -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && endDetected {
			validIdxStart = i
			break
		}

		if s[i] != ' ' && !endDetected {
			endDetected = true
			validIdxEnd = i
		}
	}

	return validIdxEnd - validIdxStart
}
