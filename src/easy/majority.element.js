'use strict';

/**
 * Given an array of size n, find the majority element. 
 * The majority element is the element that appears more than ⌊ n/2 ⌋ times.
 * You may assume that the array is non-empty and the majority element always exist in the array.
 */

/**
 * @param {number[]} nums
 * @return {number}
 */
const majorityElement = function (nums) {
    const hash = {};
    for (let item of nums) {
        if (hash[item]) {
            hash[item].count += 1;
        } else {
            hash[item] = {
                count: 1,
                val: item
            }
        }
    }
    let result = nums[0];
    for (let key in hash) {
        if (hash[key].count >= nums.length / 2) {
            result = hash[key].val;
        }
    }
    console.log(hash);
    return result;
};

// 由于majority element个数大于 n/2 ，那么，将数组排序后，处于 n/2 位置上的元素必然是 majority element。
// 蛋这种效率也不高，因为要做排序，复杂度是 O(n^2)
const majorityElement2 = function (nums) {
    nums.sort();
    return nums[nums.length / 2 >>> 0];
}

const majorityElement3 = function (nums) {
    let mj = 0;
    let cnt = 1;
    const length = nums.length;
    for (let i = 1; i < length; i++) {
        if (nums[i] === nums[mj]) {
            cnt++;
        } else {
            cnt--;
        }

        if (cnt === 0) {
            mj = i;
            cnt = 1;
        }
    }

    return nums[mj];

};

const nums = [1,2,1,1,3,1,4,1,5,1,6,7,7,7,7,1,8,1,9,1,10,1,11,1,12,1,13,1,1,1,1];
console.log(nums.length);
console.log(majorityElement(nums));