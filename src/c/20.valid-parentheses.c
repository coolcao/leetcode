/*
 * @lc app=leetcode id=20 lang=c
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (35.72%)
 * Total Accepted:    501.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

#include <stdbool.h>
#include <stdio.h>
#include <string.h>

bool isValid(char* s) {
  int length = strlen(s);
  if (length == 0) return true;
  char str[length];
  str[0] = s[0];
  int current = 0;
  for (int i = 1; i < length; i++) {
    if (current == -1) {
      current += 1;
      str[current] = s[i];
      continue;
    }
    if (s[i] == ')') {
      if (str[current] == '(') {
        str[current] = '0';
        current -= 1;
        continue;
      }
    } else if (s[i] == '}') {
      if (str[current] == '{') {
        str[current] = '0';
        current -= 1;
        continue;
      }
    } else if (s[i] == ']') {
      if (str[current] == '[') {
        str[current] = '0';
        current -= 1;
        continue;
      }
    }
    current += 1;
    str[current] = s[i];
  }
  return current == -1;
}

// int main(int argc, char const* argv[]) {
//   char* s = "()[]{([]}";
//   bool b = isValid(s);
//   printf("%d\n", b);
//   return 0;
// }
