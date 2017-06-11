'use strict';

/**
 * Given an array of integers, find if the array contains any duplicates.
 * Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
 */

/**
 * @param {number[]} nums
 * @return {boolean}
 */
const containsDuplicate = function(nums) {
    const hash = {};
    for (let item of nums) {
        if (hash[item]) {
            return true;
        } else {
            hash[item] = true;
        }
    }
    return false;
};

const nums = [1, 2, 3, 4, 5];

console.log(containsDuplicate(nums));
