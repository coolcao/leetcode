/*
 * @lc app=leetcode id=1417 lang=golang
 *
 * [1417] Reformat The String
 *
 * https://leetcode.com/problems/reformat-the-string/description/
 *
 * algorithms
 * Easy (54.49%)
 * Likes:    12
 * Dislikes: 13
 * Total Accepted:    6.8K
 * Total Submissions: 12.4K
 * Testcase Example:  '"a0b1c2"'
 *
 * Given alphanumeric string s. (Alphanumeric string is a string consisting of
 * lowercase English letters and digits).
 *
 * You have to find a permutation of the string where no letter is followed by
 * another letter and no digit is followed by another digit. That is, no two
 * adjacent characters have the same type.
 *
 * Return the reformatted string or return an empty string if it is impossible
 * to reformat the string.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "a0b1c2"
 * Output: "0a1b2c"
 * Explanation: No two adjacent characters have the same type in "0a1b2c".
 * "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "leetcode"
 * Output: ""
 * Explanation: "leetcode" has only characters so we cannot separate them by
 * digits.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "1229857369"
 * Output: ""
 * Explanation: "1229857369" has only digits so we cannot separate them by
 * characters.
 *
 *
 * Example 4:
 *
 *
 * Input: s = "covid2019"
 * Output: "c2o0v1i9d"
 *
 *
 * Example 5:
 *
 *
 * Input: s = "ab123"
 * Output: "1a2b3"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 500
 * s consists of only lowercase English letters and/or digits.
 *
 *
 */
package main

import "fmt"

func main() {
	s := "abc123"
	str := reformat(s)
	fmt.Printf("%s\n", str)
}

// @lc code=start
func reformat(s string) string {
	length := len(s)
	countOfNum, countOfLetter := 0, 0
	nums, letters := []byte{}, []byte{}
	for i := 0; i < length; i++ {
		if s[i] >= 97 {
			countOfLetter++
			letters = append(letters, s[i])
		} else {
			countOfNum++
			nums = append(nums, s[i])
		}
	}

	// fmt.Printf("%v\n", nums)
	// fmt.Printf("%v\n", letters)

	large, small := countOfLetter, countOfNum
	more := "letter"
	// fmt.Printf("%v\n", large)
	// fmt.Printf("%v\n", small)
	if countOfLetter < countOfNum {
		large, small = countOfNum, countOfLetter
		more = "num"
	}

	if large-small > 1 {
		return ""
	}

	result := []byte{}
	if more == "letter" {
		for i := 0; i < small; i++ {
			result = append(result, letters[i])
			result = append(result, nums[i])
		}
		if large > small {
			result = append(result, letters[large-1])
		}
	} else {
		for i := 0; i < small; i++ {
			result = append(result, nums[i])
			result = append(result, letters[i])
		}
		if large > small {
			result = append(result, nums[large-1])
		}
	}

	return string(result)

}

// @lc code=end
