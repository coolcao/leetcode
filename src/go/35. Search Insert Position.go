package main

import "fmt"

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0

*/

func searchInsert2(nums []int, target int) int {
	result := -1
	length := len(nums)
	if length == 0 {
		return result
	}
	for idx, val := range nums {
		if val == target {
			return idx
		}
		if val > target {
			result = idx
			break
		}
	}
	if result == -1 {
		result = length
	}
	return result
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	start, end := 0, length-1
	for start <= end {
		mid := (start + end) / 2
		fmt.Printf("%d,%d,%d\n", start, mid, end)
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return start
}

func main() {
	nums := []int{1, 2, 3, 5}
	target := 4
	result := searchInsert(nums, target)
	fmt.Printf("%d\n", result)
}
