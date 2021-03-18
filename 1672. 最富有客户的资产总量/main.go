package main

import "fmt"

func main()  {
	accounts  := [][]int{
		[]int{1,2,3},
		[]int{3,2,1},
	}
	fmt.Printf("%d", maximumWealth(accounts))
}
/**
给你一个 m x n 的整数网格 accounts ，其中 accounts[i][j] 是第 i​​​​​​​​​​​​ 位客户在第 j 家银行托管的资产数量。返回最富有客户所拥有的 资产总量 。

客户的 资产总量 就是他们在各家银行托管的资产数量之和。最富有客户就是 资产总量 最大的客户。
 */
func maximumWealth(accounts [][]int) int {
	max := 0
	for i := 0; i <= len(accounts) - 1; i++ {
		tempSum := 0
		for j := 0; j <= len(accounts[i]) - 1; j ++ {
			tempSum += accounts[i][j]
		}
		if tempSum > max {
			max = tempSum
		}
	}
	return max
}
