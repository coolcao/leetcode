#include <stdio.h>
#include "../common/utils.h"

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2,
                              int nums2Size) {
  int length = nums1Size + nums2Size;
  int array[length];
  int idx = 0;
  int i = 0, j = 0;
  while (i < nums1Size && j < nums2Size) {
    if (nums1[i] < nums2[j]) {
      array[idx] = nums1[i];
      i += 1;
    } else {
      array[idx] = nums2[j];
      j += 1;
    }
    idx += 1;
  }

  while (i < nums1Size) {
    array[idx] = nums1[i];
    i += 1;
    idx += 1;
  }
  while (j < nums2Size) {
    array[idx] = nums2[j];
    j += 1;
    idx += 1;
  }
  printArray(array, length);
  if (length % 2 == 0) {
    return ((double)array[length / 2] + (double)array[length / 2 - 1]) / 2;
  }
  return array[length / 2];
}

int main(int argc, char const* argv[]) {
  int a[5] = {1, 2, 5, 8, 9};
  int b[5] = {3, 4, 6, 9, 17};
  double median = findMedianSortedArrays(a, 5, b, 5);
  printf("%f\n", median);
  return 0;
}
