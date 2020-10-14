package main

func main() {
	ans := split(7, 3)
	print(ans)
}

/**
数的划分
 */
func split(n, k int)int  {
	dp := make([][]int, 200)
	dp = fill(dp)
	print(dp)
	return 1
}

func fill(arr [][]int) [][]int{
	for i := 0; i <= 200; i++ {
		arr[0][i] = 0
	}
	return arr
}
