package main

import "fmt"

func main(){
	c := []int{
		2,3,5,1,3,
	}
	fmt.Print(kidsWithCandies(c, 3))
}

/**

给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目
 */
func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxC := 0
	for i := 0; i <= len(candies) - 1; i ++{
		maxC = max(maxC, candies[i])
	}
	for i := 0; i <= len(candies) - 1; i++{
		res[i] = candies[i] + extraCandies >= maxC
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}