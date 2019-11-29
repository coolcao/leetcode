package main

import "fmt"

/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
*/

// @lc code=start
func firstMissingPositive(nums []int) int {
	length := len(nums)
	var tmp int
	for i := 0; i < length; i++ {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] >= length {
				break
			}
			if nums[i] == nums[nums[i]-1] {
				break
			}
			tmp = nums[i]
			nums[i] = nums[tmp-1]
			nums[tmp-1] = tmp
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}
func main() {
	nums := []int{0, 9, 7, 5, 3, 1, 2, 5, 6}
	num := firstMissingPositive(nums)
	fmt.Printf("%v\n", num)
}

// @lc code=end
