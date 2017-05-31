'use strict';
/**
 * Find the sum of all left leaves in a given binary tree.

    Example:

        3
       / \
      9  20
        /  \
       15   7

    There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
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
const sumOfLeftLeaves = function(root) {
    let sum = 0;
    let q = [root];
    if(root===null || root===undefined){
        return sum;
    }
    while (q.length > 0) {
        let node = q.pop();
        if(node.left){
            if(!node.left.left && !node.left.right){
                sum += node.left.val;
            }
            q.push(node.left);
        }
        if(node.right){
            q.push(node.right);
        }
    }
    return sum;
};

const root = {
    val: 3,
    left: {
        val: 9
    },
    right: {
        val: 20,
        left: {
            val: 15
        },
        right: {
            val: 7
        }
    }
};

console.log(sumOfLeftLeaves(root));


