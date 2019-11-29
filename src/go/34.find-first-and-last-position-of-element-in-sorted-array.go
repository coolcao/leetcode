/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

*/
package main

import "fmt"

func binarySearch(nums []int, start, end int, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	}
	if nums[mid] < target {
		return binarySearch(nums, mid+1, end, target)
	}
	return -1
}

// @lc code=start
func searchRange(nums []int, target int) []int {
	length := len(nums)
	idx := binarySearch(nums, 0, length-1, target)
	if idx == -1 {
		return []int{-1, -1}
	}
	idxStart, idxEnd := idx, idx
	result := make([]int, 2)

	for idxStart >= 0 && nums[idxStart] == target {
		idxStart--
	}
	for idxEnd < length && nums[idxEnd] == target {
		idxEnd++
	}
	result[0] = idxStart + 1
	result[1] = idxEnd - 1
	return result
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 7
	fmt.Printf("%v\n", searchRange(nums, target))
}

// @lc code=end
