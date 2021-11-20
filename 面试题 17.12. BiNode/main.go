package main

import "fmt"

// [4,2,5,1,3,null,6,0]
func main() {
	r := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	x := convertBiNode(r)
	fmt.Println(x)
}

// TreeNode 树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	node *TreeNode
	prt  *TreeNode
}

func convertBiNode(root *TreeNode) *TreeNode {
	var tree = &Tree{}
	tree.node = &TreeNode{}
	// tree.prt =tree.node;
	tree.TreeTo(root)
	return tree.node.Right
}

func (t *Tree) TreeTo(next *TreeNode) {
	if next == nil {
		return
	}
	t.TreeTo(next.Left)
	if t.prt == nil {
		t.prt = next
		t.node.Right = next
	} else {
		t.prt.Right = next
		t.prt = next
	}
	next.Left = nil
	t.TreeTo(next.Right)
}
