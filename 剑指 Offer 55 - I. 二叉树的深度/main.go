package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	
}
/**
后续遍历 DFS
 */
func maxDepth(root *TreeNode) int {
	if root == nil {return 0}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

/**
层序遍历 BFS
 */
func maxDepthBFS(root *TreeNode) int  {
	if root == nil {
		return 0
	}
	res := 0
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		tmp := make([]*TreeNode, 0)
		for _, node := range q{
			if node.Left != nil{
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		q = tmp
		res += 1
	}
	return res
}
func max(a ,b int) int {
	if a > b {
		return a
	}
	return b
}