/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

package main

import "fmt"

// @lc code=start
func rotate(nums []int, k int) {
	for k > 7 {
		k = k % 7
	}
	if k == 0 {
		return
	}
	length := len(nums)
	idx := length - 1
	value := nums[idx]
	for i := 0; i < k; i++ {
		value = nums[idx]
		for j := length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = value
	}
}

func rotate2(nums []int, k int) {
	for k > 7 {
		k = k % 7
	}
	if k == 0 {
		return
	}
	length := len(nums)
	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = nums[length-1-i]
	}
	for i := length - 1; i >= 0; i-- {
		if i >= k {
			nums[i] = nums[i-k]
		} else {
			nums[i] = tmp[k-1-i]
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 8
	rotate2(nums, k)
	fmt.Printf("%v\n", nums)
}

// @lc code=end
