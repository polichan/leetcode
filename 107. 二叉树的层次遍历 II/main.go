package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var ret [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	ret = levelOrder(root)
	res := make([][]int , len(ret))
	resIndex := 0
	for i := len(ret) - 1; i >= 0; i--{
		res[resIndex] = ret[i]
		resIndex++
	}
	return res
}

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