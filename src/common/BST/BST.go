package main

import "fmt"

type BinarySearchTree struct {
	root *TreeNode
}

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func (b *BinarySearchTree) Insert(val int){
	if b.root == nil {
		b.root = &TreeNode{val, nil, nil}
	}else {
		b.root = InsertNode(val, b.root)
	}
}

func InsertNode(val int, root *TreeNode) *TreeNode{
	if val < root.val {
		if root.left == nil {
			root.left = &TreeNode{val, nil,nil}
		}else {
			root.left = InsertNode(val, root.left)
		}
	}else {
		if root.right == nil {
			root.right = &TreeNode{val, nil, nil}
		}else {
			root.right = InsertNode(val, root.right)
		}
	}
	return root
}

func (b *BinarySearchTree) InOrder() {
	InOrder(b.root)
}

func InOrder(root *TreeNode)  {
	if root != nil {
		InOrder(root.left)
		fmt.Println(root.val)
		InOrder(root.right)
	}
}


func main() {
	var test = []int{5,2,6,1}
	tree := BinarySearchTree{nil}
	for _, num := range test{
		tree.Insert(num)
	}
	tree.InOrder()
}
