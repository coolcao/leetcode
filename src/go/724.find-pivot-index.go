/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 *
 * https://leetcode.com/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (43.15%)
 * Likes:    1059
 * Dislikes: 239
 * Total Accepted:    130.3K
 * Total Submissions: 298.3K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * Given an array of integers nums, write a method that returns the "pivot"
 * index of this array.
 *
 * We define the pivot index as the index where the sum of all the numbers to
 * the left of the index is equal to the sum of all the numbers to the right of
 * the index.
 *
 * If no such index exists, we should return -1. If there are multiple pivot
 * indexes, you should return the left-most pivot index.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,7,3,6,5,6]
 * Output: 3
 * Explanation:
 * The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the
 * sum of numbers to the right of index 3.
 * Also, 3 is the first index where this occurs.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3]
 * Output: -1
 * Explanation:
 * There is no index that satisfies the conditions in the problem
 * statement.
 *
 *
 *
 * Constraints:
 *
 *
 * The length of nums will be in the range [0, 10000].
 * Each element nums[i] will be an integer in the range [-1000, 1000].
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{-1, -1, -1, -1, -1, 0}
	idx := pivotIndex(nums)
	fmt.Printf("idx: %d\n", idx)
}

// @lc code=start
func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	left := 0
	for i := 0; i < len(nums); i++ {
		if left*2+nums[i] == sum {
			return i
		}
		left += nums[i]
	}
	return -1
}

// @lc code=end
