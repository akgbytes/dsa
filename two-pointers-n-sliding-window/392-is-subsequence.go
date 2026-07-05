package twopointersnslidingwindow

/*
https://leetcode.com/problems/is-subsequence

Problem: Is Subsequence

----------------------------------------
Approach (Two Pointers):
- Use two pointers:
  - i traverses string s
  - j traverses string t
- If characters match, move both pointers
- Otherwise, move only j
- If i reaches the end of s, then s is a subsequence of t

Time Complexity: O(n + m)
  - n = length of s
  - m = length of t

Space Complexity: O(1)
*/

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
