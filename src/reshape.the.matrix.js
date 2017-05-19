'use strict';

/**
 * @param {number[][]} nums
 * @param {number} r
 * @param {number} c
 * @return {number[][]}
 */
const matrixReshape = function(nums, r, c) {
    // 首先扁平化数组
    const tmp = nums.reduce((pre,current) => {
        pre = pre.concat(current);
        return pre;
    },[]);
    // 如果传入的行列有误，直接返回原数组
    const length = tmp.length;
    if(length !== r*c){
        return nums;
    }

    // 重新格式化数组
    const result = [];
    for(let i=0;i<r;i++){
        const item = [];
        for(let j=i*c,l=0;l<c;j++,l++){
            item.push(tmp[j]);
        }
        result.push(item);
    }
    return result;

};

console.log(matrixReshape([[2,3],[4,5],[6,8]],2,4));