package main

import "fmt"

func main() {
	node := &Node{Val: 1, Children: []*Node{
		&Node{Val: 3, Children: []*Node{
			&Node{Val: 5, Children: nil},
			&Node{Val: 6, Children: nil},
		}},
		&Node{Val: 2, Children: nil},
		&Node{Val: 4, Children: nil},
	}}
	fmt.Print(postorder(node))
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node)[]int  {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int)  {
	if root != nil {
		for _, n := range root.Children{
			dfs(n, res)
		}
		*res = append(*res, root.Val)
	}
}
