package main

import (
	"fmt"
	"strconv"
)

/*

The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.



Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"
*/

func sayCount(str string) string {
	// length := len(str)
	var current rune
	count := 0
	result := ""
	for _, val := range str {
		if current == val {
			count++
		} else {
			if count > 0 {
				result += strconv.Itoa(count) + string(current)
			}
			count = 1
			current = val
		}
	}
	if count > 0 {
		result += strconv.Itoa(count) + string(current)
	}
	return result
}

func countAndSay(n int) string {
	results := make([]string, n)
	results[0] = "1"
	for i := 1; i < n; i++ {
		results[i] = sayCount(results[i-1])
	}
	fmt.Printf("%v\n", results)
	return results[n-1]
}

func main() {
	n := 10
	result := countAndSay(n)
	fmt.Printf("result : %s\n", result)
}
