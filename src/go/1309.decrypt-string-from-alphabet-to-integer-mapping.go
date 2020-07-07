/*
 * @lc app=leetcode id=1309 lang=golang
 *
 * [1309] Decrypt String from Alphabet to Integer Mapping
 *
 * https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/description/
 *
 * algorithms
 * Easy (76.20%)
 * Likes:    253
 * Dislikes: 25
 * Total Accepted:    29.2K
 * Total Submissions: 38K
 * Testcase Example:  '"10#11#12"'
 *
 * Given a string s formed by digits ('0' - '9') and '#' . We want to map s to
 * English lowercase characters as follows:
 *
 *
 * Characters ('a' to 'i') are represented by ('1' to '9') respectively.
 * Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.
 *
 *
 * Return the string formed after mapping.
 *
 * It's guaranteed that a unique mapping will always exist.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "10#11#12"
 * Output: "jkab"
 * Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "1326#"
 * Output: "acz"
 *
 *
 * Example 3:
 *
 *
 * Input: s = "25#"
 * Output: "y"
 *
 *
 * Example 4:
 *
 *
 * Input: s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
 * Output: "abcdefghijklmnopqrstuvwxyz"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s[i] only contains digits letters ('0'-'9') and '#' letter.
 * s will be valid string such that mapping is always possible.
 *
 */
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str = `1326#`
	ans := freqAlphabets(str)
	fmt.Printf("%v\n", ans)
}

// @lc code=start

var re = regexp.MustCompile(`\d{2}#|\d`)
var hmap = map[string]byte{
	"1":   byte('a'),
	"2":   byte('b'),
	"3":   byte('c'),
	"4":   byte('d'),
	"5":   byte('e'),
	"6":   byte('f'),
	"7":   byte('g'),
	"8":   byte('h'),
	"9":   byte('i'),
	"10#": byte('j'),
	"11#": byte('k'),
	"12#": byte('l'),
	"13#": byte('m'),
	"14#": byte('n'),
	"15#": byte('o'),
	"16#": byte('p'),
	"17#": byte('q'),
	"18#": byte('r'),
	"19#": byte('s'),
	"20#": byte('t'),
	"21#": byte('u'),
	"22#": byte('v'),
	"23#": byte('w'),
	"24#": byte('x'),
	"25#": byte('y'),
	"26#": byte('z'),
}

func freqAlphabets(s string) string {

	bs := []byte{}

	for _, match := range re.FindAllString(s, -1) {
		bs = append(bs, hmap[match])
	}

	return string(bs)
}

// @lc code=end
