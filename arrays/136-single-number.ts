/*
https://leetcode.com/problems/single-number/
Problem: Single Number

Approach:
- Use XOR to cancel out duplicate elements
- Remaining value is the unique number

Time Complexity: O(n)
Space Complexity: O(1)
*/

function singleNumber(nums: number[]): number {
  let xor = 0;

  for (let i = 0; i < nums.length; i++) {
    xor ^= nums[i];
  }

  return xor;
}
