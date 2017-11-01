'use strict';
/*
Given an unsorted array of integers, find the length of longest continuous increasing subsequence.

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.
Note: Length of the array will not exceed 10,000.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
const findLengthOfLCIS = function(nums) {
    const length = nums.length;
    if (length === 0) {
        return 0;
    }
    let max = 1;
    let sub_length = 1;
    let start = 0;
    for (let i = 0; i < length-1; i++) {
        if (nums[i+1]-nums[i] > 0) {
            sub_length += 1;
            // start += 1;
        } else {
            sub_length = 1;
            start = i + 1;
        }
        if (max < sub_length) {
            max = sub_length;
        }
    }
    return max;
};

const nums = [1, 3, 5, 7, 1, 9, 11, 13, 15, 17, 19,0,1,59,32,1,4,2,1,4,5,7,8,9,11,16,23,26,34,36,38,40,47,58,60];
console.log(findLengthOfLCIS(nums));
