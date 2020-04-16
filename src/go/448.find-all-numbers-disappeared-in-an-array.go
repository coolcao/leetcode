/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (55.38%)
 * Likes:    2446
 * Dislikes: 223
 * Total Accepted:    224.1K
 * Total Submissions: 404.5K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
 * elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the
 * returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDisappearedNumbers2(nums)
	fmt.Printf("%v\n", result)
}

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	for _, v := range nums {
		result[v-1] = v
	}
	count := 0
	for i, v := range result {
		if v == 0 {
			result[count] = i + 1
			count++
		}
	}
	return result[0:count]
}

// 原地修改
// 我们遍历一次数组，然后将每个元素对应位置上的数改为负数
// 最后再遍历一遍数组，为负数的说明出现过，正数的说明未出现过
func getPositive(num int) int {
	if num > 0 {
		return num
	}
	return -1 * num
}
func getNegative(num int) int {
	if num < 0 {
		return num
	}
	return -1 * num
}
func findDisappearedNumbers2(nums []int) []int {
	for _, v := range nums {
		nums[getPositive(v)-1] = getNegative(nums[getPositive(v)-1])
	}
	count := 0
	for i, v := range nums {
		if v > 0 {
			nums[count] = i + 1
			count++
		}
	}
	return nums[:count]
}

// @lc code=end
