/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root
node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.


*/

#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

int visit(struct TreeNode *root, int depth) {
    if (root == NULL) return depth;
    int left = visit(root->left, depth + 1);
    int right = visit(root->right, depth + 1);
    return left>right?left:right;
}

int maxDepth(struct TreeNode *root) {
    if (root == NULL) return 0;
    return visit(root, 0);
}
