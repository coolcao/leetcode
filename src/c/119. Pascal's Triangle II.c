/*
Given a non-negative index k where k ≤ 33, return the kth index row of the
Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above
it.

Example:

Input: 3
Output: [1,3,3,1]

**********************
生成杨辉三角的第n行
杨辉三角：
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/
#include <stdio.h>
#include <stdlib.h>
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* getRow(int rowIndex, int* returnSize) {
  *returnSize = rowIndex + 1;
  int* pre = malloc(sizeof(int) * (*returnSize));
  pre[0] = 1;
  if (rowIndex == 0) return pre;

  pre[1] = 1;
  if (rowIndex == 1) return pre;

  int* current = malloc(sizeof(int) * (*returnSize));
  current[0] = 1;
  current[1] = 1;

  for (int row = 2; row <= rowIndex; row++) {
    for (int i = 1; i <= row; i++) {
      if (i == row) {
        current[i] = 1;
      } else {
        current[i] = pre[i] + pre[i - 1];
      }
    }
    int* tmp = pre;
    pre = current;
    current = tmp;
  }
  return pre;
}

int main(int argc, char const* argv[]) {
  int returnSize = 0;
  int rowIndex = 5;
  int* result = getRow(rowIndex, &returnSize);
  printf("returnSize: %d\n", returnSize);
  for (int i = 0; i < returnSize; i++) {
    printf("%d\t", result[i]);
  }
  printf("\n");

  return 0;
}
