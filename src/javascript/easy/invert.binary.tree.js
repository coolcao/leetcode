'use strict';

const utils = require('./utils');

/**
 * Invert a binary tree.

         4
       /   \
      2     7
     / \   / \
    1   3 6   9

    to
         4
       /   \
      7     2
     / \   / \
    9   6 3   1

    Trivia:
    This problem was inspired by this original tweet by Max Howell:
    Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so fuck off.
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
var invertTree = function(root) {

    if (!root) {
        return null;
    }

    const nodes = [root];

    while (nodes.length > 0) {
        let node = nodes.shift();
        let tmp = node.left;
        node.left = node.right;
        node.right = tmp;

        if (node.left) {
            nodes.push(node.left);
        }
        if (node.right) {
            nodes.push(node.right);
        }
    }


    return root;
};

const root = utils.createTree([4, 1, null, 2, null, 3]);

console.log(JSON.stringify(root));
console.log(JSON.stringify(invertTree(root)));