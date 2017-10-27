'use strict';

/**
Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

Example:

Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
 */

/**
 * 递归实现
 * 对于BST，节点顺序是 左<根<右，因此解决该题需要按照从 右-根-左的顺序遍历BST即可。
 * 每遍历一个节点，将该节点的值累加到下一个遍历的节点上。
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
 * @return {TreeNode}
 */
const convertBST = function (root) {
  let sum = 0;
  const rightOrder = function(node) {
    if(node) {
      rightOrder(node.right);
      sum += node.val;
      node.val = sum;
      rightOrder(node.left);
    }
  }
  rightOrder(root);
  return root;
};


const utils = require('../utils/utils');
const root = utils.createTree([9, 5, 12, 2, 6, 10, 15]);
// const root = utils.createTree([12,10,15]);

const croot = convertBST(root);
console.log(JSON.stringify(utils.expandTree(croot), null, 2));




