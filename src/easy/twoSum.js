/**
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * ====================
 * ------Easy----------
 * ====================
 */


let twoSum = function(nums, target) {
    let result = [];
    let map = new Map();
    for (let i = 0; i < nums.length; i++) {
        let value = nums[i];
        let an = target - value;
        let diff = map.get(value);
        if(diff){
            if(diff.diff === value){
                result.push(diff.index,i);
                break;
            }
        }else{
            map.set(an,{index:i,diff:an});
        }
    }
    return result;
}

let nums = [2, 7, 11, 15];
console.log(twoSum(nums, 26));