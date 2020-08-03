package main

func main() {
	test := []int{-2,1,-3,4,-1,2,1,-5,4}
	println(maxSubArray(test))
}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	dp := make([]int, l)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i - 1] + nums[i], nums[i])
		result = max(dp[i], result)
	}

	return result
}

func max(x,y int)int  {
	if x > y {
		return x
	}
	return y
}