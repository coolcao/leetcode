/*
 * @lc app=leetcode id=101 lang=c
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (42.49%)
 * Total Accepted:    349.3K
 * Total Submissions: 822.1K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 *
 *
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 * But the following [1,2,2,null,3,null,3]  is not:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * Note:
 * Bonus points if you could solve it both recursively and iteratively.
 *
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

bool compare(struct TreeNode *node1, struct TreeNode *node2) {
  if (node1 == NULL && node2 == NULL) return true;

  if (node1 == NULL || node2 == NULL) return false;

  return (node1->val == node2->val && compare(node1->left, node2->right) && compare(node1->right, node2->left));

}

bool isSymmetric(struct TreeNode *root) {
  if (root == NULL) {
    return true;
  }
  return compare(root->left, root->right);
}

struct TreeNode *makeNode(int value) {
  struct TreeNode *root = (struct TreeNode *)malloc(sizeof *root);
  root->val = value;
  root->left = NULL;
  root->right = NULL;
  return root;
}

int main(int argc, char const *argv[]) {
  struct TreeNode *root = makeNode(1);
  struct TreeNode *left = makeNode(2);
  struct TreeNode *right = makeNode(2);
  struct TreeNode *leftleft = makeNode(3);
  struct TreeNode *leftright = makeNode(4);
  struct TreeNode *rightleft = makeNode(4);
  struct TreeNode *rightright = makeNode(3);

  left->left = leftleft;
  left->right = leftright;

  right->left = rightleft;
  right->right = rightright;

  root->left = left;
  root->right = right;
  bool b = isSymmetric(root);
  printf("%d\n", b);
  return 0;
}
