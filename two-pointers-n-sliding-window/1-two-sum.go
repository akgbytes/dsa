package twopointersnslidingwindow

/*
https://leetcode.com/problems/two-sum

Problem: Two Sum

----------------------------------------
Approach (HashMap):
- Traverse the array once
- For each element, calculate its complement (target - current element)
- If the complement already exists in the hashmap, return both indices
- Otherwise, store the current element and its index in the hashmap

Time Complexity: O(n)
Space Complexity: O(n)
*/

func TwoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, exists := hashmap[complement]; exists {
			return []int{idx, i}
		}

		hashmap[num] = i
	}

	return nil
}
