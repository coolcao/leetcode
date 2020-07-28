/*
 * @lc app=leetcode id=1122 lang=golang
 *
 * [1122] Relative Sort Array
 *
 * https://leetcode.com/problems/relative-sort-array/description/
 *
 * algorithms
 * Easy (67.28%)
 * Likes:    689
 * Dislikes: 47
 * Total Accepted:    59.2K
 * Total Submissions: 87.5K
 * Testcase Example:  '[2,3,1,3,2,4,6,7,9,2,19]\n[2,1,4,3,9,6]'
 *
 * Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all
 * elements in arr2 are also in arr1.
 *
 * Sort the elements of arr1 such that the relative ordering of items in arr1
 * are the same as in arr2.  Elements that don't appear in arr2 should be
 * placed at the end of arr1 in ascending order.
 *
 *
 * Example 1:
 * Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
 * Output: [2,2,2,1,4,3,3,9,6,7,19]
 *
 *
 * Constraints:
 *
 *
 * arr1.length, arr2.length <= 1000
 * 0 <= arr1[i], arr2[i] <= 1000
 * Each arr2[i] is distinct.
 * Each arr2[i] is in arr1.
 *
 *
 */
package main

import "fmt"

func main() {
	arr1, arr2 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}
	result := relativeSortArray(arr1, arr2)
	fmt.Printf("%v\n", result)
}

// @lc code=start
// var hmap map[int]int

// type SortByArr2 []int

// func (a SortByArr2) Len() int      { return len(a) }
// func (a SortByArr2) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// func (a SortByArr2) Less(i, j int) bool {
// 	ivalue, iok := hmap[a[i]]
// 	jvalue, jok := hmap[a[j]]
// 	if iok && jok {
// 		return ivalue < jvalue
// 	}
// 	if iok {
// 		return true
// 	}
// 	if jok {
// 		return false
// 	}
// 	return a[i] < a[j]
// }

// func relativeSortArray(arr1 []int, arr2 []int) []int {
// 	hmap = map[int]int{}
// 	for i, a2 := range arr2 {
// 		hmap[a2] = i + 1
// 	}
// 	sort.Sort(SortByArr2(arr1))
// 	return arr1
// }

func relativeSortArray(arr1 []int, arr2 []int) []int {
	count := [1001]int{}
	result := []int{}
	for _, num := range arr1 {
		count[num]++
	}
	for _, num := range arr2 {
		for count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}
	for i, c := range count {
		for c > 0 {
			result = append(result, i)
			c--
		}
	}
	return result
}

// @lc code=end
