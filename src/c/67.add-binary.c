/*
 * @lc app=leetcode id=67 lang=c
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (37.63%)
 * Total Accepted:    270.7K
 * Total Submissions: 718.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// void printStr(char* s) {
//   int len = strlen(s);
//   for (int i = 0; i < len; i++) {
//     printf("%c", s[i]);
//   }
//   printf("\n");
// }
char* addBinary(char* a, char* b) {
  int alen = strlen(a);
  int blen = strlen(b);
  int maxlen = alen > blen ? alen : blen;

  char array_a[alen+1];
  char array_b[blen+1];
  strcpy(array_a, a);
  strcpy(array_b, b);

  int flag = 0;
  char* result = NULL;
  result = (char*)malloc(sizeof(char*) * (maxlen + 1));

  int aidx = alen - 1;
  int bidx = blen - 1;
  while (aidx >= 0 && bidx >= 0) {
    int anum = array_a[aidx] - 48;
    int bnum = array_b[bidx] - 48;

    int sum = anum + bnum + flag;
    flag = sum / 2;
    sum = sum % 2;

    array_a[aidx] = sum + 48;
    array_b[bidx] = sum + 48;

    aidx -= 1;
    bidx -= 1;
  }

  int idx = 0;

  if (alen > blen) {
    while (aidx >= 0) {
      int anum = array_a[aidx] - 48;
      int sum = anum + flag;
      flag = sum / 2;
      sum = sum % 2;
      array_a[aidx] = sum + 48;

      aidx -= 1;
    }

    if (flag) {
      result[idx++] = flag + 48;
    }

    for (int i = 0; i < alen; i++) {
      result[idx++] = array_a[i];
    }
  } else {
    while (bidx >= 0) {
      int bnum = array_b[bidx] - 48;
      int sum = bnum + flag;
      flag = sum / 2;
      sum = sum % 2;
      array_b[bidx] = sum + 48;

      bidx -= 1;
    }
    if (flag) {
      result[idx++] = flag + 48;
    }
    for (int i = 0; i < blen; i++) {
      result[idx++] = array_b[i];
    }
  }
  result[idx++] = '\0';
  return result;
}

// int main(int argc, char const* argv[]) {
//   char* a = "110011011001";
//   char* b = "100110";
//   char* c = addBinary(a, b);
//   printStr(c);
//   return 0;
// }
