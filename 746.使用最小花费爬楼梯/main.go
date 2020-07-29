package main

func main() {
	test := []int{10,15,20}
	minCostClimbingStairs(test)
}


func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return cost[0]
	}
	cost = append(cost,0)
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = min(cost[0] + cost[1], cost[1])
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i - 2], dp[i-1])
	}
	return dp[len(cost) - 1]
}

func min(x,y int) int{
	if   x < y{
		return x
	}
	return y
}