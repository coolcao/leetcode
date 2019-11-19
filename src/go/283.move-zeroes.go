/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.

*/
package main

import "fmt"

// @lc code=start
func moveZeroes(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	count := 0
	for idx, val := range nums {
		if val != 0 {
			tmp := nums[count]
			nums[count] = val
			nums[idx] = tmp
			count++
		}
	}
}

func main() {
	nums := []int{1, 0, 2, 0, 5, 0, 6, 8, 9, 0}
	moveZeroes(nums)
	fmt.Printf("%v\n", nums)
}

// @lc code=end
