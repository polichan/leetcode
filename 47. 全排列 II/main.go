package main

import (
	"fmt"
)

func main() {
	fmt.Println(permuteUnique([]int{
		1, 1, 2,
	}))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	var backtrace func(nums []int, track []int, used []bool)
	backtrace = func(nums []int, track []int, used []bool) {
		if len(nums) == len(track) {
			res = append(res, track)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i-1 >= 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			backtrace(nums, append(track, nums[i]), used)
			used[i] = false
		}
	}
	backtrace(nums, track, used)
	return res
}
