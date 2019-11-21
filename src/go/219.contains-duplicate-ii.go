/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/
package main

import "fmt"

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	idxMap := make(map[int]int)
	for idx, num := range nums {
		if lastIdx, ok := idxMap[num]; ok {
			if idx-lastIdx <= k {
				return true
			}
			idxMap[num] = idx
		} else {
			idxMap[num] = idx
		}
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	b := containsNearbyDuplicate(nums, k)
	fmt.Printf("%v\n", b)
}

// @lc code=end
