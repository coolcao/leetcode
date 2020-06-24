/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 *
 * https://leetcode.com/problems/binary-search/description/
 *
 * algorithms
 * Easy (50.83%)
 * Likes:    645
 * Dislikes: 43
 * Total Accepted:    135K
 * Total Submissions: 261.7K
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * Given a sorted (in ascending order) integer array nums of n elements and a
 * target value, write a function to search target in nums. If target exists,
 * then return its index, otherwise return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,0,3,5,9,12], target = 9
 * Output: 4
 * Explanation: 9 exists in nums and its index is 4
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-1,0,3,5,9,12], target = 2
 * Output: -1
 * Explanation: 2 does not exist in nums so return -1
 *
 *
 *
 *
 * Note:
 *
 *
 * You may assume that all elements in nums are unique.
 * n will be in the range [1, 10000].
 * The value of each element in nums will be in the range [-9999, 9999].
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 8, 9, 10, 12, 14, 19, 23}
	target := 18
	idx := search(nums, target)
	fmt.Printf("%v\n", idx)
}

// @lc code=start

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		}
		if nums[mid] < target {
			start = mid + 1
		}
	}
	return -1
}

// @lc code=end
