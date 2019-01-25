/**
Given an array nums of n integers and an integer target, are there elements a,
b, c, and d in nums such that a + b + c + d = target? Find all unique
quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */

#include <stdio.h>
#include <stdlib.h>
#include "../common/uthash.h"
struct _map {
  char* key;
  UT_hash_handle hh;
};
typedef struct _map* Map;

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
char* getKeyFromNum(int* nums, int numsSize) {
  char* strkey = (char*)malloc(sizeof(char*) * (numsSize * 2 + 1));
  int idx = 0;

  for (int i = 0; i < numsSize; i++) {
    if (nums[i] < 0) {
      strkey[idx++] = '-';
      strkey[idx++] = -nums[i] + 48;
    } else {
      strkey[idx++] = nums[i] + 48;
    }
  }
  strkey[idx++] = '\0';
  return strkey;
}
int** fourSum(int* nums, int numsSize, int target, int* returnSize) {
  Map map = NULL;
  sort(nums, numsSize, 0, numsSize - 1);

  int** result = (int**)malloc(sizeof(int*) * numsSize * numsSize);
  int resutIdx = 0;

  for (int i = 0; i < numsSize - 1; i++) {
    for (int j = i + 1; j < numsSize; j++) {
      int sub = target - nums[i] - nums[j];
      int start = j + 1;
      int end = numsSize - 1;
      while (start < end) {
        int two_sum = nums[start] + nums[end];

        if (two_sum < sub) {
          start += 1;
        } else if (two_sum > sub) {
          end -= 1;
        } else {
          int keyNums[4] = {nums[i], nums[j], nums[start], nums[end]};
          char* strkey = getKeyFromNum(keyNums, 4);

          Map m = NULL;
          HASH_FIND_STR(map, strkey, m);

          if (m == NULL) {
            m = (Map)malloc(sizeof(*m));
            m->key = strkey;
            HASH_ADD_STR(map, key, m);

            int* tmp = (int*)malloc(sizeof(int*) * 4);
            tmp[0] = nums[i];
            tmp[1] = nums[j];
            tmp[2] = nums[start];
            tmp[3] = nums[end];

            result[resutIdx++] = tmp;
          }

          start += 1;
          end -= 1;
        }
      }
    }
  }

  *returnSize = resutIdx;
  return result;
}
int main(int argc, char const* argv[]) {
#define LEN 8
  int nums[LEN] = {-3, -2, -1, 0, 0, 1, 2, 3};
  int target = 0;
  int returnSize = 0;
  int** result = fourSum(nums, LEN, target, &returnSize);
  printf("size: %d\n", returnSize);

  for (int i = 0; i < returnSize; i++) {
    printf("%d %d %d %d\n", result[i][0], result[i][1], result[i][2],
           result[i][3]);
  }

  return 0;
}
