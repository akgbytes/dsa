package strings

/*
https://leetcode.com/problems/jewels-and-stones

Problem: Jewels and Stones

----------------------------------------
Approach (HashMap Lookup):
- Store all jewel characters in a hashmap
- Traverse stones string
- If stone exists in hashmap, increment count

Time Complexity: O(n + m)
Space Complexity: O(n)
*/

func NumJewelsInStones(jewels string, stones string) int {
	count := 0

	hashmap := make(map[rune]bool)

	// store jewels
	for _, ch := range jewels {
		hashmap[ch] = true
	}

	// count matching stones
	for _, ch := range stones {
		if hashmap[ch] {
			count++
		}
	}

	return count
}
