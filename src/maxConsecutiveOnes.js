'use strict';

/**
 * Given a binary array, find the maximum number of consecutive 1s in this array.
 * Example 1:
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive 1s.
 *     The maximum number of consecutive 1s is 3.
 * Note:
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 */

/**
 * @param {number[]} nums
 * @return {number}
 */
const findMaxConsecutiveOnes = function(nums) {
    let result = [];
    let length = nums.length;
    let max = 0;
    for(let i=0;i<length;i++){
        if(nums[i] === 1){
            max ++;
        }else if(nums[i] === 0  ){
            result.push(max);
            max = 0;
        }
        if(i === length - 1){
            result.push(max);
        }
    }

    result.forEach(item => {
        if(item > max){
            max = item;
        }
    });
    return max;
};

const findMaxConsecutiveOnes2 = function (nums) {
    let length = nums.length;
    let max  = 0;
    let len = 0;
    for(let i=0;i<length;i++){
        if(nums[i] === 1){
            len ++;
        }else if(nums[i] === 0  ){
            if(len > max){
                max = len;
            }
            len = 0;
        }
        if(i === length - 1){
            if(len > max){
                max = len;
            }
        }
    }
    return max;
}

console.log(findMaxConsecutiveOnes2([1,0,1,1,0,1]));