/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (41.23%)
 * Likes:    431
 * Dislikes: 190
 * Total Accepted:    140.7K
 * Total Submissions: 341.2K
 * Testcase Example:  '16'
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a
 * power of 4.
 * 
 * Example 1:
 * 
 * 
 * Input: 16
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 5
 * Output: false
 * 
 * 
 * Follow up: Could you solve it without loops/recursion?
 */

// @lc code=start
// 判断一个整数是否是4的幂
// 这里可以和判断一个数是否是2的幂类似，使用二进制操作进行
// 不同的是，观察4的幂的二进制，只有奇数位上的数字是1，而2的幂是所有位都可能是1
// 就用这个不同来做判断
func isPowerOfFour(num int) bool {
	if num <=0 {
		return false
	}
	//0x55555555 的二进制表示为 1010101010101010101010101010101
	return num&(num-1) == 0 && (num&0x55555555)==num
}
// @lc code=end

