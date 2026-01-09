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

function findMaxConsecutiveOnes(nums: number[]): number {
  let max = 0;
  let count = 0;

  for (const num of nums) {
    if (num === 0) {
      max = Math.max(max, count);
      count = 0;
    } else {
      count++;
    }
  }

  return Math.max(max, count);
}
