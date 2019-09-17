package main

import "fmt"

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

func addBinary(a string, b string) string {
	result := make([]uint8, 0)
	la, lb := len(a), len(b)
	length := la
	if la < lb {
		length = lb
	}

	var flag uint8
	flag = 0
	for i := 0; i < length; i++ {
		if la-i > 0 && lb-i > 0 {
			result = append(result, a[la-1-i]-48+b[lb-1-i]-48+flag)
		} else if la-i > 0 {
			result = append(result, a[la-1-i]-48+flag)
		} else if lb-i > 0 {
			result = append(result, b[lb-1-i]-48+flag)
		}
		flag = result[i] / 2
		if result[i] > 1 {
			result[i] = result[i] % 2
		}
	}
	if flag > 0 {
		result = append(result, 1)
	}

	res := []byte{}
	for i := len(result) - 1; i >= 0; i-- {
		res = append(res, byte(result[i]+48))
	}
	return string(res)
}

func main() {
	a, b := "11", "1"
	res := addBinary(a, b)
	fmt.Printf("%s\n", res)
}
