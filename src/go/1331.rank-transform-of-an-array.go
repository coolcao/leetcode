/*
 * @lc app=leetcode id=1331 lang=golang
 *
 * [1331] Rank Transform of an Array
 *
 * https://leetcode.com/problems/rank-transform-of-an-array/description/
 *
 * algorithms
 * Easy (58.63%)
 * Likes:    160
 * Dislikes: 11
 * Total Accepted:    10.4K
 * Total Submissions: 17.8K
 * Testcase Example:  '[40,10,20,30]'
 *
 * Given an array of integers arr, replace each element with its rank.
 *
 * The rank represents how large the element is. The rank has the following
 * rules:
 *
 *
 * Rank is an integer starting from 1.
 * The larger the element, the larger the rank. If two elements are equal,
 * their rank must be the same.
 * Rank should be as small as possible.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [40,10,20,30]
 * Output: [4,1,2,3]
 * Explanation: 40 is the largest element. 10 is the smallest. 20 is the second
 * smallest. 30 is the third smallest.
 *
 * Example 2:
 *
 *
 * Input: arr = [100,100,100]
 * Output: [1,1,1]
 * Explanation: Same elements share the same rank.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [37,12,28,9,100,56,80,5,12]
 * Output: [5,3,4,2,8,6,7,1,3]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= arr.length <= 10^5
 * -10^9 <= arr[i] <= 10^9
 *
 */
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	arr := []int{1, -1, 2}
	fmt.Printf("%v\n", time.Now())
	res := arrayRankTransform(arr)
	fmt.Printf("%v\n", time.Now())
	fmt.Printf("%v\n", res)
}

// @lc code=start
// func arrayRankTransform(arr []int) []int {
// 	length := len(arr)
// 	result := make([]int, length)
// 	if length == 0 {
// 		return result
// 	}

// 	min, max := 1000000001, -1000000001
// 	for _, v := range arr {
// 		if v > max {
// 			max = v
// 		}
// 		if v < min {
// 			min = v
// 		}
// 	}

// 	sortedLength := max - min + 1
// 	sortedArr := make([]int, sortedLength)

// 	for _, v := range arr {
// 		idx := v - min
// 		sortedArr[idx]++
// 	}

// 	rankMap := map[int]int{}

// 	r := 1
// 	for i, v := range sortedArr {
// 		if v != 0 {
// 			rankMap[min+i] = r
// 			r++
// 		}
// 	}

// 	for i := 0; i < length; i++ {
// 		result[i] = rankMap[arr[i]]
// 	}
// 	return result
// }

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func arrayRankTransform(arr []int) []int {
	length := len(arr)
	result := make([]int, length)
	if length == 0 {
		return result
	}

	for i, v := range arr {
		result[i] = v
	}

	sort.Sort(IntSlice(result))

	rank := 1
	pre := result[0]

	rankMap := map[int]int{}

	for _, v := range result {

		if pre != v {
			rank++
		}
		rankMap[v] = rank
		pre = v
	}

	for i, v := range arr {
		result[i] = rankMap[v]
	}

	return result
}

// @lc code=end
