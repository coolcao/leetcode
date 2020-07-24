/*
 * @lc app=leetcode id=1200 lang=golang
 *
 * [1200] Minimum Absolute Difference
 *
 * https://leetcode.com/problems/minimum-absolute-difference/description/
 *
 * algorithms
 * Easy (66.55%)
 * Likes:    300
 * Dislikes: 21
 * Total Accepted:    38.4K
 * Total Submissions: 57.6K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Given an array of distinct integers arr, find all pairs of elements with the
 * minimum absolute difference of any two elements.
 *
 * Return a list of pairs in ascending order(with respect to pairs), each pair
 * [a, b] follows
 *
 *
 * a, b are from arr
 * a < b
 * b - a equals to the minimum absolute difference of any two elements in
 * arr
 *
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [4,2,1,3]
 * Output: [[1,2],[2,3],[3,4]]
 * Explanation: The minimum absolute difference is 1. List all pairs with
 * difference equal to 1 in ascending order.
 *
 * Example 2:
 *
 *
 * Input: arr = [1,3,6,10,15]
 * Output: [[1,3]]
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [3,8,-10,23,19,-4,-14,27]
 * Output: [[-14,-10],[19,23],[23,27]]
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= arr.length <= 10^5
 * -10^6 <= arr[i] <= 10^6
 *
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	result := minimumAbsDifference(arr)
	fmt.Printf("%v\n", result)
}

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	result := [][]int{}
	sort.Ints(arr)
	min := 1000000
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < min {
			min = diff
		}
	}
	for i := 1; i < len(arr); i++ {
		if min == arr[i]-arr[i-1] {
			tmp := []int{arr[i-1], arr[i]}
			result = append(result, tmp)
		}
	}
	return result
}

// @lc code=end
