/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 *
 * https://leetcode.com/problems/guess-number-higher-or-lower/description/
 *
 * algorithms
 * Easy (41.78%)
 * Likes:    338
 * Dislikes: 1566
 * Total Accepted:    140.5K
 * Total Submissions: 335.9K
 * Testcase Example:  '10\n6'
 *
 * We are playing the Guess Game. The game is as follows:
 *
 * I pick a number from 1 to n. You have to guess which number I picked.
 *
 * Every time you guess wrong, I'll tell you whether the number is higher or
 * lower.
 *
 * You call a pre-defined API guess(int num) which returns 3 possible results
 * (-1, 1, or 0):
 *
 *
 * -1 : My number is lower
 * ⁠1 : My number is higher
 * ⁠0 : Congrats! You got it!
 *
 *
 * Example :
 *
 *
 *
 * Input: n = 10, pick = 6
 * Output: 6
 *
 *
 *
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
package main

import "fmt"

func main() {
	n := 10
	fmt.Println(guessNumber(n))
}
func guess(num int) int {
	n := 6
	if num < n {
		return -1
	}
	if num > n {
		return 1
	}
	return 0
}
func guessNumber(n int) int {
	start, end := 1, n
	for start <= end {
		mid := (start + end) / 2
		g := guess(mid)
		if g == 0 {
			return mid
		}
		if g == -1 {
			start = mid + 1
		}
		if g == 1 {
			end = mid - 1
		}
	}
	return -1
}

// @lc code=end
