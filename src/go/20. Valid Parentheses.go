package main

import "fmt"

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

func isValid(s string) bool {

	sLen := len(s)
	if sLen == 0 {
		return true
	}
	if sLen%2 == 1 {
		return false
	}

	symbolMap := map[byte]int{
		'(': 1,
		')': -1,
		'{': 2,
		'}': -2,
		'[': 3,
		']': -3,
	}

	arr := make([]byte, sLen)

	// 标记栈顶
	p := 0

	for i := 0; i < sLen; i++ {
		c := s[i]
		sm := symbolMap[c]
		if i == 0 {
			arr[p] = c
			p++
			continue
		}
		if p == 0 {
			arr[p] = c
			p++
			continue
		}
		smp := symbolMap[arr[p-1]]
		if (sm + smp) == 0 {
			p--
			arr[p] = 0
			continue
		}
		arr[p] = c
		p++
	}
	if arr[0] == 0 {
		return true
	}

	return false

}

func main() {
	s := "()[]{}"
	valid := isValid(s)
	fmt.Printf("%t\n", valid)
}
