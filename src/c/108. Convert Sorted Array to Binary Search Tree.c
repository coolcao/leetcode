/*
Given an array where elements are sorted in ascending order, convert it to a
height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in
which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following
height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
*/
#include <stdio.h>
#include <stdlib.h>
struct TreeNode {
  int val;
  struct TreeNode* left;
  struct TreeNode* right;
};

struct TreeNode* createNode(int num) {
    struct TreeNode* node = malloc(sizeof *node);
    node->val = num;
    node->left = NULL;
    node->right = NULL;
    return node;
}

// start包含，end不包含
struct TreeNode* createTree(int* nums, int start, int end) {
    if (start >= end) return NULL;
    int idx = (start + end) / 2;
    struct TreeNode* root = createNode(nums[idx]);
    root->left = createTree(nums, start, idx);
    root->right = createTree(nums, idx+1, end);
    return root;
}

struct TreeNode* sortedArrayToBST(int* nums, int numsSize) {
    if(numsSize == 0) return NULL;
    return createTree(nums, 0, numsSize);
}

int main(int argc, char const *argv[])
{
    int nums[10] = {1,2,3,4,5,6,7,8,9,10};
    sortedArrayToBST(nums, 10);
    return 0;
}

