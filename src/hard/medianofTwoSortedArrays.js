'use strict';

/**
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 * Example 1:
 * nums1 = [1, 3]
 * nums2 = [2]
 * The median is 2.0
 * Example 2:
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * The median is (2 + 3)/2 = 2.5
 * ====================
 * ------Hard----------
 * ====================
 */

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
const findMedianSortedArrays = function(nums1, nums2) {
    let i = 0;
    let m = nums1.length;
    let n = nums2.length;
    let length = m + n;
    let index = (length / 2) >>> 0;
    if (index == 0) {
        return nums1[0] || nums2[0] || null;
    }
    let tmp = [];
    while (nums1.length > 0 && nums2.length > 0) {
        i++;
        if(tmp.length <= index){
            if (nums1[0] < nums2[0]) {
                tmp.push(nums1.shift());
            } else {
                tmp.push(nums2.shift());
            }
        }else{
            break;
        }
    }
    while (nums1.length > 0) {
        i++;
        if(tmp.length <= index){
            tmp.push(nums1.shift());
        }else{
            break;
        }
    }
    while (nums2.length > 0) {
        i++;
        if(tmp.length <= index){
            tmp.push(nums2.shift());
        }else{
            break;
        }
    }
    console.log('i = ' + i);
    if(length % 2 === 0){
        return (tmp.pop() + tmp.pop())/2;
    }else{
        return tmp.pop();
    }
};


console.log(findMedianSortedArrays([1,3,5,7,9], [2]));