package strings

/*
https://leetcode.com/problems/jewels-and-stones

Problem: Jewels and Stones

----------------------------------------
Approach (Set Lookup):
- Store all jewel characters in a set
- Traverse stones string
- If stone exists in set, increment count

Time Complexity: O(n + m)
Space Complexity: O(n)
*/

func NumJewelsInStones(jewels string, stones string) int {
	count := 0

	set := make(map[rune]struct{})

	// store jewels
	for _, ch := range jewels {
		set[ch] = struct{}{}
	}

	// count matching stones
	for _, ch := range stones {
		if _, ok := set[ch]; ok {
			count++
		}
	}

	return count
}
