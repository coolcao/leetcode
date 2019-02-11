/*
 * @lc app=leetcode id=26 lang=c
 *
 * [26] Remove Duplicates from Sorted Array
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
 *
 * algorithms
 * Easy (39.35%)
 * Total Accepted:    511.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted array nums, remove the duplicates in-place such that each
 * element appear only once and return the new length.
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * Example 1:
 *
 *
 * Given nums = [1,1,2],
 *
 * Your function should return length = 2, with the first two elements of nums
 * being 1 and 2 respectively.
 *
 * It doesn't matter what you leave beyond the returned length.
 *
 * Example 2:
 *
 *
 * Given nums = [0,0,1,1,1,2,2,3,3,4],
 *
 * Your function should return length = 5, with the first five elements of nums
 * being modified to 0, 1, 2, 3, and 4 respectively.
 *
 * It doesn't matter what values are set beyond the returned length.
 *
 *
 * Clarification:
 *
 * Confused why the returned value is an integer but your answer is an array?
 *
 * Note that the input array is passed in by reference, which means
 * modification to the input array will be known to the caller as well.
 *
 * Internally you can think of this:
 *
 *
 * // nums is passed in by reference. (i.e., without making a copy)
 * int len = removeDuplicates(nums);
 *
 * // any modification to nums in your function would be known by the caller.
 * // using the length returned by your function, it prints the first len
 * elements.
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 *
 */
#include <stdio.h>
// int removeDuplicates(int* nums, int numsSize) {
//   if (numsSize <= 1) {
//     return numsSize;
//   }
//   int current = nums
//       [0];  // 由于数组是已排好序的，而且题目要求使用O(1)的额外空间来存储，所以使用一个变量来存储已遍历到的数
//   int removed = 0;  // 记录已删除的元素个数
//   for (int i = 1; i < numsSize; i++) {
//     // 当前元素和之前的元素相同，要删除操作
//     int idx = i - removed;
//     if (nums[idx] == nums[idx - 1]) {
//       for (int j = idx; j < numsSize - removed; j++) {
//         nums[j] = nums[j + 1];
//       }
//       removed += 1;
//     } else {
//       // 当前元素和之前元素不同，不进行删除，替换current为当前元素
//       current = nums[idx];
//     }
//   }
//   return numsSize - removed;
// }

int removeDuplicates(int* nums, int numsSize) {
  if (numsSize <= 1) return numsSize;
  int current = nums[0];
  int currentIdx = 1;
  for (int i = 1; i < numsSize; i++) {
    if (nums[i] == current) {
      continue;
    } else {
      current = nums[i];
      nums[currentIdx] = current;
      currentIdx += 1;
    }
  }
  return currentIdx;
}

// int main(int argc, char const* argv[]) {
// #define LEN 20
//   int a[LEN] = {0, 0, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6};
//   int len = removeDuplicates(a, LEN);
//   printf("%d\n", len);
//   for (int i = 0; i < len; i++) {
//     printf("%d ", a[i]);
//   }
//   printf("\n");
//   return 0;
// }
