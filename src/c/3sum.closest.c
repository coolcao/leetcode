/**
Given an array nums of n integers and an integer target, find three integers in
nums such that the sum is closest to target. Return the sum of the three
integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
====
Medium
====
 */

#include <math.h>
#include <stdio.h>

void sort(int a[], int length, int start, int end) {
  if (start >= end) {
    return;
  }
  int idx = start;
  for (int i = start + 1; i <= end; i++) {
    // 将所有小于基准的值放到该基准值的前面
    if (a[idx] > a[i]) {
      int tmp = a[i];
      int sidx = i;
      while (sidx > idx) {
        a[sidx] = a[sidx - 1];
        sidx -= 1;
      }
      a[idx] = tmp;
      idx += 1;
    }
  }
  if (idx > 0) {
    sort(a, length, start, idx - 1);
  }
  if (idx < length - 1) {
    sort(a, length, idx + 1, end);
  }
}

int threeSumClosest(int* nums, int numsSize, int target) {
  // 使用快速排序先将原数组排序
  sort(nums, numsSize, 0, numsSize - 1);

  int minSub = 2147483647;
  int sum = 0;

  for (int i=0;i<numsSize;i++) {
    printf("%d ", nums[i]);
  }
  printf("\n");

  for (int i = 0; i < numsSize - 2; i++) {
    int start = i + 1, end = numsSize - 1;
    while (start < end) {
      int tmpsum = nums[i] + nums[start] + nums[end];
      int sub = target - tmpsum;
      int sub_abs = sub<0?-sub:sub;
      if (sub_abs < minSub) {
        minSub = sub_abs;
        sum = tmpsum; 
      }
      // printf("sub_abs:%d\n", sub_abs);
      // printf("\n");
      if (sub < 0) {
        end -= 1;
      } else if (sub > 0) {
        start += 1;
      } else {
        break;
      }
    }
  }
  return sum;
}

#define LEN 15
int main(int argc, char const* argv[]) {
  int nums[LEN] = {-1, 2, 1, -4, 3, -2, 8, 6, 5, 7, 12, 24, 25, 18, 15};
  int target = 344;
  int sum = threeSumClosest(nums, LEN, target);
  printf("%d\n", sum);
  return 0;
}
