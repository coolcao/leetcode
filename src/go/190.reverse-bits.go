/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
Reverse bits of a given 32 bits unsigned integer.



Example 1:

Input: 00000010100101000001111010011100
Output: 00111001011110000010100101000000
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
Example 2:

Input: 11111111111111111111111111111101
Output: 10111111111111111111111111111111
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.


Note:

Note that in some languages such as Java, there is no unsigned integer type. In this case, both input and output will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above the input represents the signed integer -3 and the output represents the signed integer -1073741825.


Follow up:

If this function is called many times, how would you optimize it?
*/
package main

import (
	"fmt"
	"math"
)

func uint32ToArray(num uint32) [32]uint32 {
	var result [32]uint32
	idx := 31
	for num >= 2 {
		result[idx] = num % 2
		num /= 2
		idx--
	}
	result[idx] = num
	return result
}

func arryTouint32(array [32]uint32) uint32 {
	var num uint32
	for i := 0; i < 32; i++ {
		num += array[i] * (uint32)(math.Pow(2, float64(i)))
	}
	return num
}

// @lc code=start
func reverseBits(num uint32) uint32 {
	arr := uint32ToArray(num)
	return arryTouint32(arr)
}

// 使用位运算
// 这个方法很巧妙，依次取得num最右边的一位，放到ret，然后再将ret向左移动一位。
// 这样就达到了ret最后相对于num是二进制位反转过来的。
func reverseBits2(num uint32) uint32 {
	var ret uint32
	ret = 0
	var i uint32
	for i = 0; i < 32; i++ {
		ret = ret << 1
		ret = ret | ((num >> i) & 1)
	}
	return ret
}

func main() {
	var num uint32
	num = 43261596
	rnum := reverseBits(num)
	rnum2 := reverseBits2(num)
	fmt.Printf("%v\n", rnum)
	fmt.Printf("%v\n", rnum2)
}

// @lc code=end
