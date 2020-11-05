package main

import "fmt"

func main() {
	test := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 7, Right: nil, Left: nil}}
	fmt.Print(searchBST(test, 2))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	return helper(root,val)
}

func helper(root *TreeNode, val int) *TreeNode{
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}else{
		return searchBST(root.Right, val)
	}
}

func helperBySecondSolution(root *TreeNode, val int)*TreeNode  {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		}else{
			root = root.Right
		}
	}
	return root
}