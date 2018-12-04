'use strict';

/**
 * Given two arrays, write a function to compute their intersection.

    Example:
    Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

    Note:
    Each element in the result should appear as many times as it shows in both arrays.
    The result can be in any order.
    Follow up:
    What if the given array is already sorted? How would you optimize your algorithm?
    What if nums1's size is small compared to nums2's size? Which algorithm is better?
    What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
 */

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
const intersect = function(nums1, nums2) {
    const l1 = nums1.length;
    const l2 = nums2.length;
    const l = l1 > l2 ? l1 : l2;
    let m1 = new Map();
    let m2 = new Map();
    for (let i = 0; i < l; i++) {
        const item1 = nums1[i];
        const item2 = nums2[i];
        if(item1!==null && item1!==undefined){
            if(m1.has(item1)){
                m1.set(item1,m1.get(item1) + 1);
            }else{
                m1.set(item1,1);
            }
        }
        if(item2!==null && item2!==undefined){
            if(m2.has(item2)){
                m2.set(item2,m2.get(item2) + 1);
            }else{
                m2.set(item2,1);
            }
        }
    }
    const result = [];
    for(let [k,v] of m1.entries()){
        if(m2.has(k)){
            const tmplength = m2.get(k)>v?v:m2.get(k);
            for(let i=0;i<tmplength;i++){
                result.push(k);
            }
        }
    }
    return result;
};

const intersect2 = function (nums1,nums2) {

    const m1 = {};
    for(const item of nums1){
        if(m1[item]){
            m1[item] += 1;
        }else{
            m1[item] = 1;
        }
    }
    const result = [];
    for(const item of nums2){
        if(m1[item]){
            m1[item] -= 1;
            result.push(item);
        }
    }
    return result;

}



console.log(intersect2([8,0,0,1],[0,0,2]));