'use strict';

/**
Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2
Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
*/

const utils = require('../utils/utils.js');
const TreeNode = require('../utils/TreeNode');
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} L
 * @param {number} R
 * @return {TreeNode}
 */
const trimBST = function(root, L, R) {
  const trim = function(node) {
    if (node) {
      if (node.val >= L && node.val <= R) {
        node.left = trim(node.left);
        node.right = trim(node.right);
      } else if (node.val >= L) {
        node = trim(node.left);
      } else if (node.val <= R) {
        node = trim(node.right);
      }
    }
    return node;
  }
  return trim(root)
};

let root = utils.createTree([3,0,4,null,2,null,null,1]);
// console.log(JSON.stringify(root,null,2));
console.log(JSON.stringify(trimBST(root,1,3),null,2));
