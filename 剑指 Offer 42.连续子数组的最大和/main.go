package main

func main() {
	test := []int{-2,1,-3,4,-1,2,1,-5,4}
	println(maxSubArray(test))
}


func maxSubArray(nums []int) int {
	numsLength := len(nums)
	if numsLength == 0 {
		return 0
	}
	if numsLength == 1 {
		return nums[0]
	}
	dp := make([]int, numsLength)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < numsLength; i++{
		dp[i] = max(nums[i] + dp[i - 1], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}