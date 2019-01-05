/**
Write a function to find the longest common prefix string amongst an array of
strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.

====
Easy
====
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* longestCommonPrefix(char** strs, int strsSize) {
  int length = strlen(strs[0]);
  char* tmp = (char*)malloc(sizeof(char*) * (length + 1));
  char c;
  int idx = 0;
  for (int i = 0; i < length; i++) {
    c = strs[0][i];
    int all = 1;
    for (int j = 1; j < strsSize; j++) {
      char* str = strs[j];
      if (str[i] != c) {
        all = 0;
        break;
      }
    }
    if (all) {
      tmp[idx++] = c;
    } else {
      break;
    }
  }
  tmp[idx] = '\0';
  return tmp;
}

#define LEN 2
int main(int argc, char const* argv[]) {
  char* strs[LEN] = {"aca","cba"};
  char* str = longestCommonPrefix(strs, LEN);
  printf("%s\n", str);
  return 0;
}
