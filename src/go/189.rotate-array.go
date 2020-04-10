/*
 * @lc app=leetcode id=189 lang=golang
 * [189] Rotate Array

Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?

难度：简单
*/

package main

import "fmt"

// @lc code=start
func rotate(nums []int, k int) {
	length := len(nums)
	if k >= length {
		k = k % length
	}
	if k == 0 {
		return
	}
	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = nums[length-k+i]
	}
	for i := length - 1; i >= k; i-- {
		nums[i], nums[i-k] = nums[i-k], nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}

func rotate2(nums []int, k int) {
	length := len(nums)
	if k >= length {
		k = k % length
	}
	if k == 0 {
		return
	}
	tmp := make([]int, length)
	for i := 0; i < k; i++ {
		tmp[i] = nums[length-k+i]
	}
	for i := 0; i < length-k; i++ {
		tmp[k+i] = nums[i]
	}
	for i := 0; i < length; i++ {
		nums[i] = tmp[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Printf("%v\n", nums)
}

// @lc code=end
