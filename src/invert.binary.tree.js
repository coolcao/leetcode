'use strict';

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

    const nodes = [];
    nodes.push(root);
    if (root.left) {
        nodes.push(root.left);
    }
    if (root.right){
        nodes.push(root.right);
    }

    while (nodes.length > 0) {
        let node = nodes.pop();
        let tmp = node.left;
        node.left = node.right;
        node.right = tmp;
    }


    return root;
};

const root = {
    val: 1,
    left: {
        val: 2,
        left: {
            val: 4,
            left: null,
            right: null
        },
        right: null
    },
    right: {
        val: 3,
        left: null,
        right: null
    }
};

console.log(JSON.stringify(invertTree(root),null,2));