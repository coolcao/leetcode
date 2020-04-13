/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 *
 * https://leetcode.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (43.31%)
 * Likes:    574
 * Dislikes: 1014
 * Total Accepted:    202.5K
 * Total Submissions: 467.4K
 * Testcase Example:  '"hello"'
 *
 * Write a function that takes a string as input and reverse only the vowels of
 * a string.
 *
 * Example 1:
 *
 *
 * Input: "hello"
 * Output: "holle"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "leetcode"
 * Output: "leotcede"
 *
 *
 * Note:
 * The vowels does not include the letter "y".
 *
 *
 *
 */

// @lc code=start
package main

import "fmt"

func main() {
	s := "skfjieuwoekfjksjfiwejowsklaskdfjiejfsdlkfjei"
	if reverseVowels(s) != reverseVowels2(s) {
		fmt.Println("Error")
		fmt.Println(reverseVowels2(s))
	} else {
		fmt.Println("ok")
	}
}

// --------- 先把元音字母拿出来，然后倒序再插进去---------------
func reverseVowels(s string) string {
	vowelsMap := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	length := len(s)
	if length == 0 {
		return ""
	}
	vowels := []byte{}
	for i := 0; i < length; i++ {
		b := s[i]
		if _, ok := vowelsMap[b]; ok {
			vowels = append(vowels, b)
		}
	}
	vc := 0
	bytes := make([]byte, length)
	vl := len(vowels)
	for i := 0; i < length; i++ {
		b := s[i]
		if _, ok := vowelsMap[b]; ok {
			bytes[i] = vowels[vl-1-vc]
			vc++
		} else {
			bytes[i] = b
		}
	}
	return string(bytes)
}

// -----------两个指针，一前一后，遇到元音交换----------
func reverseVowels2(s string) string {
	vowelsMap := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	length := len(s)
	front, back := 0, length-1
	bytes := []byte(s)
	for front < back {
		fb, bb := bytes[front], bytes[back]
		_, fok := vowelsMap[fb]
		_, bok := vowelsMap[bb]
		if fok && bok {
			bytes[front], bytes[back] = bytes[back], bytes[front]
			front++
			back--
		} else if fok {
			back--
		} else if bok {
			front++
		} else {
			front++
			back--
		}
	}
	return string(bytes)
}

// @lc code=end
