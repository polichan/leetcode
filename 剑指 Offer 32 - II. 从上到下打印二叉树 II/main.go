package main

func main() {
	t := &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil} , Right: &TreeNode{Val: 20, Right: &TreeNode{Val: 7, Left: nil, Right: nil} ,Left: &TreeNode{Val: 15, Left: nil , Right: nil}}}
	levelOrder(t)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var ret [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {return nil}
	ret = make([][]int,1)
	ret[0] = []int{root.Val}
	if root.Left != nil {subFunc(root.Left, 1)}
	if root.Right != nil {subFunc(root.Right, 1)}
	return ret
}

func subFunc (root *TreeNode, depth int) {
	if depth >= len(ret) {
		ret = append(ret, []int{root.Val})
	} else {
		ret[depth] = append(ret[depth], root.Val)
	}
	if root.Left != nil {
		subFunc(root.Left, depth+1)
	}
	if root.Right != nil {
		subFunc(root.Right, depth+1)
	}
}