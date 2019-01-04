/**
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and
M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added
together. Twelve is written as, XII, which is simply X + II. The number twenty
seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written
as IV. Because the one is before the five we subtract it making four. The same
principle applies to the number nine, which is written as IX. There are six
instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be
within the range from 1 to 3999.

Example 1:
Input: 3
Output: "III"

Example 2:
Input: 4
Output: "IV"

Example 3:
Input: 9
Output: "IX"

Example 4:
Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

Example 5:
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

====
Medium
====
 */

/**
 * 输入一个整数，范围在1-3999，输出其罗马计数法表示的字符串形式
 * 难度：中等
 * 直接把各个特殊数字的表示一一列出来，然后计算每个数字的个数，再拼起来即可。不知道为啥难度标记为中等。
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* intToRoman(int num) {
  int nums_of_eachs[13] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
  int nums_of_each = 0;
  int bases[13] = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
  char* symbols[13] = {"M",  "CM", "D",  "CD", "C",  "XC", "L",
                       "XL", "X",  "IX", "V",  "IV", "I"};

  int tmp = num;
  for (int i = 0; i < 13; i++) {
    int base = bases[i];
    int each = tmp / base;
    nums_of_eachs[i] = each;
    nums_of_each += each;
    tmp = tmp % base;
  }

  int idx = 0;
  char* str = (char*)malloc(sizeof(char*) * nums_of_each * 2 + 1);
  for (int i = 0; i < 13; i++) {
    int each = nums_of_eachs[i];
    if (each != 0) {
      char* ss = symbols[i];
      if (each == 1) {
        int length = strlen(ss);
        for (int j = 0; j < length; j++) {
          str[idx++] = ss[j];
        }
      } else {
        for (int j = 0; j < each; j++) {
          str[idx++] = ss[0];
        }
      }
    }
  }
  str[idx] = '\0';

  return str;
}

int main(int argc, char const* argv[]) {
  int num = 3999;
  char* s = intToRoman(num);
  printf("%s\n", s);
  return 0;
}
