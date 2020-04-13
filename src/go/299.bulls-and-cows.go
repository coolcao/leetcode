/*
 * @lc app=leetcode id=299 lang=golang
 *
 * [299] Bulls and Cows
 *
 * https://leetcode.com/problems/bulls-and-cows/description/
 *
 * algorithms
 * Easy (41.67%)
 * Likes:    539
 * Dislikes: 641
 * Total Accepted:    139.2K
 * Total Submissions: 334.1K
 * Testcase Example:  '"1807"\n"7810"'
 *
 * You are playing the following Bulls and Cows game with your friend: You
 * write down a number and ask your friend to guess what the number is. Each
 * time your friend makes a guess, you provide a hint that indicates how many
 * digits in said guess match your secret number exactly in both digit and
 * position (called "bulls") and how many digits match the secret number but
 * locate in the wrong position (called "cows"). Your friend will use
 * successive guesses and hints to eventually derive the secret number.
 *
 * Write a function to return a hint according to the secret number and
 * friend's guess, use A to indicate the bulls and B to indicate the cows.
 *
 * Please note that both secret number and friend's guess may contain duplicate
 * digits.
 *
 * Example 1:
 *
 *
 * Input: secret = "1807", guess = "7810"
 *
 * Output: "1A3B"
 *
 * Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
 *
 * Example 2:
 *
 *
 * Input: secret = "1123", guess = "0111"
 *
 * Output: "1A1B"
 *
 * Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a
 * cow.
 *
 * Note: You may assume that the secret number and your friend's guess only
 * contain digits, and their lengths are always equal.
 */

package main

import "fmt"

func main() {
	secret := "3746737838773737272882833"
	guess := "3747748384774738384877338"
	fmt.Println(getHint(secret, guess))
}

// @lc code=start
func getHint(secret string, guess string) string {
	bulls := 0
	length := len(secret)
	guesses := []byte{}
	m := map[byte]int{}
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			m[secret[i]]++
			guesses = append(guesses, guess[i])
		}
	}
	cows := 0
	for i := 0; i < len(guesses); i++ {
		b := guesses[i]
		if g, ok := m[b]; ok && g > 0 {
			cows++
			m[b]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

// @lc code=end
