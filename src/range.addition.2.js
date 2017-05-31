'use strict'

/**
 * Given an m * n matrix M initialized with all 0's and several update operations.
    Operations are represented by a 2D array, and each operation is represented by an array with two positive integers a and b, which means M[i][j] should be added by one for all 0 <= i < a and 0 <= j < b.
    You need to count and return the number of maximum integers in the matrix after performing all the operations.

    Example 1:
    Input: 
    m = 3, n = 3
    operations = [[2,2],[3,3]]
    Output: 4
    Explanation: 
    Initially, M = 
    [[0, 0, 0],
     [0, 0, 0],
     [0, 0, 0]]

    After performing [2,2], M = 
    [[1, 1, 0],
     [1, 1, 0],
     [0, 0, 0]]

    After performing [3,3], M = 
    [[2, 2, 1],
     [2, 2, 1],
     [1, 1, 1]]

    So the maximum integer in M is 2, and there are four of it in M. So return 4.
    Note:
    The range of m and n is [1,40000].
    The range of a is [1,m], and the range of b is [1,n].
    The range of operations size won't exceed 10,000.
 */

/**
 * @param {number} m
 * @param {number} n
 * @param {number[][]} ops
 * @return {number}
 */
const maxCount = function(m, n, ops) {
    const arr = [];
    for(let i=0;i<m;i++){
        arr[i] = [];
        for(let j=0;j<n;j++){
            arr[i][j] = 0;
        }
    }
    for(let op of ops){
        for(let i=0;i<op[0];i++){
            for(let j=0;j<op[1];j++){
                arr[i][j] += 1;
            }
        }
    }
    console.log(JSON.stringify(arr));
    let max = arr[0][0];
    let sum = 0;
    for(let i=0;i<m;i++){
        for(let j=0;j<n;j++){
            if(max < arr[i][j]){
                sum = 0;
                max = arr[i][j];
            }else if(max === arr[i][j]){
                sum += 1;
            }
        }
    }
    return sum;
};

const maxCount2 = function (m,n,ops) {
    if(!ops[0]){
        return m * n;
    }
    let minR = ops[0][0];
    let minC = ops[0][1];
    for(let item of ops){
        if(minR > item[0]){
            minR = item[0];
        }
        if(minC > item[1]){
            minC = item[1];
        }
    }
    return minR * minC;
}


console.log(maxCount2(5,5,[[2,3],[4,2],[1,5]]));


