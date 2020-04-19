/*
 * @lc app=leetcode id=1394 lang=golang
 *
 * [1394] Find Lucky Integer in an Array
 *
 * https://leetcode.com/problems/find-lucky-integer-in-an-array/description/
 *
 * algorithms
 * Easy (68.74%)
 * Likes:    91
 * Dislikes: 5
 * Total Accepted:    16.8K
 * Total Submissions: 24.9K
 * Testcase Example:  '[2,2,3,4]'
 *
 * Given an array of integers arr, a lucky integer is an integer which has a
 * frequency in the array equal to its value.
 *
 * Return a lucky integer in the array. If there are multiple lucky integers
 * return the largest of them. If there is no lucky integer return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [2,2,3,4]
 * Output: 2
 * Explanation: The only lucky number in the array is 2 because frequency[2] ==
 * 2.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [1,2,2,3,3,3]
 * Output: 3
 * Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [2,2,2,3,3]
 * Output: -1
 * Explanation: There are no lucky numbers in the array.
 *
 *
 * Example 4:
 *
 *
 * Input: arr = [5]
 * Output: -1
 *
 *
 * Example 5:
 *
 *
 * Input: arr = [7,7,7,7,7,7,7]
 * Output: 7
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 500
 * 1 <= arr[i] <= 500
 *
 */
package main

import "fmt"

func main() {
	arr := []int{7, 7, 7, 7, 7, 7, 7}
	lucky := findLucky(arr)
	fmt.Printf("%v\n", lucky)
}

// @lc code=start
func findLucky(arr []int) int {
	length := len(arr)
	tmp := make([]int, length+1)

	for _, v := range arr {
		if v > length {
			continue
		}
		tmp[v]++
	}

	for i := length; i > 0; i-- {
		if tmp[i] == i {
			return i
		}
	}

	return -1
}

// @lc code=end
