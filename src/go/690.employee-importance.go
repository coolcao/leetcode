/*
 * @lc app=leetcode id=690 lang=golang
 *
 * [690] Employee Importance
 *
 * https://leetcode.com/problems/employee-importance/description/
 *
 * algorithms
 * Easy (56.12%)
 * Likes:    594
 * Dislikes: 609
 * Total Accepted:    70.3K
 * Total Submissions: 124.2K
 * Testcase Example:  '[[1,2,[2]], [2,3,[]]]\n2'
 *
 * You are given a data structure of employee information, which includes the
 * employee's unique id, his importance value and his direct subordinates' id.
 *
 * For example, employee 1 is the leader of employee 2, and employee 2 is the
 * leader of employee 3. They have importance value 15, 10 and 5, respectively.
 * Then employee 1 has a data structure like [1, 15, [2]], and employee 2 has
 * [2, 10, [3]], and employee 3 has [3, 5, []]. Note that although employee 3
 * is also a subordinate of employee 1, the relationship is not direct.
 *
 * Now given the employee information of a company, and an employee id, you
 * need to return the total importance value of this employee and all his
 * subordinates.
 *
 * Example 1:
 *
 *
 * Input: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
 * Output: 11
 * Explanation:
 * Employee 1 has importance value 5, and he has two direct subordinates:
 * employee 2 and employee 3. They both have importance value 3. So the total
 * importance value of employee 1 is 5 + 3 + 3 = 11.
 *
 *
 *
 *
 * Note:
 *
 *
 * One employee has at most one direct leader and may have several
 * subordinates.
 * The maximum number of employees won't exceed 2000.
 *
 *
 *
 *
 */

package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// @lc code=start

func getImportance(employees []*Employee, id int) int {
	employeeImportances := map[int]int{}
	employeeMap := map[int]*Employee{}
	for _, e := range employees {
		employeeImportances[e.Id] = e.Importance
		employeeMap[e.Id] = e
	}

	sum := 0

	all := []*Employee{employeeMap[id]}

	for len(all) > 0 {
		e := all[0]
		sum += e.Importance
		all = all[1:]
		for _, subid := range e.Subordinates {
			all = append(all, employeeMap[subid])
		}
	}

	return sum
}

// @lc code=end
