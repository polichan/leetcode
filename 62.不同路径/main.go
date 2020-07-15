package main


func main() {
	uniquePaths(2,3)
}

func uniquePaths(m int, n int) int {
	dp := make([]int,n)
	for index, _ := range dp{
		dp[index] = 1
	}

	for i := 1; i< m;i++ {
		for j := 1; j< n; j++ {
			dp[j] = dp[j -1] + dp[j]
		}
	}
	return dp[n - 1]
}