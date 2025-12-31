/*
Problem: Find Second Largest Element in Array
Approach: Track second largest in single pass
Time: O(n), Space: O(1)
*/

function secondLargest(nums: number[]) {
  let first = -Infinity;
  let second = -Infinity;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] > second) {
      if (nums[i] > first) {
        second = first;
        first = nums[i];
      }
      // Handle duplicates
      else if (nums[i] == first) {
        continue;
      } else {
        second = nums[i];
      }
    }
  }

  return second;
}

console.log(secondLargest([5, 5, 1]));
