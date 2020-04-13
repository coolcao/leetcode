/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 *
 * https://leetcode.com/problems/first-bad-version/description/
 *
 * algorithms
 * Easy (33.48%)
 * Likes:    1059
 * Dislikes: 580
 * Total Accepted:    314.4K
 * Total Submissions: 937.8K
 * Testcase Example:  '5\n4'
 *
 * You are a product manager and currently leading a team to develop a new
 * product. Unfortunately, the latest version of your product fails the quality
 * check. Since each version is developed based on the previous version, all
 * the versions after a bad version are also bad.
 *
 * Suppose you have n versions [1, 2, ..., n] and you want to find out the
 * first bad one, which causes all the following ones to be bad.
 *
 * You are given an API bool isBadVersion(version) which will return whether
 * version is bad. Implement a function to find the first bad version. You
 * should minimize the number of calls to the API.
 *
 * Example:
 *
 *
 * Given n = 5, and version = 4 is the first bad version.
 *
 * call isBadVersion(3) -> false
 * call isBadVersion(5) -> true
 * call isBadVersion(4) -> true
 *
 * Then 4 is the first bad version.
 *
 *
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

package main

import "fmt"

// 模拟isBadVersion接口
func isBadVersion(version int) bool {
	v := 1150769282
	if version >= v {
		return true
	}
	return false
}
func firstBadVersion(n int) int {
	start, end := 1, n
	mid := (start + end) / 2
	for mid != start {
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid
		}
		mid = (start + end) / 2
	}
	if isBadVersion(mid) {
		return mid
	}
	return mid + 1
}

func main() {
	n := 1420736637
	fmt.Println(firstBadVersion(n))
}

// @lc code=end
