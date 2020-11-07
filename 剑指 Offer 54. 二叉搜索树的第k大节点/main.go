package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil,
		Right: nil}}, Right: &TreeNode{Val: 4}}
	fmt.Printf("%d", kthLargest(tree, 1))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var skip int
var res int
func kthLargest(root *TreeNode, k int) int {
	skip = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		skip--
		if skip == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
}