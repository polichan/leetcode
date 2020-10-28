package main

import "fmt"

func main() {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 18}}}
	fmt.Print(maxDepth(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}


func max(x,y int)int{
	if x > y {
		return x
	}
	return y
}