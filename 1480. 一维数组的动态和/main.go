package main

import "fmt"

func main() {
	fmt.Print(runningSum([]int{
		1,2,3,4,
	}))
}


func runningSum(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	dp := make([]int, len(nums))
	// dp[i] 表示第 i 次保存的上一次计算值
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] =	nums[i] + dp[i-1]
	}
	return dp
}