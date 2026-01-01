/*
https://leetcode.com/problems/palindrome-number/
Problem:  Palindrome Number

Approach 1: Reverse the entire number and compare

Idea:
- Negative numbers can never be palindrome
- Numbers ending with 0 (except 0 itself) are not palindrome
- Reverse the full number and compare with original

Time Complexity: O(log n)
Space Complexity: O(1)

*/

function isPalindrome(x: number): boolean {
  if (x < 0 || (x % 10 === 0 && x !== 0)) return false;

  let temp = x;
  let rev = 0;

  while (temp > 0) {
    rev = rev * 10 + (temp % 10);
    temp = Math.floor(temp / 10);
  }

  return rev === x;
}

/*
Approach 2: Reverse only half of the number (Optimized)

Idea:
- Negative numbers and numbers ending with 0 (except 0) are not palindrome
- Reverse digits only until reversedHalf >= remaining number
- Compare first half and second half
- For odd digit numbers, ignore the middle digit

Time Complexity: O(log n)
Space Complexity: O(1)

Why better?
- Avoids reversing the entire number
- Prevents overflow issues in other languages e.g. x = 2147483647  // INT_MAX -> rev = 7463847412 (limit cross)
*/

function isPalindrome2(x: number): boolean {
  if (x < 0 || (x % 10 === 0 && x !== 0)) return false;

  let reversedHalf = 0;
  while (x > reversedHalf) {
    reversedHalf = reversedHalf * 10 + (x % 10);
    x = Math.floor(x / 10);
  }

  return x === reversedHalf || x === Math.floor(reversedHalf / 10);
}
