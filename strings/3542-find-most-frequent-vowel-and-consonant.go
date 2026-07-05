package strings

/*
https://leetcode.com/problems/find-most-frequent-vowel-and-consonant

Problem: Maximum Frequency of Vowels and Consonants

----------------------------------------
Approach (HashMap):
- Count the frequency of each character using a hashmap
- Traverse the hashmap
- Track the maximum frequency among vowels
- Track the maximum frequency among consonants
- Return the sum of both maximum frequencies

Time Complexity: O(n)
Space Complexity: O(1)
  - At most 26 lowercase English letters are stored
*/

func MaxFreqSum(s string) int {
	hashmap := make(map[rune]int)

	// count character frequencies
	for _, ch := range s {
		hashmap[ch]++
	}

	maxFreqVowel := 0
	maxFreqConst := 0

	for ch, freq := range hashmap {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			maxFreqVowel = max(maxFreqVowel, freq)
		default:
			maxFreqConst = max(maxFreqConst, freq)
		}
	}

	return maxFreqVowel + maxFreqConst
}
