/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 *
 * https://leetcode.com/problems/jewels-and-stones/description/
 *
 * algorithms
 * Easy (84.80%)
 * Likes:    2086
 * Dislikes: 364
 * Total Accepted:    511.6K
 * Total Submissions: 593.6K
 * Testcase Example:  '"aA"\n"aAAbbbb"'
 *
 * You're given strings J representing the types of stones that are jewels, and
 * S representing the stones you have.  Each character in S is a type of stone
 * you have.  You want to know how many of the stones you have are also
 * jewels.
 *
 * The letters in J are guaranteed distinct, and all characters in J and S are
 * letters. Letters are case sensitive, so "a" is considered a different type
 * of stone from "A".
 *
 * Example 1:
 *
 *
 * Input: J = "aA", S = "aAAbbbb"
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: J = "z", S = "ZZ"
 * Output: 0
 *
 *
 * Note:
 *
 *
 * S and J will consist of letters and have length at most 50.
 * The characters in J are distinct.
 *
 *
 */
package main

import "fmt"

func main() {
	J := "z"
	S := "ZZZZ"
	c := numJewelsInStones(J, S)
	fmt.Printf("%v\n", c)
}

// @lc code=start
func numJewelsInStones(J string, S string) int {
	hmap := map[rune]bool{}
	count := 0
	for _, c := range J {
		hmap[c] = true
	}
	for _, c := range S {
		if hmap[c] {
			count++
		}
	}
	return count
}

// @lc code=end
