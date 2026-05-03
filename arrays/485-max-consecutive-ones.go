package arrays

/*
https://leetcode.com/problems/max-consecutive-ones

Problem: Max Consecutive Ones

Approach:
- Traverse array once
- Count consecutive 1s
- Reset count on 0
- Track maximum streak

Time Complexity: O(n)
Space Complexity: O(1)
*/

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	c := 0

	for _, v := range nums {
		if v != 1 {
			res = max(res, c)
			c = 0
		} else {
			c++
		}

	}

	return max(res, c)
}
