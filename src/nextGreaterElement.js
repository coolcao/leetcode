'use strict';

/**
 * You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2. Find all the next greater numbers for nums1's elements in the corresponding places of nums2.
 * The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.
 * Example 1:
 * Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
 * Output: [-1,3,-1]
 * Explanation:
 *     For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
 *     For number 1 in the first array, the next greater number for it in the second array is 3.
 *     For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
 * Example 2:
 * Input: nums1 = [2,4], nums2 = [1,2,3,4].
 * Output: [3,-1]
 * Explanation:
 *     For number 2 in the first array, the next greater number for it in the second array is 3.
 *     For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
 * Note:
 * All elements in nums1 and nums2 are unique.
 * The length of both nums1 and nums2 would not exceed 1000.
 */

/**
 * 给定两个数组（无重复）nums1 与 nums2，其中nums1的元素是nums2的子集。
 * 在nums2中寻找大于nums1中对应位置且距离最近的元素。如果对应位置不存在这样的元素，则输出-1。
 * 注意：
 * nums1与nums2中的所有元素都是唯一的。
 * nums1与nums2的元素个数不超过1000。
 */

/**
 * @param {number[]} findNums
 * @param {number[]} nums
 * @return {number[]}
 */

const nextGreaterElement = function (findNums,nums) {
    let length = nums.length;
    for(let i=0;i<findNums.length;i++){
        let index = nums.indexOf(findNums[i]);
        for(let j=index;j<length;j++){
            if(nums[j]>findNums[i]){
                findNums[i] = nums[j];
                break;
            }else if(j == length - 1){
                findNums[i] = -1;
            }
        }
    }
    return findNums;
}

console.log(nextGreaterElement([4,1,2],[1,3,4,2]));