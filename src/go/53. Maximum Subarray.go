package main

import (
	"fmt"
)

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max := nums[0]
	start, end := 0, 0

	for i := 0; i < length; i++ {
		tmpmax := nums[i]
		for j := i; j < length; j++ {
			if i != j {
				tmpmax += nums[j]
			}
			if max < tmpmax {
				max = tmpmax
				start, end = i, j
			}
		}
	}
	fmt.Printf("%v\n", nums[start:end+1])
	return max

}

func maxSubArray2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Printf("%v\n", dp)
	return max
}

func main() {
	nums := []int{-2, -1, -3, -4, -5, -6}
	max := maxSubArray(nums)
	max2 := maxSubArray2(nums)
	fmt.Printf("max=%d\n", max)
	fmt.Printf("max=%d\n", max2)
}
