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
const moveZeroes = function (nums) {
    let sum = 0;
    const length = nums.length;
    const tmp = [];

    nums.reduceRight((pre, current) => {
        if (current === 0) {
            pre.push(current);
        } else {
            pre.unshift(current);
        }
        return pre;
    }, tmp);

    for (let i = 0; i < tmp.length; i++) {
        nums[i] = tmp[i];
    }

};

/**
 * 严格意义上说，上面这个方法是错的，因为额外用的了tmp数组
 * 其实，还有一种更简单的方法，将所有不为0的元素往前移，后面的所有位补0即可
 */

const moveZeroes2 = function (nums) {
    const length = nums.length;
    let sum = 0;
    for (let i = 0; i < length; i++) {
        if (nums[i] === 0) {
            sum += 1;
        } else {
            nums[i - sum] = nums[i];
        }
    }

    for (let i = length - 1; sum > 0; i--) {
        nums[i] = 0;
        sum -= 1;
    }
}

const nums = [0, 1, 1, 2, 0, 3, 4, 0, 20, 0];

moveZeroes2(nums);

console.log(nums);