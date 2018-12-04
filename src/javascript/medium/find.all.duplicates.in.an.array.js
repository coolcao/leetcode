'use strict';

/**
 * Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

    Find all the elements that appear twice in this array.

    Could you do it without extra space and in O(n) runtime?

    Example:
    Input:
    [4,3,2,7,8,2,3,1]

    Output:
    [2,3]
 */



/**
 * 正负位标记法
 * @param {number[]} nums
 * @return {number[]}
 */
const findDuplicates = function(nums) {
    const result = [];
    for(let i=0;i<nums.length;i++){
        let index = nums[i];
        if(nums[Math.abs(index) - 1] < 0){
            result.push(Math.abs(index));
        }else{
            nums[Math.abs(index) - 1] *= -1;
        }
    }
    return result;
};

const findDuplicates2 = function (nums) {
    const result = [];
    const length = nums.length;
    for(let i=0;i<length;i++){
        while (nums[i] && nums[i] !== (i+1)) {
            let n = nums[i];
            if(nums[i] == nums[n - 1]){
                result.push(n);
                nums[i] = 0;
            }else{
                nums[i] = nums[n-1];
                nums[n-1] = n;
            }
        }
    }
    console.log(nums);
    return result;
}
const nums = [4,3,2,7,8,2,3,1];
console.log(findDuplicates2(nums));



