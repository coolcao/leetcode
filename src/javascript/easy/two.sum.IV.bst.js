/**
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True
Example 2:
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False
*/

const utils = require('../utils/utils.js');

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

class TreeNode {
    constructor(val) {
        this.val = val;
        this.left = this.right = null;
    }
}


/**
 * @param {TreeNode} root
 * @param {number} k
 * @return {boolean}
 */
const findTarget = function(root, k) {
    const hash = {};
    const check = function(node) {
        if (node) {
            if (hash[node.val]) {
                return true;
            }
            hash[k - node.val] = 1;
            return check(node.left) || check(node.right);
        } else {
            return false;
        }
    }
    return check(root);
};

const root = utils.createTree([5, 3, 6, 2, 4, 7]);
console.log(findTarget(root, 29));
