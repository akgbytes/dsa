package strings

import "strings"

/*
https://leetcode.com/problems/find-words-containing-character

Problem: Find Words Containing Character

----------------------------------------
Approach (String Search):
- Traverse all words
- Check if the target character exists in the current word
- If found, store the index

Time Complexity: O(n * m)
  - n = number of words
  - m = average word length

Space Complexity: O(k)
  - k = number of matching indices
*/

func FindWordsContaining(words []string, x byte) []int {
	result := make([]int, 0)

	for i, word := range words {
		if strings.ContainsRune(word, rune(x)) {
			result = append(result, i)
		}
	}

	return result
}
