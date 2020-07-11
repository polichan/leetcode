package main

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
	count int
	leftCount int
}

func InsertNode(val int, root *TreeNode) int{
	if val == root.val {
		root.count++
		return root.leftCount
	}else if  val < root.val{
		root.leftCount++
		if root.left == nil {
			root.left = &TreeNode{val, nil,nil, 1, 0}
			return 0
		}
		return InsertNode(val, root.left)
	} else {
		if root.right == nil {
			root.right = &TreeNode{val, nil, nil, 1, 0}
			return root.lessOrEqual()
		}
		return root.lessOrEqual() + InsertNode(val, root.right)
	}
}

func (node *TreeNode)lessOrEqual()int  {
	return node.count + node.leftCount
}

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var ans = make([]int,len(nums))
	nums = reverse(nums)
	node := &TreeNode{nums[0], nil, nil, 1, 0}
	for index := 1; index < len(nums) ; index ++  {
		ans[index] = InsertNode(nums[index], node)
	}
	return reverse(ans)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
