/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two

Given an integer, write a function to determine if it is a power of two.

Example 1:

Input: 1
Output: true
Explanation: 20 = 1
Example 2:

Input: 16
Output: true
Explanation: 24 = 16
Example 3:

Input: 218
Output: false
*/
package main

import "fmt"

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d is power of two? %v\n", i, isPowerOfTwo(i))
	}
}

// @lc code=end
