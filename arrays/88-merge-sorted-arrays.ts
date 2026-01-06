/*
https://leetcode.com/problems/merge-sorted-array
Problem: Merge Sorted Array

Approach:
- Start merging from the end to avoid overwriting nums1 values
- i points to last valid element in nums1
- j points to last element in nums2
- k points to last index in nums1

Time Complexity: O(m + n)
Space Complexity: O(1)
*/

function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  let i = m - 1;
  let j = n - 1;
  let k = nums1.length - 1;

  // Merge from the end
  while (i >= 0 && j >= 0) {
    if (nums1[i] >= nums2[j]) {
      nums1[k] = nums1[i];
      i--;
    } else {
      nums1[k] = nums2[j];
      j--;
    }
    k--;
  }

  // Copy remaining elements from nums2 (if any)
  while (j >= 0) {
    nums1[k] = nums2[j];
    j--;
    k--;
  }
}
