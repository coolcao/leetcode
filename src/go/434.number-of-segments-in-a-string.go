/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 *
 * https://leetcode.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (37.49%)
 * Likes:    210
 * Dislikes: 717
 * Total Accepted:    74.8K
 * Total Submissions: 198.7K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * Count the number of segments in a string, where a segment is defined to be a
 * contiguous sequence of non-space characters.
 *
 * Please note that the string does not contain any non-printable characters.
 *
 * Example:
 *
 * Input: "Hello, my name is John"
 * Output: 5
 *
 *
 */
package main

import "fmt"

// @lc code=start
func isSpace(c byte) bool {
	return c == ' '
}
func countSegments(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	// 记录前一个位置是否是空格
	count := 0
	pre := true
	for i := 0; i < length; i++ {
		current := isSpace(s[i])
		if !pre && current {
			count++
		}
		pre = current
	}
	if !pre {
		count++
	}
	return count
}

// @lc code=end

func main() {
	type TestCase struct {
		Case   string
		Except int
	}
	cases := []*TestCase{
		&TestCase{Case: "", Except: 0},
		&TestCase{Case: " ", Except: 0},
		&TestCase{Case: "  ", Except: 0},
		&TestCase{Case: "        ", Except: 0},
		&TestCase{Case: "a", Except: 1},
		&TestCase{Case: " a", Except: 1},
		&TestCase{Case: " a ", Except: 1},
		&TestCase{Case: "a b", Except: 2},
		&TestCase{Case: " a b", Except: 2},
		&TestCase{Case: " a b ", Except: 2},
		&TestCase{Case: " a         b ", Except: 2},
		&TestCase{Case: "abc", Except: 1},
		&TestCase{Case: "a bc", Except: 2},
		&TestCase{Case: "a b c", Except: 3},
		&TestCase{Case: "a b c ", Except: 3},
		&TestCase{Case: " Hello, My Good Boy", Except: 4},
		&TestCase{Case: " Hello, My Good Boy         ", Except: 4},
		&TestCase{Case: "a	bc", Except: 1},
	}
	for _, tc := range cases {
		r := countSegments(tc.Case)
		if tc.Except != r {
			fmt.Printf("%s want %d, but got %d\n", tc.Case, tc.Except, r)
		}
	}
	fmt.Println("all cases over!")
}
