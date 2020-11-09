package main

func main() {
	
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	res := make([]int , 0)
	order(root, &res)
	mp := make(map[int]int, 0)
	for _, m := range res{
		mp[m]++
	}
	return len(mp) == 1
}

func order(root *TreeNode, res *[]int)  {
	if root == nil {
		return
	}
	order(root.Left, res)
	*res = append(*res, root.Val)
	order(root.Right, res)
}