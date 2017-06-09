'use strict';
/**
 * Given a binary tree, return the tilt of the whole tree.

    The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

    The tilt of the whole tree is defined as the sum of all nodes' tilt.

    Example:
    Input: 
             1
           /   \
          2     3
    Output: 1
    Explanation: 
    Tilt of node 2 : 0
    Tilt of node 3 : 0
    Tilt of node 1 : |2-3| = 1
    Tilt of binary tree : 0 + 0 + 1 = 1
    Note:

    The sum of node values in any subtree won't exceed the range of 32-bit integer.
    All the tilt values won't exceed the range of 32-bit integer.

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

const utils = require('../utils/utils');

const findTilt = function(root) {
    const sum = function (node) {
        let s = 0;
        if(node){
            s += (node.val + sum(node.left) + sum(node.right));
        }
        return s;
    }
    const tilt = function (node) {
        if(!node.left && !node.right){
            return 0;
        }else if(!node.left){
            return Math.abs(sum(node.right));
        }else if(!node.right){
            return Math.abs(sum(node.left));
        }else{
            return Math.abs(sum(node.left) - sum(node.right));
        }
    }
    let _s = 0;
    if(!root){
        return 0;
    }
    const q = [root];
    while (q.length > 0) {
        const node = q.pop();
        _s += tilt(node);
        if (node.left) {
            q.push(node.left);
        }
        if (node.right) {
            q.push(node.right);
        }
    }
    return _s;
};

const root = utils.createTree([-8,3,0,-8,null,null,null,null,-1,null,8]);
console.log(findTilt(root));


