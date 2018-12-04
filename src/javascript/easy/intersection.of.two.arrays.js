'use strict';

/**
 * Given two arrays, write a function to compute their intersection.

    Example:
    Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

    Note:
    Each element in the result must be unique.
    The result can be in any order.

---------
    计算两个数组的交集。注意集合的概念。

 */


/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
const intersection = function(nums1, nums2) {
    const hash = {};
    let result = [];
    const l1 = nums1.length;
    const l2 = nums2.length;
    for(let item of nums1){
        hash[item] = 1;
    }

    for(let item of nums2){
        if(hash[item] && hash[item] === 1){
            result.push(item);
            hash[item] += 1;
        }
    }

    return result;

};

const intersection2 = function (nums1,nums2) {
    const s1 = new Set(nums1);
    const s2 = new Set(nums2);
    return [...s1].filter(num => {
        return s2.has(num);
    });
}

const nums1 = [1,2,2,1];
const nums2 = [2,3,4];

console.log(intersection2(nums1,nums2));


