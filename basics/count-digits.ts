/*
Problem: Count digits in a number
Approach: Repeated division by 10
Time: O(n), Space: O(1)
*/

function countDigits(num: number): number {
  if (num === 0) return 1;

  // Handle negatives
  num = Math.abs(num);
  let count = 0;

  while (num > 0) {
    num = Math.floor(num / 10);
    count++;
  }

  return count;
}

console.log(countDigits(-85));
