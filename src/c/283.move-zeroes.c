/*
Given an array nums, write a function to move all 0's to the end of it while
maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/
#include <stdio.h>
#include <stdlib.h>
void moveZeroes(int* nums, int numsSize) {
  if (numsSize == 0) return;
  int count = 0;
  for (int i = 0; i < numsSize; i++) {
    if (nums[i] != 0) {
      int tmp = nums[i];
      nums[i] = nums[count];
      nums[count] = tmp;
      count++;
    }
  }
}
void print(int* nums, int numsSize) {
  for (int i = 0; i < numsSize; i++) {
    printf("%d,", nums[i]);
  }
  printf("\n");
}
int main(int argc, char const* argv[]) {
#define SIZE 10
  int nums[SIZE] = {0, 1, 0, 3, 12,0,4,0,5,9};
  moveZeroes(nums, SIZE);
  print(nums, SIZE);
  return 0;
}
