/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.
*/
package main

import "fmt"

// @lc code=start
func isIsomorphic(s string, t string) bool {
	length := len(s)
	charMap := make(map[byte]byte)  // 正向映射
	rcharMap := make(map[byte]bool) // 反向映射
	for i := 0; i < length; i++ {
		cs, ct := s[i], t[i]
		if ec, ok := charMap[cs]; ok {
			if ec != ct {
				return false
			}
		} else {
			if _, ok := rcharMap[ct]; ok {
				return false
			}
			charMap[cs] = ct
			rcharMap[ct] = true
		}
	}
	return true
}
func main() {
	s := "eggxy"
	t := "addxx"
	b := isIsomorphic(s, t)
	fmt.Printf("%v\n", b)
}

// @lc code=end
