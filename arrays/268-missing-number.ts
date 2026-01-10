/*
Approach 1: Math Sum
- Expected sum = n * (n + 1) / 2
- Actual sum = sum of array
- Difference is missing number

Time: O(n)
Space: O(1)
*/

function missingNumber(nums: number[]): number {
  const n = nums.length;
  const expected = (n * (n + 1)) / 2;
  const actual = nums.reduce((a, b) => a + b, 0);

  return expected - actual;
}

/*
Approach 2: XOR
- XOR all indices (0 to n) and all array values
- Duplicate values cancel out
- Remaining value is the missing number

Time Complexity: O(n)
Space Complexity: O(1)
*/

function missingNumber2(nums: number[]): number {
  let xor = nums.length;

  for (let i = 0; i < nums.length; i++) {
    xor ^= i ^ nums[i];
  }

  return xor;
}
