'use strict';

/**
 * Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

    Example:

    Input:

       1
        \
         3
        /
       2

    Output:
    1

    Explanation:
    The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
    Note: There are at least two nodes in this BST.


 */

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
const getMinimumDifference = function(root) {
    const getOrderValues = function (root) {
        const values = [];
        const readInOrder = function (root,values) {
            if(root){
                readInOrder(root.left,values);
                values.push(root.val);
                readInOrder(root.right,values);
            }
        }
        readInOrder(root,values);
        return values;
    }
    const orderedValues = getOrderValues(root);

    let min = orderedValues[1] - orderedValues[0];

    for(let i=2;i<orderedValues.length;i++){
        let diff = orderedValues[i] - orderedValues[i-1];
        if(min > diff){
            min = diff;
        }
    }

    return min;

};

const root = {
    val: 8,
    left: {
        val: 4,
        left: { val: 1 },
        right: { val: 6 }
    },
    right: {
        val: 12,
        left: { val: 10 },
        right: { val: 15 }
    }
};

console.log(getMinimumDifference(root));