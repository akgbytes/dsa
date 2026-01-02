/*
https://leetcode.com/problems/remove-element
Problem: Remove Element from Sorted Array

Approach:
- Use two pointers
- Pointer j scans the array
- Pointer i keeps track of position to place elements not equal to val
- Modify array in-place

Time Complexity: O(n)
Space Complexity: O(1)
*/

function removeElement(nums: number[], val: number): number {
  let i = 0;
  for (let j = 0; j < nums.length; j++) {
    if (nums[j] !== val) {
      nums[i] = nums[j];
      i++;
    }
  }
  return i;
}
console.log(removeElement([0, 1, 2, 2, 3, 0, 4, 2], 3));
