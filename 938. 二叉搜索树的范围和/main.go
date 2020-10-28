package main

import "fmt"

func main() {
	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 18}}}
	fmt.Print(rangeSumBST(root, 7, 15))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	arr := make([]int, 0)
	inOrder(root, &arr)
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		if cur >= L && cur <= R{
			res += cur
		}
	}
	return res
}

func inOrder(node *TreeNode, arr *[]int){
	if node == nil {
		return
	}
	inOrder(node.Left, arr)
	*arr = append(*arr, node.Val)
	inOrder(node.Right,arr)
}