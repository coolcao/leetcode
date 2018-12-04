'use strict'

/**
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 1:

Input: [1,3,5,6], 0
Output: 0

 */

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
const searchInsert = function(nums, target) {
    let idx = -1;
    const length = nums.length;
    for(let i=0;i<length;i++) {
        if (nums[i] === target) {
            idx = i;
            break;
        } else if (nums[i] > target) {
            idx = i;
            break;
        }
    }
    if (idx === -1) idx = length;
    return idx;
};


/**
 * 
 * @param {number[]} nums 
 * @param {number} target 
 */
const searchInsert2 = function(nums, target) {
    const length = nums.length;
    let start = 0, end = length;
    while (start <= end) {
        let mid = ((start + end) / 2) >>> 0;
        if (nums[mid] < target) {
            start = mid + 1;
        } else if (nums[mid] > target) {
            end = mid - 1;
        } else {
            return mid;
        }
    }
    return start;
};

const nums = [1,3,5,6];

console.log(searchInsert2(nums, 0));