package main

import "fmt"

func main() {
	test := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}}
	fmt.Print(increasingBST(test))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	slice := make([]int, 0)
	dfs(root, &slice)
	res := &TreeNode{Val: slice[0], Right: nil, Left: nil}
	currentNode := res
	for i := 1; i < len(slice); i++{
		next := &TreeNode{Val: slice[i], Right: nil, Left: nil}
		currentNode.Right = next
		currentNode = next
	}
	return res
}

func dfs(root *TreeNode, slice *[]int)  {
	if root == nil {
		return
	}
	dfs(root.Left, slice)
	*slice = append(*slice, root.Val)
	dfs(root.Right,slice)
}
