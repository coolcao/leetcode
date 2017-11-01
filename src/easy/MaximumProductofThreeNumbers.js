'use strict';

/**
Given an integer array, find three numbers whose product is maximum and output the maximum product.

Example 1:
Input: [1,2,3]
Output: 6
Example 2:
Input: [1,2,3,4]
Output: 24
Note:
The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.

 */

/**
 * @param {number[]} nums
 * @return {number}
 */
const maximumProduct = function(nums) {
    const length = nums.length;
    let max = -1000*-1000*-1000;
    let positives = 0, negatives = 0;
    let max1 = -1001, max2 = -1001,max3 = -1001;
    let min1 = 1001, min2 = 1001,min3 = 1001;
    for(let num of nums) {
        if (num > 0) positives += 1;
        else if(num<0) negatives += 1;

        if (num > max1) {
            max3 = max2;
            max2 = max1;
            max1 = num;
        } else if (num > max2) {
            max3 = max2;
            max2 = num;
        } else if (num > max3) {
            max3 = num;
        }
        if (num < min1) {
            min3 = min2;
            min2 = min1;
            min1 = num;
        } else if (num < min2) {
            min3 = min2;
            min2 = num;
        } else if (num < min3) {
            min3 = num;
        }
    }
    const tmp = [max1,max2,max3];
    if (length===4) tmp.push(min1);
    if (length===5) tmp.push(min1,min2);
    if (length>=6) tmp.push(min1,min2,min3);
    for(let x=0;x<tmp.length-2;x++) {
        for(let y=x+1;y<tmp.length-1;y++){
            for(let z=y+1;z<tmp.length;z++){
                if (x!==y && x!==z) {
                    const product = tmp[x]*tmp[y]*tmp[z];
                    if (max < product) {
                        max = product;
                    }
                }
            }
        }
    }
    // console.log(positives + '|' + negatives);
    // console.log(max1 + '|' + max2 + '|' + max3);
    // console.log(min1 + '|' + min2 + '|' + min3);
    // console.log(tmp);
    return max;
};

const nums = [3,4,2,1,5,5,3,12,4,2,1,3,32,32,4,54,-1,3,-2,43,-43,-5,3,-42,0];
// const nums = [-1,3,-2,0,1];

console.log(maximumProduct(nums));