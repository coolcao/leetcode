'use strict';

/**
 * Given two binary trees, write a function to check if they are equal or not.

    Two binary trees are considered equal if they are structurally identical and the nodes have the same value.

    Subscribe to see which companies asked this question.
 */

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
const isSameTree = function(p, q) {
    if (p && q) {
        return (p.val === q.val) && isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    } else if (p === q) {
        return true;
    } else {
        return false;
    }
};

const p = {
    val: 1,
    left: {
        val: 2
    },
    right: {
        val: 3
    }
}
const q = {
    val: 1,
    left: {
        val: 2
    },
    right: {
        val: 4
    }
}

console.log(isSameTree(p, q));