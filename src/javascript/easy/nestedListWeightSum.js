'use strict';

/**
 * Given a nested list of integers, return the sum of all integers in the list weighted by their depth.
 * Each element is either an integer, or a list -- whose elements may also be integers or other lists.
 * Example 1:
 * Given the list [[1,1],2,[1,1]], return 10. (four 1's at depth 2, one 2 at depth 1)
 * Example 2:
 * Given the list [1,[4,[6]]], return 27. (one 1 at depth 1, one 4 at depth 2, and one 6 at depth 3; 1 + 4*2 + 6*3 = 27)
 */

 let depthSum = function (nums) {
    let sum = 0;
    let map = new Map();
    let q = [];
    nums.forEach(item => {
        q.push({depth:1,value:item});
    });
    while (q.length > 0) {
        let item = q.pop();
        if(typeof item.value == 'number'){
            sum += (item.value * item.depth);
        }else{
            item.value.forEach(_item => {
                q.push({depth:item.depth+1,value:_item});
            });
        }
    }
    return sum;
 }

 console.log(depthSum([1,[2,3],[4,5,[6,7],[8],[9]]]));