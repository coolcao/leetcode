/*
 * @lc app=leetcode id=1160 lang=golang
 *
 * [1160] Find Words That Can Be Formed by Characters
 *
 * https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/description/
 *
 * algorithms
 * Easy (67.31%)
 * Likes:    359
 * Dislikes: 68
 * Total Accepted:    49.2K
 * Total Submissions: 72.9K
 * Testcase Example:  '["cat","bt","hat","tree"]\n"atach"'
 *
 * You are given an array of strings words and a string chars.
 *
 * A string is good if it can be formed by characters from chars (each
 * character can only be used once).
 *
 * Return the sum of lengths of all good strings in words.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["cat","bt","hat","tree"], chars = "atach"
 * Output: 6
 * Explanation:
 * The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 =
 * 6.
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
 * Output: 10
 * Explanation:
 * The strings that can be formed are "hello" and "world" so the answer is 5 +
 * 5 = 10.
 *
 * Note:
 *
 * 1 <= words.length <= 1000
 * 1 <= words[i].length, chars.length <= 100
 * All strings contain lowercase English letters only.
 *
 */
package main

import "fmt"

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Printf("%v\n", countCharacters(words, chars))
}

// @lc code=start
func contain(word string, chars string) int {
	count := 0
	hmap := map[byte]int{}
	for i := 0; i < len(chars); i++ {
		hmap[chars[i]]++
	}
	for i := 0; i < len(word); i++ {
		if hmap[word[i]] == 0 {
			return 0
		}
		hmap[word[i]]--
		count++
	}
	return count
}
func countCharacters(words []string, chars string) int {
	sum := 0
	for _, w := range words {
		sum += contain(w, chars)
	}
	return sum
}

// @lc code=end
