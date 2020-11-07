package main

func main() {
	sortedArrayToBST([]int{
		-10,-3,0,5,9,
	})
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]
	node := &TreeNode{nums[mid], sortedArrayToBST(left), sortedArrayToBST(right)}
	return node
}