package main

func main() {

}

/**
给你二叉树的根节点root 和一个表示目标和的整数targetSum ，
判断该树中是否存在 根节点到叶子节点 的路径，
这条路径上所有节点值相加等于目标和targetSum 。

叶子节点 是指没有子节点的节点。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return false
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	preOrder(root.Right)
}
