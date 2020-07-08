/*
 * @lc app=leetcode id=1299 lang=golang
 *
 * [1299] Replace Elements with Greatest Element on Right Side
 *
 * https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/
 *
 * algorithms
 * Easy (78.17%)
 * Likes:    349
 * Dislikes: 84
 * Total Accepted:    53.7K
 * Total Submissions: 70.6K
 * Testcase Example:  '[17,18,5,4,6,1]'
 *
 * Given an array arr, replace every element in that array with the greatest
 * element among the elements to its right, and replace the last element with
 * -1.
 *
 * After doing so, return the array.
 *
 *
 * Example 1:
 * Input: arr = [17,18,5,4,6,1]
 * Output: [18,6,6,6,1,-1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 1 <= arr[i] <= 10^5
 *
 */
package main

import "fmt"

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Printf("%v\n", replaceElements(arr))
}

// @lc code=start
func replaceElements(arr []int) []int {
	length := len(arr)
	max := -1

	for i := length - 1; i >= 0; i-- {
		item := arr[i]
		arr[i] = max
		if max < item {
			max = item
		}
	}
	return arr
}

// @lc code=end
