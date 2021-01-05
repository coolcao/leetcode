/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/
package main

import "fmt"

// @lc code=start
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

func toSearch(nums []int, start, end int, target int) int {
	if start >= end {
		if nums[start] == target {
			return start
		}
		return -1
	}
	mid := (start + end) / 2
	// 找出升序区间
	ascStart, ascEnd := start, mid
	nStart, nEnd := mid+1, end
	if nums[mid] < nums[end] {
		ascStart = mid
		ascEnd = end
		nStart = start
		nEnd = mid - 1
	}

	// 判断target是否在升序区间内，如果在，直接二分
	if target >= nums[ascStart] && target <= nums[ascEnd] {
		return binarySearch(nums, ascStart, ascEnd, target)
	}
	// 如果不在，则继续在剩余非升序区间查找
	return toSearch(nums, nStart, nEnd, target)

}

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	return toSearch(nums, 0, length-1, target)
}

// @lc code=end

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	idx := search(nums, target)
	fmt.Printf("idx: %d\n", idx)
}
