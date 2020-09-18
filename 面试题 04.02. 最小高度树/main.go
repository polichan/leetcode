package main

func main() {
	
}


/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {return &TreeNode{Val: nums[0]}}
	middle := len(nums) / 2
	root := &TreeNode{Val: nums[middle]}
	root.Left = sortedArrayToBST(nums[:middle])
	root.Right = sortedArrayToBST(nums[middle + 1:])
	return root
}