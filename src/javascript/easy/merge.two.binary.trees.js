'use strict';
/**
 * Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

    You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

    Example 1:
    Input: 
        Tree 1                     Tree 2                  
              1                         2                             
             / \                       / \                            
            3   2                     1   3                        
           /                           \   \                      
          5                             4   7                  
    Output: 
    Merged tree:
             3
            / \
           4   5
          / \   \ 
         5   4   7
    Note: The merging process must start from the root nodes of both trees.
 */

const utils = require('../utils/utils');
const TreeNode = require('../utils/TreeNode');

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} t1
 * @param {TreeNode} t2
 * @return {TreeNode}
 */
const mergeTrees = function(t1, t2) {
    let node = null;
    if (t1 && t2) {
        node = new TreeNode(t1.val + t2.val);
        node.left = mergeTrees(t1.left, t2.left);
        node.right = mergeTrees(t1.right, t2.right);
    } else if (t1) {
        node = new TreeNode(t1.val);
        node.left = mergeTrees(t1.left, null);
        node.right = mergeTrees(t1.right, null);
    } else if (t2) {
        node = new TreeNode(t2.val);
        node.left = mergeTrees(null, t2.left);
        node.right = mergeTrees(null, t2.right);
    }
    return node;
};

const t1 = utils.createTree([1, 3, 2, 5]);
const t2 = utils.createTree([2, 1, 3, null, 4, null, 7]);

console.log(mergeTrees(t1, t2));