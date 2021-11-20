package main

import "fmt"

func main() {
	fmt.Println(permute([]int{
		1, 2, 3,
	}))
}

/**
permute 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
*/
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	var backtrack func(path []int, track []int)
	backtrack = func(path []int, track []int) {
		if len(track) == len(nums) {
			res = append(res, track)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 如果重复了，跳过
			if exist(track, nums[i]) {
				continue
			}
			backtrack(nums, append(track, nums[i]))
		}
	}
	backtrack(nums, track)
	return res
}

func exist(n []int, num int) bool {
	for _, item := range n {
		if item == num {
			return true
		}
	}
	return false
}
