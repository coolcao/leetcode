/**
 * Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.
 *
 * Example 1:
 * Input: [1,4,3,2]
 * 
 * Output: 4
 * Explanation: n is 2, and the maximum sum of pairs is 4.
 * Note:
 * n is a positive integer, which is in the range of [1, 10000].
 * All the integers in the array will be in the range of [-10000, 10000].
 */

/**
 * 分组的模式很简单，就是先将原数组排序，然后按照排序的进行分组即可
 * 因此，最简单的方法就是，先排序，然后取偶数位上的数字进行加和即可
 */


/**
 * @param {number[]} nums
 * @return {number}
 */
const arrayPairSum = function(nums) {
    // 先排序
    nums = nums.sort((a,b) => {
        return a-b;
    });
    let sum = 0;
    let length = nums.length;
    // 偶数位加和
    for(let i=0;i<length;i=i+2){
        sum += nums[i];
    }
    return sum;
};

const array = [6214, -2290, 2833, -7908];

console.log(arrayPairSum(array));