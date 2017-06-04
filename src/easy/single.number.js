'use strict';

/**
 * Given an array of integers, every element appears twice except for one. Find that single one.
 *  Note:
 *  Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 *  Subscribe to see which companies asked this question.
 */

/**
 * [singleNumber description]
 * @param  {number[]} nums 
 * @return {number}      
 */
const singleNumber1 = function(nums) {
    const set = new Set();
    for(const num of nums){
        if(set.has(num)){
            set.delete(num);
        }else{
            set.add(num);
        }
    }
    return set.values().next().value;
};

const singleNumber2 = function (nums) {
    const obj = {};
    for(const num of nums){
        if(obj[num]){
            obj[num] ++
        }else{
            obj[num] = 1;
        }
    }
    for(const key of Object.getOwnPropertyNames(obj)){
        if(obj[key] === 1){
            return Number(key);
        }
    }
}

const singleNumber3 = function (nums) {
    let result = 0;
    for(const num of nums){
        result = result ^ num;
    }
    return result;
}

const array = [1,1,1,1,2,2,3,4,4,5,5,6,6,7,7,8,8,9,9,0,0,1,1];

console.time('set');
console.log(singleNumber1(array));
console.timeEnd('set');

console.time('obj');
console.log(singleNumber2(array));
console.timeEnd('obj');

console.time('xor');
console.log(singleNumber3(array));
console.timeEnd('xor');
