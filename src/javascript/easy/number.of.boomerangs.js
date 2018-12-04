'use strict';

/**
Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:
Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]

 */

/**
 * @param {number[][]} points
 * @return {number}
 */
const numberOfBoomerangs = function(points) {
    const length = points.length;
    let count = 0;
    for(let i=0;i<length;i++) {
        const lengths = {};
        for(let j=0;j<length;j++) {
            if(i !== j) {
                const x = points[i];
                const y = points[j];
                const distance = Math.pow(x[0]-y[0],2)+Math.pow(x[1]-y[1],2);
                if (lengths[distance]) count += lengths[distance] * 2;
                if (lengths[distance]) lengths[distance] += 1 
                else lengths[distance] = 1;
            }
        }
    }
    return count;
};

const points = [[0,0],[1,0],[1,1],[1,-1],[2,0]];
// const points = [[0,0],[1,0],[1,1],[2,0]];
// const points = [[0,0],[1,0],[1,4],[2,0],[5,1],[5,2]];

console.log(numberOfBoomerangs(points));