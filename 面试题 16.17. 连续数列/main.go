package main

func main() {
	
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
	res := dp[0]
	for i := 1; i< l; i++ {
		dp[i] = max(dp[i-1] + nums[i], nums[i])
		if  dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(x,y int)int  {
	if x> y {
		return x
	}
	return y
}