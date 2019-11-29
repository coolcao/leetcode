/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/
package main

import "fmt"

// @lc code=start
func missingNumber(nums []int) int {
	length := len(nums)
	sum := length * (length + 1) >> 1
	for idx := 0; idx < length; idx++ {
		sum -= nums[idx]
	}
	return sum
}

func main() {
	nums := []int{3, 0, 2}
	num := missingNumber(nums)
	fmt.Printf("%d\n", num)
}

// @lc code=end
