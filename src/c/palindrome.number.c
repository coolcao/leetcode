/**
Determine whether an integer is a palindrome. An integer is a palindrome when it
reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes
121-. Therefore it is not a palindrome. Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?
 */

/**
 * 判断一个整数是否是回文整数
 * 要求不要转换成字符串进行解决
 * 不用字符串的实质其实就是，避免直接使用语言现成的反转函数，自己写回文比较函数
 * 我们可以使用数组解决。
 * 当然也可以计算反转后的值是否和原值相等，但这样可能会出现数值溢出（从逻辑上讲，对于溢出的值，肯定不是回文数，其实这也可以正确执行）
 */

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
bool isPalindrome(int x) {
  if (x >= 0 && x <= 9) return true;
  if (x < 0 || (x && x % 10 == 0)) return false;
  int *nums = (int *)malloc(sizeof(int) * 10);

  int idx = 0;
  int num = x;
  while (num > 0) {
    nums[idx++] = num % 10;
    num /= 10;
  }
  for (int i = 0; i < idx / 2; i++) {
    if (nums[i] != nums[idx - 1 - i]) {
      return false;
    }
  }
  return true;
}

int main(int argc, char const *argv[]) {
  int x = -121;
  bool re = isPalindrome(x);
  printf("%d is palindrome? %d\n", x, re);
  return 0;
}
