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
	if len(slice) == 1{
		return &TreeNode{Val: slice[0], Right: nil, Left: nil}
	}
	res := &TreeNode{Val: slice[0], Right: &TreeNode{}, Left: nil}
	currentNode := res.Right
	for i := 1; i < len(slice); i++{
		currentNode.Val = slice[i]
		if i != len(slice) - 1 {
			currentNode.Right = &TreeNode{}
		}
		currentNode = currentNode.Right
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
