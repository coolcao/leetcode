/*
 * @lc app=leetcode id=1408 lang=golang
 *
 * [1408] String Matching in an Array
 *
 * https://leetcode.com/problems/string-matching-in-an-array/description/
 *
 * algorithms
 * Easy (59.49%)
 * Likes:    47
 * Dislikes: 21
 * Total Accepted:    11.8K
 * Total Submissions: 19.9K
 * Testcase Example:  '["mass","as","hero","superhero"]'
 *
 * Given an array of string words. Return all strings in words which is
 * substring of another word in any order.
 *
 * String words[i] is substring of words[j], if can be obtained removing some
 * characters to left and/or right side of words[j].
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["mass","as","hero","superhero"]
 * Output: ["as","hero"]
 * Explanation: "as" is substring of "mass" and "hero" is substring of
 * "superhero".
 * ["hero","as"] is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["leetcode","et","code"]
 * Output: ["et","code"]
 * Explanation: "et", "code" are substring of "leetcode".
 *
 *
 * Example 3:
 *
 *
 * Input: words = ["blue","green","bu"]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= words.length <= 100
 * 1 <= words[i].length <= 30
 * words[i] contains only lowercase English letters.
 * It's guaranteed that words[i] will be unique.
 *
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"dzf", "dzfay", "hrl", "ydzfo"}
	match := stringMatching(words)
	fmt.Println(match)
}

// @lc code=start
func stringMatching(words []string) []string {
	// length := len(words)
	wordLens := make([][]string, 31)
	for _, w := range words {
		wl := len(w)
		wordLens[wl] = append(wordLens[wl], w)
	}
	result := []string{}
	used := map[string]bool{}
	for i := 1; i < 30; i++ {
		for j := i + 1; j <= 30; j++ {
			if len(wordLens[i]) == 0 || len(wordLens[j]) == 0 {
				continue
			}
		L1:
			for _, x := range wordLens[i] {
				if used[x] {
					continue
				}
				for _, y := range wordLens[j] {
					if used[x] {
						continue L1
					}
					if strings.Contains(y, x) {
						used[x] = true
						result = append(result, x)
					}
				}
			}

		}
	}

	return result
}

// @lc code=end
