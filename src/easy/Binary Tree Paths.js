'use strict'

/**
Given a binary tree, return all root-to-leaf paths.

For example, given the following binary tree:

   1
 /   \
2     3
 \
  5
All root-to-leaf paths are:

["1->2->5", "1->3"]
Credits:
Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.


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
 * @return {string[]}
 */
const binaryTreePaths = function(root) {
    const path = function(node,pre,level,result) {
        if (node) {
            if (level) pre += '->' + node.val;
            else pre += node.val;
            if (!node.left && !node.right) {
                result.push(pre);
            }
            if (node.left) {
                path(node.left, pre, level + 1, result)
            }
            if (node.right) {
                path(node.right, pre, level + 1, result);
            }
        }
    }
    const result = [];
    path(root, '', 0,result);
    return result;
};

const utils = require('../utils/utils');

const root = utils.createTree([1,2,3,4,5,null,null,7,null,null,8]);

console.log(binaryTreePaths(root));