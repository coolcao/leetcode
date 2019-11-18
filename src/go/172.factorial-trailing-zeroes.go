/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
Given an integer n, return the number of trailing zeroes in n!.

Example 1:

Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.
Example 2:

Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.
Note: Your solution should be in logarithmic time complexity.

*/

package main

import "fmt"

func trailingZeroes(n int) int {
	tmp := n / 5
	result := tmp
	for tmp >= 5 {
		tmp = tmp / 5
		result += tmp
	}
	return result
}

func main() {
	n := 40
	fmt.Printf("%d\n", trailingZeroes(n))
}

// @lc code=end
