'use strict';

/**
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

    Find all the elements of [1, n] inclusive that do not appear in this array.

    Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

    Example:

    Input:
    [4,3,2,7,8,2,3,1]

    Output:
    [5,6]
 */

const findDisappearedNumbers = function(nums) {
    const result = [];
    let sum = 0;
    const length = nums.length;
    for(let i=0;i<length;i++){
        result[nums[i]-1] = nums[i];
    }
    for(let i=0;i<length;i++){
        if(result[i]){
            continue;
        }else{
            result[sum] = i+1;
            sum += 1;
        }
    }
    result.length = sum;
    return result;
};

const findDisappearedNumbers2 = function (nums) {
    const length = nums.length;
    for(let i=0;i<length;i++){
        let index = Math.abs(nums[i]) - 1;
        if(nums[index]>0){
            nums[index] = -nums[index];
        }
    }
    const result = [];
    for(let i=0;i<length;i++){
        if(nums[i]>0){
            result.push(i+1);
        }
    }
    return result;
}

console.log(findDisappearedNumbers2([1,6,2,2,3,4,3,2]));