/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array
Problem: Remove Duplicates from Sorted Array

Approach:
- Use two pointers
- Pointer i tracks position of last unique element
- Pointer j scans the array
- When nums[j] != nums[i], move i forward and overwrite nums[i]

Time Complexity: O(n)
Space Complexity: O(1)
*/

function removeDuplicates(nums: number[]): number {
  let i = 0;
  for (let j = 0; j < nums.length; j++) {
    if (nums[i] !== nums[j]) {
      i++;
      nums[i] = nums[j];
    }
  }
  return i + 1;
}

console.log(removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]));
