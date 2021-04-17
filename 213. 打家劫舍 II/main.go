package main

func main() {
	nums := []int{
		1,2,3,
	}
	rob(nums)
}
/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的
房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。
 */
func rob(nums []int) int {
	n := len(nums)
	if n == 1{
		return nums[0]
	}
	// 两间房屋只能偷一个
	if n == 2{
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func _rob(nums[]int) int  {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:]{
		first, second = second, max(first + v, second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}