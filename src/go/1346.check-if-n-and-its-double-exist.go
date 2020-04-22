/*
 * @lc app=leetcode id=1346 lang=golang
 *
 * [1346] Check If N and Its Double Exist
 *
 * https://leetcode.com/problems/check-if-n-and-its-double-exist/description/
 *
 * algorithms
 * Easy (40.66%)
 * Likes:    111
 * Dislikes: 13
 * Total Accepted:    18.2K
 * Total Submissions: 45K
 * Testcase Example:  '[10,2,5,3]'
 *
 * Given an array arr of integers, check if there exists two integers N and M
 * such that N is the double of M ( i.e. N = 2 * M).
 *
 * More formally check if there exists two indices i and j such that :
 *
 *
 * i != j
 * 0 <= i, j < arr.length
 * arr[i] == 2 * arr[j]
 *
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [10,2,5,3]
 * Output: true
 * Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [7,1,14,11]
 * Output: true
 * Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [3,1,7,11]
 * Output: false
 * Explanation: In this case does not exist N and M, such that N = 2 * M.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= arr.length <= 500
 * -10^3 <= arr[i] <= 10^3
 *
 *
 */

package main

import "fmt"

func main() {
	arr := []int{-95, -299, 357, 715, -438, -759, 255, 418, -647, -805, -546, -182, -523, 13, -79, -227, 537, -655, 993, -526, -518, 679, -420, -53, 120, 187, -203, -567, -75, 464, -472, -324, 16, 215, -773, 862, -563, -839, -906, -969, 633, -990, 756, -17, -346, 820, -216, 736, 188, -412, 881, -599, -181, -673, 802, 688, 553, 323, -689, 625, 871, -938, -969, -207, -703, 794, 361, 111, -884, 156, -223, -480, -734, -838, -53, 335, 720, -379, 855, -971, -928, 99, -876, 75, 721, -736, -913, 911}
	b := checkIfExist2(arr)
	fmt.Println(b)
}

// @lc code=start
func checkIfExist(arr []int) bool {
	positives, negatives := make([]bool, 1001), make([]bool, 1001)
	zero := 0
	for _, v := range arr {
		if v < 0 {
			negatives[-1*v] = true
		} else if v > 0 {
			positives[v] = true
		} else {
			zero++
			if zero > 1 {
				return true
			}
		}
	}
	for i := 1; i <= 500; i++ {
		if (positives[i] && positives[i*2]) || (negatives[i] && negatives[i*2]) {
			return true
		}
	}
	return false
}

// @lc code=end
