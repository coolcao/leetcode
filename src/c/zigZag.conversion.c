/**

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of
rows like this: (you may want to display this pattern in a fixed font for better
legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number
of rows:

string convert(string s, int numRows);

Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* convert(char* s, int numRows) {
  if (numRows == 1) return s;
  int steps[numRows];
  for (int i = 0; i < numRows; i++) {
    steps[i] = (numRows - 1 - i) * 2;
  }
  int length = strlen(s);
  char* result = (char*)malloc(sizeof(char*) * (length + 1));
  int nums = 0;
  for (int i = 0; i < numRows; i++) {
    int idx = i;
    int p1 = steps[i];
    int p2 = steps[numRows - 1 - i];
    int flag = 0;
    while (idx < length) {
      printf("%c,", s[idx]);
      char c = s[idx];
      if (flag % 2 == 0) {
        if (p1 == 0) {
          idx += p2;
        } else {
          idx += p1;
        }
      } else {
        if (p2 == 0) {
          idx += p1;
        } else {
          idx += p2;
        }
      }
      result[nums] = c;
      flag += 1;
      nums += 1;
    }
  }
  result[length] = '\0';
  return result;
}

int main(int argc, char const* argv[]) {
  char* result = convert("AB", 1);
  printf("\n%s\n", result);
  return 0;
}
