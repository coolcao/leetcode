'use strict';

/**
 * Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

    Example:

    Input:
    [1,2,3]

    Output:
    3

    Explanation:
    Only three moves are needed (remember each move increments two elements):

    [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

 */

/**
 * @param {number[]} nums
 * @return {number}
 */
const minMoves = function(nums) {
    const length = nums.length;
    const isAllEqual = function (array) {
        let data = array[0];
        for(let i=1;i<length;i++){
            if(array[i] !== data){
                return false;
            }
        }
        return true;
    }

    let sum = 0;

    const increments = function (array) {
        sum += 1;

        let max = array[0];
        let midx = 0;

        for(let i=0;i<length;i++){
            if(max <= array[i]){
                max = array[i];
                midx = i;
            }
            array[i] += 1;
        }
        array[midx] -= 1;

    }

    while (!isAllEqual(nums)) {
        increments(nums);
    }

    return sum;
};

const minMoves2 = function (nums) {
    const length = nums.length;
    const findMin = function () {
        let min = nums[0];
        for(let i=1;i<length;i++){
            const item = nums[i];
            if(min > item){
                min = item;
            }
        }
        return min;
    }

    let min = findMin();
    let result = 0;

    for(let i=0;i<length;i++){
        result += (nums[i]-min);
    }
    return result;
}

const nums = [1,2147483647];

console.log(minMoves2(nums));