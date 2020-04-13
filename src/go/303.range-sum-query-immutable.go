/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 *
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (42.55%)
 * Likes:    748
 * Dislikes: 970
 * Total Accepted:    183.4K
 * Total Submissions: 430.6K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * Example:
 *
 * Given nums = [-2, 0, 3, -5, 2, -1]
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 *
 *
 * Note:
 *
 * You may assume that the array does not change.
 * There are many calls to sumRange function.
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{}
	na := Constructor(nums)
	c := na.SumRange(0, 0)
	fmt.Printf("c: %d\n", c)
}

// @lc code=start

// ---------- 第一种方法，把所有可能列出来，直接返回-----------
// type NumArray struct {
// 	length   int
// 	data     []int
// 	sumCount []int
// }
// func Constructor(nums []int) NumArray {
// 	length := len(nums)
// 	na := new(NumArray)
// 	all := (length + 1) * length / 2
// 	na.data = make([]int, length)
// 	na.sumCount = make([]int, all)
// 	na.length = length

// 	a := 0
// 	for i := 0; i < length; i++ {
// 		na.data[i] = nums[i]
// 		sum := nums[i]
// 		for j := i; j < length; j++ {
// 			if i == j {
// 				na.sumCount[a] = sum
// 			} else {
// 				sum += nums[j]
// 				na.sumCount[a] = sum
// 			}
// 			a++
// 		}
// 	}

// 	return *na
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	idx := 0
// 	for x := 0; x < i; x++ {
// 		idx += (this.length - x)
// 	}
// 	idx += (j - i)
// 	return this.sumCount[idx]
// }

// -----------------第二种方案------------------
type NumArray struct {
	length   int
	sumCount []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)

	if length == 0 {
		return NumArray{
			length:   0,
			sumCount: []int{},
		}
	}

	na := new(NumArray)
	na.length = length
	na.sumCount = make([]int, length)
	na.sumCount[0] = nums[0]

	for i := 1; i < length; i++ {
		na.sumCount[i] = nums[i] + na.sumCount[i-1]
	}

	fmt.Printf("length: %d\n", na.length)
	fmt.Printf("%v\n", na.sumCount)
	return *na
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		if this.length == 0 {
			return 0
		}
		return this.sumCount[j]
	}
	return this.sumCount[j] - this.sumCount[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
