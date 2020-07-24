/*
 * @lc app=leetcode id=1189 lang=golang
 *
 * [1189] Maximum Number of Balloons
 *
 * https://leetcode.com/problems/maximum-number-of-balloons/description/
 *
 * algorithms
 * Easy (60.97%)
 * Likes:    277
 * Dislikes: 35
 * Total Accepted:    36.7K
 * Total Submissions: 60K
 * Testcase Example:  '"nlaebolko"'
 *
 * Given a string text, you want to use the characters of text to form as many
 * instances of the word "balloon" as possible.
 *
 * You can use each character in text at most once. Return the maximum number
 * of instances that can be formed.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: text = "nlaebolko"
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: text = "loonbalxballpoon"
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: text = "leetcode"
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= text.length <= 10^4
 * text consists of lower case English letters only.
 *
 */
package main

import "fmt"

func main() {
	text := "nlaebollllllllko"
	max := maxNumberOfBalloons(text)
	fmt.Printf("%v\n", max)
}

// @lc code=start
func maxNumberOfBalloons(text string) int {
	counts := [5]int{}
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case byte('b'):
			counts[0]++
		case byte('a'):
			counts[1]++
		case byte('l'):
			counts[2]++
		case byte('o'):
			counts[3]++
		case byte('n'):
			counts[4]++
		}
	}
	fmt.Printf("%v\n", counts)
	min := len(text)
	for i := 0; i < 5; i++ {
		if i == 2 || i == 3 {
			if counts[i]/2 < min {
				min = counts[i] / 2
			}
		} else {
			if counts[i] < min {
				min = counts[i]
			}
		}
	}
	return min
}

// @lc code=end
