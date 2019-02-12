/*
 * @lc app=leetcode id=28 lang=c
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (31.06%)
 * Total Accepted:    373.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */
#include <stdio.h>
#include <string.h>
int strStr(char* haystack, char* needle) {
  int hlen = strlen(haystack);
  int nlen = strlen(needle);
  // printf("hlen=%d\n", hlen);
  // printf("nlen=%d\n", nlen);
  if (hlen == 0 && nlen != 0) return -1;
  if (nlen == 0) return 0;
  int idx = -1;

  for (int i=0;i<hlen + 1 - nlen;i++) {
    if (haystack[i] == needle[0]) {
      for (int j=0;j<nlen;j++) {
        if (haystack[i+j] == needle[j]) {
          if (j == nlen - 1) {
            idx = i + j;
            goto hear;
          }
          continue;
        } else {
          idx = -1;
          break;
        }
      }
    } else {
      continue;
    }
  }
  hear:
  // printf("idx=%d\n", idx);
  // printf("nlen=%d\n", nlen);
  return idx == -1 ? -1 : idx + 1 - nlen;
}
// int main(int argc, char const* argv[]) {
//   char* haystack = "a";
//   char* needle = "a";
//   int idx = strStr(haystack, needle);
//   printf("%d\n", idx);
//   return 0;
// }
