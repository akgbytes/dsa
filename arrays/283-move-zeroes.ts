/*
https://leetcode.com/problems/move-zeroes
Problem: Move Zeroes

Approach:
- Use two pointers
- Pointer k tracks the index for next non-zero element
- Pointer i scans the array
- When a non-zero is found, place it at index k and move zero to current index

Time Complexity: O(n)
Space Complexity: O(1)
*/

function moveZeroes(nums: number[]): void {
  let k = 0;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== 0) {
      if (i !== k) {
        nums[k] = nums[i];
        nums[i] = 0;
      }
      k++;
    }
  }
}
