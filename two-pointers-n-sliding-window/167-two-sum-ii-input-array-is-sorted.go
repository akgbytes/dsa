package twopointersnslidingwindow

/*
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

Problem: Two Sum II - Input Array Is Sorted

----------------------------------------
Approach (Two Pointers):
- Initialize two pointers:
  - left at the beginning
  - right at the end
- Calculate the sum of both elements
- If sum equals target, return their 1-based indices
- If sum is greater than target, move right pointer left
- If sum is smaller than target, move left pointer right

Time Complexity: O(n)
Space Complexity: O(1)
*/

func TwoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
