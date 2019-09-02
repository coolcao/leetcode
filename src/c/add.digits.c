/**
Given a non-negative integer num, repeatedly add all its digits until the result
has only one digit.

Example:

Input: 38
Output: 2
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
             Since 2 has only one digit, return it.
Follow up:
Could you do it without any loop/recursion in O(1) runtime?
 */

#include <stdio.h>
int addDigits(int num) {
  if (num < 10) {
    return num;
  }
  int n = num % 9;
  return n == 0 ? 9 : n;
}

int main(int argc, char const *argv[]) {
  int num = 876;
  printf("%d\n", addDigits(num));
  return 0;
}
