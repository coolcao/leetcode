'use strict';

/**
 * Given scores of N athletes, find their relative ranks and the people with the top three highest scores, who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

    Example 1:
    Input: [5, 4, 3, 2, 1]
    Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
    Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal". 
    For the left two athletes, you just need to output their relative ranks according to their scores.
    Note:
    N is a positive integer and won't exceed 10,000.
    All the scores of athletes are guaranteed to be unique.
 */



/**
 * @param {number[]} nums
 * @return {string[]}
 */
const findRelativeRanks = function(nums) {
    const tmp = [];
    for (let item of nums) {
        tmp[item] = item;
    }

    const hash = {};

    let index = 1;
    const result = [];
    const length = tmp.length;
    for (let i = length - 1; i >= 0; i--) {
        if (tmp[i] !== undefined) {
            hash[tmp[i]] = '' + index;
            index += 1;
        }
    }
    console.log(hash);
    for(let item of nums){
        if(hash[item] === '1'){
            result.push('Gold Medal');
        }else if(hash[item] === '2'){
            result.push('Silver Medal');
        }else if(hash[item] === '3'){
            result.push('Bronze Medal');
        }else{
            result.push(hash[item]);
        }
    }

    return result;
};


const nums = [123123,11921,1,0,123];

console.log(findRelativeRanks(nums));

