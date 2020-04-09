/*
 * @lc app=leetcode id=22 lang=golang
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 * [22] Generate Parentheses
 难度：中等
*/

package main

import "fmt"

// @lc code=start
func generate(left, right int, out string, res *[]string) {
	if left < 0 || right < 0 || left > right {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, out)
		return
	}
	generate(left-1, right, out+"(", res)
	generate(left, right-1, out+")", res)
}
func generateParenthesis(n int) []string {
	result := []string{}
	generate(n, n, "", &result)
	return result
}

func main() {
	re := generateParenthesis(3)
	fmt.Printf("%v\n", re)
}

// @lc code=end
