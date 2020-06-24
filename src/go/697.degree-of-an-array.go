/*
 * @lc app=leetcode id=697 lang=golang
 *
 * [697] Degree of an Array
 *
 * https://leetcode.com/problems/degree-of-an-array/description/
 *
 * algorithms
 * Easy (53.18%)
 * Likes:    842
 * Dislikes: 760
 * Total Accepted:    81.2K
 * Total Submissions: 151.6K
 * Testcase Example:  '[1,2,2,3,1]'
 *
 * Given a non-empty array of non-negative integers nums, the degree of this
 * array is defined as the maximum frequency of any one of its elements.
 * Your task is to find the smallest possible length of a (contiguous) subarray
 * of nums, that has the same degree as nums.
 *
 * Example 1:
 *
 * Input: [1, 2, 2, 3, 1]
 * Output: 2
 * Explanation:
 * The input array has a degree of 2 because both elements 1 and 2 appear
 * twice.
 * Of the subarrays that have the same degree:
 * [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
 * The shortest length is 2. So return 2.
 *
 *
 *
 *
 * Example 2:
 *
 * Input: [1,2,2,3,1,4,2]
 * Output: 6
 *
 *
 *
 * Note:
 * nums.length will be between 1 and 50,000.
 * nums[i] will be an integer between 0 and 49,999.
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4, 5, 4, 3, 2, 2, 1}
	fmt.Printf("min:%v\n", findShortestSubArray(nums))
}

// @lc code=start
func getDegree(nums []int) int {
	hmap := map[int]int{}
	degree := 0
	for _, n := range nums {
		hmap[n]++
		if hmap[n] > degree {
			degree = hmap[n]
		}
	}
	return degree
}
func findShortestSubArray(nums []int) int {
	// 先计算数组的度
	degree := getDegree(nums)
	length := len(nums)

	// 使用滑动窗口，计算最小子数组长度
	left, right := 0, 0

	// 保存窗口内元素计数
	windowMap := map[int]int{}
	// 记录窗口的度
	windowDegree := 0

	// 记录窗口最小长度
	min := length

	for right < length {
		// 入窗
		windowMap[nums[right]]++
		if windowDegree < windowMap[nums[right]] {
			windowDegree = windowMap[nums[right]]
		}
		right++

		for left < right && windowDegree == degree {
			if windowDegree == windowMap[nums[left]] {
				windowDegree--
			}
			windowMap[nums[left]]--

			if min > right-left {
				min = right - left
			}

			left++
		}

	}

	return min
}

// @lc code=end
