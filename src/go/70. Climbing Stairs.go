package main

import "fmt"

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

// 直接使用递归的方式进行
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 思想还是递归的思想，只不过使用一个数组将递归之前的数缓存起来，避免重复计算。
func climbStairs2(n int) int {
	tmp := make([]int, n)
	for idx := 0; idx < n; idx++ {
		if idx == 0 || idx == 1 {
			tmp[idx] = idx + 1
			continue
		}
		tmp[idx] = tmp[idx-1] + tmp[idx-2]
	}

	fmt.Printf("%v\n", tmp)
	return tmp[n-1]
}

func main() {
	n := 50
	// steps := climbStairs(n)
	steps2 := climbStairs2(n)
	// fmt.Printf("%d阶台阶总共有%d种走法\n", n, steps)
	fmt.Printf("%d阶台阶总共有%d种走法\n", n, steps2)
}
