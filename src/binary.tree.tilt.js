'use strict';
/**
 * Given a binary tree, return the tilt of the whole tree.

    The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

    The tilt of the whole tree is defined as the sum of all nodes' tilt.

    Example:
    Input: 
             1
           /   \
          2     3
    Output: 1
    Explanation: 
    Tilt of node 2 : 0
    Tilt of node 3 : 0
    Tilt of node 1 : |2-3| = 1
    Tilt of binary tree : 0 + 0 + 1 = 1
    Note:

    The sum of node values in any subtree won't exceed the range of 32-bit integer.
    All the tilt values won't exceed the range of 32-bit integer.

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
const findTilt = function(root) {
    const reduces = [];
    const q = [root];

    while (q.length > 0) {
        const node = q.pop();
        if(node.left && node.right){
            q.push(node.left);
            q.push(node.right);
            reduces.push(Math.abs(node.left.val - node.right.val));
        }else if(node.left){
            reduces.push(node.left.val);
            q.push(node.left);
        }else if(node.right){
            reduces.push(node.right.val);
            q.push(node.right);
        }
    }

    return reduces.reduce(function (pre,current) {
        pre += current;
        return pre;
    },0);

};

const root = {
    val: 1,
    left:{
        val: 2,
        right: {
            val: 4
        }
    },
    right:{
        val: 3
    }
};


console.log(findTilt(root));


