'use strict';

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
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
 * @return {number[][]}
 */
const levelOrderBottom = function(root) {
  const result = [];
  const traver = function(node, level) {
    if (result[level]) {
      result[level].push(node.val);
    } else {
      result[level] = [];
      result[level].push(node.val);
    }
    if (node.left) {
      traver(node.left, level + 1)
    }
    if (node.right) {
      traver(node.right, level + 1)
    }
  }
  if (root) {
    traver(root, 0);
  }
  return result.reverse();
};
const utils = require('../utils/utils');
const array = [3, 9, 20, null, null, 15, 7];
const root = utils.createTree(array);
const result = levelOrderBottom(root);
console.log(result);
