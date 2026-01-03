/*
https://leetcode.com/problems/reverse-string
Problem: Reverse String

Approach:
- Use two pointers from both ends
- Swap characters in-place until pointers meet

Time Complexity: O(n)
Space Complexity: O(1)
*/

function reverseString(s: string[]): void {
  let i = 0;
  let j = s.length - 1;

  while (i < j) {
    [s[i], s[j]] = [s[j], s[i]];
    i++;
    j--;
  }
}
