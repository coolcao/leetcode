/**
 Given an array nums of n integers, are there elements a, b, c in nums such that
a + b + c = 0? Find all unique triplets in the array which gives the sum of
zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
====
Medium
====
 */

#include <stdio.h>
# include <stdlib.h>
#include "../common/uthash.h"


struct _map {
  char* key;
  UT_hash_handle hh;
};
typedef struct _map *Map;
Map map = NULL;

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

/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int** threeSum(int* nums, int numsSize, int* returnSize) {
  // 先排序
  sort(nums, numsSize, 0, numsSize - 1);

  int** result = (int**)malloc(sizeof(int*)*numsSize*numsSize);
  int resutIdx = 0;

  for (int i = 0; i < numsSize; i++) {
    int sub = 0 - nums[i];
    int start = i + 1;
    int end = numsSize - 1;
    while (start < end) {
      int two_sum = nums[start] + nums[end];
      if (two_sum < sub) {
        start += 1;
      } else if (two_sum > sub) {
        end -= 1;
      } else {
        // 三个数相加等于0，则返回相应的数组
        char* strkey = (char*)malloc(sizeof(char*)*4);
        strkey[0] = nums[i];
        strkey[1] = nums[start];
        strkey[2] = nums[end];
        strkey[3] = '\0';
        Map m = NULL;
        HASH_FIND_STR(map, strkey, m);
        if (m == NULL) {
          m = (Map)malloc(sizeof(*m));
          m->key = strkey;
          HASH_ADD_STR(map, key, m);

          int* tmp = (int*)malloc(sizeof(int*)*3);
          tmp[0] = nums[i];
          tmp[1] = nums[start];
          tmp[2] = nums[end];

          result[resutIdx++] = tmp;
          printf("%d %d %d\n", result[resutIdx-1][0],result[resutIdx-1][1],result[resutIdx-1][2]);
        }
        start += 1;
        end -= 1;
      }
    }
  }
  *returnSize = resutIdx;
  return result;
}



int main(int argc, char const* argv[]) {
  int a[20] = {-1,-2,-3,-4,-5,-6,-7,-8,-9,0,1,2,3,4,5,6,7,8,9,0};
  int returnSize = 0;
  int** result = threeSum(a, 20, &returnSize);
  printf("size:%d\n", returnSize);
  return 0;
}
