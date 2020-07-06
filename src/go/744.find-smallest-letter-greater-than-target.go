/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 *
 * https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
 *
 * algorithms
 * Easy (44.99%)
 * Likes:    387
 * Dislikes: 501
 * Total Accepted:    73.4K
 * Total Submissions: 162K
 * Testcase Example:  '["c","f","j"]\n"a"'
 *
 *
 * Given a list of sorted characters letters containing only lowercase letters,
 * and given a target letter target, find the smallest element in the list that
 * is larger than the given target.
 *
 * Letters also wrap around.  For example, if the target is target = 'z' and
 * letters = ['a', 'b'], the answer is 'a'.
 *
 *
 * Examples:
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "a"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "c"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "d"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "g"
 * Output: "j"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "j"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "k"
 * Output: "c"
 *
 *
 *
 * Note:
 *
 * letters has a length in range [2, 10000].
 * letters consists of lowercase letters, and contains at least 2 unique
 * letters.
 * target is a lowercase letter.
 *
 *
 */
package main

import "fmt"

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('d')
	res := nextGreatestLetter(letters, target)
	fmt.Printf("%c\n", res)

}

// @lc code=start
func search(letters []byte, target byte) int {
	start, end := 0, len(letters)-1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if letters[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Printf("mid: %d\n", mid)
	return mid
}
func nextGreatestLetter(letters []byte, target byte) byte {
	idx := search(letters, target)
	if idx == len(letters)-1 && letters[idx] <= target {
		return letters[0]
	}
	if letters[idx] <= target {
		return letters[idx+1]
	}
	return letters[idx]
}

// @lc code=end
