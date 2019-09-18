/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one
sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m +
n) to hold additional elements from nums2. Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

#include <stdio.h>

void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n) {
  int length = m + n;
  int i = length - 1, i1 = m - 1, i2 = n - 1;
  while (i1 >= 0 && i2 >= 0) {
    if (nums1[i1] > nums2[i2]) {
      nums1[i] = nums1[i1];
      i1--;
    } else {
      nums1[i] = nums2[i2];
      i2--;
    }
    i--;
  }
  while (i1 >= 0) {
    nums1[i] = nums1[i1];
    i1--;
    i--;
  }
  while (i2 >= 0) {
    nums1[i] = nums2[i2];
    i2--;
    i--;
  }
}

int main(int argc, char const* argv[]) {
#define nums1Size 6
#define nums2Size 3
  int nums1[nums1Size] = {1, 2, 3, 0, 0, 0};
  int nums2[nums2Size] = {2, 5, 6};
  merge(nums1, nums1Size, 3, nums2, nums2Size, 3);
  for (int i = 0; i < 6; i++) {
    printf("%d\t", nums1[i]);
  }
  printf("\n");

  return 0;
}
