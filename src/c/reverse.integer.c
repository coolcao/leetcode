/**
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output: 321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only store integers within
the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this
problem, assume that your function returns 0 when the reversed integer
overflows.
 */

#include <math.h>
#include <stdio.h>
int reverse(int x) {
  int plus[10] = {2, 1, 4, 7, 4, 8, 3, 6, 4, 7};
  int minus[10] = {2, 1, 4, 7, 4, 8, 3, 6, 4, 8};
  int flag = x > 0 ? 1 : -1;

  int num = x;
  if (flag == -1) num = -x;
  int tmp[10] = {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1};

  int idx = 0;
  while (num > 0) {
    tmp[idx++] = num % 10;
    num = num / 10;
  }

  int sum = 0;

  // 这里需要判断反转之后是否溢出
  if (idx == 10) {
    if (flag == 1) {
      for (int i = 0; i < idx; i++) {
        int sub = tmp[i] - plus[i];
        if (sub > 0) {
          return 0;
        } else if (sub < 0) {
          break;
        }
      }
    } else {
      for (int i = 0; i < idx; i++) {
        int sub = tmp[i] - minus[i];
        if (sub > 0) {
          return 0;
        } else if (sub < 0) {
          break;
        }
      }
    }
  }

  for (int i = 0; i < idx; i++) {
    sum += tmp[i] * pow(10, idx - 1 - i);
  }

  return sum * flag;
}

int main(int argc, char const *argv[]) {
  int x = 4833;
  int rnum = reverse(x);
  printf("reverse number of %d is %d\n", x, rnum);
  return 0;
}
