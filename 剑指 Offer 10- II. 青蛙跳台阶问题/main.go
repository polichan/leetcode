package main

import "fmt"

func main() {
	fmt.Printf("%d", numWays(7))
}


func numWays(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
		dp[i] %= 1000000007
	}
	return dp[n]
}

func max(a, b int)int  {
	if a > b {
		return a
	}
	return b
}