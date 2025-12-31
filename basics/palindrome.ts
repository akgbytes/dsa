/*
Problem: Check if a string is palindrome
Approach: Two pointers from both ends
Time: O(n), Space: O(1)
*/

function isPalindrome(str: string): boolean {
  let i = 0;
  let j = str.length - 1;

  while (i < j) {
    if (str[i] !== str[j]) return false;
    i++;
    j--;
  }

  return true;
}
