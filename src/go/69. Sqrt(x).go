package main

import "fmt"

/*
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	for i := 1; ; i++ {
		if (i+1)*(i+1) > x {
			return i
		}
	}
}

func mySqrt2(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	min, max := 1, x
	mid := (min + max) / 2
	midmid := mid * mid

	for midmid != x && min < max {
		// fmt.Printf("min:%d \t mid:%d\t max:%d\n", min, mid, max)
		if midmid < x {
			min = mid
		} else if midmid > x {
			max = mid
		}
		mid = (min + max) / 2
		midmid = mid * mid
		if mid == min || mid == max {
			return mid
		}
	}
	return mid

}

func main() {
	// x := 6
	// st := mySqrt(x)
	// st2 := mySqrt2(x)
	// fmt.Printf("%d\n", st)
	// fmt.Printf("%d\n", st2)

	for i := 1; i < 1000000; i++ {
		if mySqrt(i) != mySqrt2(i) {
			fmt.Printf("不相等的：%d, mySqrt:%d mySqrt2: %d\n", i, mySqrt(i), mySqrt2(i))
		}
	}
	fmt.Println("ok")

}
