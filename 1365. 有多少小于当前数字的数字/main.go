package main

import (
	"sort"
)

func main() {
	smallerNumbersThanCurrent([]int{
		8,1,2,2,3,
	})
}

/**
给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。

换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。

以数组形式返回答案。
 */

type pair struct{ v, pos int }

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	data := make([]pair, n)
	for i, v := range nums {
		data[i] = pair{v, i}
	}
	sort.Slice(data, func(i, j int) bool { return data[i].v < data[j].v })
	ans := make([]int, n
	prev := -1
	for i, d := range data {
		if prev == -1 || d.v != data[i-1].v {
			prev = i
		}
		ans[d.pos] = prev
	}
	return ans
}