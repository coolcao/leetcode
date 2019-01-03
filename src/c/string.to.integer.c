/**
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the
first non-whitespace character is found. Then, starting from this character,
takes an optional initial plus or minus sign followed by as many numerical
digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral
number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid
integral number, or if no such sequence exists because either str is empty or it
contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within
the 32-bit signed integer range: [−2^31,  2^31 − 1]. If the numerical value is
out of the range of representable values, INT_MAX (2^31 − 1) or INT_MIN (−2^31)
is returned.

Example 1:

Input: "42"
Output: 42

Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.

Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a
numerical digit.

Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be
performed.

Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed
integer. Thefore INT_MIN (−2^31) is returned.

 */


/**
 * 这个问题其实就是让自己写一个parseInt函数，将一个字符串解析成32位数字类型
 * 难度其实不难，只要了解了parseInt内部的实现逻辑流程即可，但其细节点非常多，想要一次性搞定，其实不容易。
 * parseInt的解析流程大致如下，只要按照流程解析就可以了：
 * 1. 首先去掉字符串中的所有前置空格，例如将 "   123dd" 去掉前置空格为 "123dd"
 * 2. 然后判断字符串的第一个字符，如果第一个字符是 正号(+)或负号(-)，将其去掉，并标记正负。如果不是正负号，则字符串不变。
 * 3. 截取字符串中连续的数字字串，例如 "123def" 将截取为 "123"
 * 4. 去掉前置的所有0， 例如将 "000123" 去掉前置0后为 "123"
 * 5. 判断剩下的这些数字是否在32位整数的范围内，如果超过范围，直接返回临界值
 * 6. 如果在范围内，则根据其ASCII码直接将其转换成32位整数即可。
 */


#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define min -pow(2, 31)
#define max pow(2, 31) - 1

void printStr(char *s) {
  int length = strlen(s);
  printf("|");
  for (int i = 0; i < length; i++) {
    printf("%c", s[i]);
  }
  printf("|\n");
}

char *removeLeadSpace(char *s) {
  int length = strlen(s);
  char *tmp = (char *)malloc(sizeof(char *) * length + 1);
  // 去除所有前置空格
  int space = 0;
  int idx = 0;
  for (int i = 0; i < length; i++) {
    char c = s[i];
    if (c == ' ') {
      if (space == 0) {
        continue;
      }
    } else {
      space = 1;
    }
    tmp[idx++] = c;
  }
  tmp[idx] = '\0';
  return tmp;
}

char *removeFlag(char *s) {
  int length = strlen(s);
  char *tmp = (char *)malloc(sizeof(char *) * length + 1);
  int idx = 0;
  int start = 0;
  if (s[0] == '+' || s[0] == '-') {
    start += 1;
  }

  for (int i=start;i<length;i++) {
    tmp[idx++] = s[i];
  }
  tmp[idx] = '\0';
  return tmp;
}

char *getNumbers(char *s) {
  int length = strlen(s);
  int isNumber = 1;
  char *tmp = (char *)malloc(sizeof(char*) * length + 1);
  int idx = 0;
  for (int i=0;i<length;i++) {
    char c = s[i];
    if (isNumber) {
      if (c>=48 && c <= 57) {
        tmp[idx++] = c ;
      } else {
        isNumber = 0;
      }
    } else {
      break;
    }
  }
  tmp[idx] = '\0';
  return tmp;
}

char *removeLeadZero(char *s) {
  int length = strlen(s);
  char *tmp = (char *)malloc(sizeof(char *) * length + 1);
  // 去除所有前置空格
  int zero = 0;
  int idx = 0;
  for (int i = 0; i < length; i++) {
    char c = s[i];
    if (c == '0') {
      if (zero == 0) {
        continue;
      }
    } else {
      zero = 1;
    }
    tmp[idx++] = c;
  }
  tmp[idx] = '\0';
  return tmp;
}



int myAtoi(char *s) {
  int length = strlen(s);
  if (length == 0) {
    return 0;
  }
  int plus[10] = {2, 1, 4, 7, 4, 8, 3, 6, 4, 7};
  int minus[10] = {2, 1, 4, 7, 4, 8, 3, 6, 4, 8};
  int flag = 1; // 标识符号，正数为1，负数为-1

  char* noneLeadZero = removeLeadSpace(s);

  // 判断首位是否是正负号，标记flag
  if (noneLeadZero[0] == '-') {
    flag = -1;
  }
  char* noneFlag = removeFlag(noneLeadZero);

  char* nums = getNumbers(noneFlag);

  char* noneZeroNums = removeLeadZero(nums);

  printStr(noneZeroNums);

  int lastLen = strlen(noneZeroNums);
  // 如果没有一个数字，则返回0
  if (lastLen == 0) return 0;
  // 如果长度已大于10，直接返回边界
  if (lastLen > 10) {
    return (flag>0?max:min);
  }
  // 长度为10， 逐一校验每位数字大小
  if (lastLen == 10) {
    if (flag > 0) {
      for (int i=0;i<lastLen;i++) {
        int sub = noneZeroNums[i]-48-plus[i];
        if (sub > 0) {
          return max;
        } else if (sub < 0) {
          break;
        }
      }
    } else {
      for (int i=0;i<lastLen;i++) {
        int sub = noneZeroNums[i]-48-minus[i];
        if (sub > 0) {
          return min;
        } else if (sub < 0) {
          break;
        }
      }
    }
  }

  int sum = 0;
  for (int i=0;i<lastLen;i++) {
    sum += (noneZeroNums[i]-48) * pow(10, lastLen-1-i);
  }

  return flag*sum;
}

int main(int argc, char const *argv[]) {
  char* s = "1095502006p8";
  int num = myAtoi(s);
  printf("%d\n", num);
  //
  // char* ss[] =
  // { "+-2",
  //   "-91283472332",
  //   "2147483648",
  //   "20000000000000000000",
  //   "   -42",
  //   "-   234",
  //   "0-1",
  //   "-5-",
  //   "  0000000000012345678",
  //   "-000000000000001",
  //   "  0000000000012345678",
  //   "010",
  //   "0-1" };
  
  // int n ;
  // for (int i=0;i<13;i++) {
  //   n = myAtoi(ss[i]);
  //   printf("%s = %d\n", ss[i], n);
  //   printf("===============\n");
  // }



  return 0;
}
