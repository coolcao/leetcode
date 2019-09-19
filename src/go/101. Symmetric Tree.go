package main

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Note:
Bonus points if you could solve it both recursively and iteratively.
*/

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compareNode(node1 *TreeNode, node2 *TreeNode) {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && compareNode(node1.Left, node2.Right) && compareNode(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareNode(root.Left, root.Right)
}

func main() {

}
