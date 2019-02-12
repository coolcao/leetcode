/*
 * @lc app=leetcode id=58 lang=c
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.14%)
 * Total Accepted:    241.9K
 * Total Submissions: 752.8K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 * 
 * If the last word does not exist, return 0.
 * 
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 * 
 * Example:
 * 
 * Input: "Hello World"
 * Output: 5
 * 
 * 
 */
#include <stdio.h>
#include <string.h>
int lengthOfLastWord(char* s) {
  int len = strlen(s);
  if (len == 0) return 0;
  int lastSpaceStart = 0;
  int lastSpaceEnd = 0;
  int begin = 0;
  int read = 0;
  for (int i=0;i<len;i++) {
    if ((s[i]>=65 && s[i]<=90) || (s[i] >= 97 && s[i] <= 122)) {
      read = 1;
      if (begin == 0) {
        begin = 1;
        lastSpaceStart = i;
        if (i == len -1 ) lastSpaceEnd = i;
      } else {
        lastSpaceEnd = i;
        continue;
      }
    } else {
      if (begin) {
        lastSpaceEnd = i - 1;
        begin = 0;
      }
    }
  }
  if (read == 0) return 0;
  return lastSpaceEnd - lastSpaceStart + 1;
}

// int main(int argc, char const *argv[])
// {
//   char* s = "";
//   int length = lengthOfLastWord(s);
//   printf("%d\n", length);
//   return 0;
// }
