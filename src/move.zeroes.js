'use strict';

/**
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

    For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].

    Note:
    You must do this in-place without making a copy of the array.
    Minimize the total number of operations.

 */

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
const moveZeroes = function(nums) {
    let sum = 0;
    const length = nums.length;

    const swap = function(a, b) {
        let tmp = nums[a];
        nums[a] = nums[b];
        nums[b] = tmp;
    }

    for (let i = 0; i < length; i++) {

        if (i + sum === length) {
            return;
        }
        if (nums[i] === 0) {
            sum += 1;
            while (nums[length - sum] === 0) {
                sum += 1;
            }
            swap(i, length - sum);
        }
    }
};

const array = [0, 1, 2, 0, 1, 9, 0, 8, 7, 4, 0, 0, 0, 0];

moveZeroes(array);

console.log(array);