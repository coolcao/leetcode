/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

package main

import "fmt"

// 将一个元素插入到一个已经排好序的数组
func insert(nums []int, length int, value int) {
	if length == 0 {
		nums[0] = value
	}
	for idx := length - 1; idx >= 0; idx-- {
		if value < nums[idx] {
			nums[idx+1] = nums[idx]
			nums[idx] = value
		} else {
			nums[idx+1] = value
			break
		}
	}
}

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	for index := 0; index < n; index++ {
		value := nums2[index]
		insert(nums1, m+index, value)
		fmt.Printf("%v\n", nums1)
	}
}

func main() {
	nums1 := []int{1, 2, 0}
	m := 2
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Printf("%v\n", nums1)
}

// @lc code=end
