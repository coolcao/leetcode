/*
 * @lc app=leetcode id=66 lang=c
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.55%)
 * Total Accepted:    345.2K
 * Total Submissions: 851.1K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 *
 */
/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */

#include <stdio.h>
#include <stdlib.h>

int* plusOne(int* digits, int digitsSize, int* returnSize) {
  int* result;
  int flag = 1;
  int idx = 0;

  for (int i = digitsSize - 1; i >= 0; i--) {
    int sum = flag + digits[i];
    flag = sum / 10;
    sum = sum % 10;
    digits[i] = sum;
  }
  if (flag != 0) {
    *returnSize = digitsSize + 1;
    result = malloc(sizeof(int) * (digitsSize + 1));
    result[idx++] = flag;
  } else {
    *returnSize = digitsSize;
    result = malloc(sizeof(int) * digitsSize);
  }
  for (int i = 0; i < digitsSize; i++) {
    result[idx++] = digits[i];
  }

  return result;
}

/**
 * 这个方法是之前的思路，将整数倒序放到数组进行计算。倒序的原因是计算时容易处理。
 * 但这里只是加1，其实可以不用倒序 
 */
int* plusOne2(int* digits, int digitsSize, int* returnSize) {
  int* result;
  result = malloc(sizeof(int) * (digitsSize + 1));
  int flag = 1;
  for (int i = 0; i < digitsSize + 1; i++) {
    int sum = 0;
    if (i == digitsSize) {
      sum = 0 + flag;
    } else {
      sum = digits[digitsSize - 1 - i] + flag;
    }
    flag = sum / 10;
    if (sum > 9) {
      sum %= 10;
    }
    result[i] = sum;
    // printf("%d\n", result[i]);
  }

  *returnSize = digitsSize + 1;
  if (result[digitsSize] == 0) {
    *returnSize = digitsSize;
  }
  // printf("returnSize: %d\n", *returnSize);
  for (int i = 0; i < *returnSize / 2; i++) {
    int tmp = result[i];
    result[i] = result[*returnSize - 1 - i];
    result[*returnSize - 1 - i] = tmp;
  }

  return result;
}

// int main(int argc, char const* argv[]) {
// #define LEN 5
//   int a[LEN] = {2, 6, 8, 9, 5};
//   int returnSize;
//   int* result = plusOne(a, LEN, &returnSize);
//   for (int i = 0; i < returnSize; i++) {
//     printf("%d", result[i]);
//   }
//   printf("\n");
//   return 0;
// }
