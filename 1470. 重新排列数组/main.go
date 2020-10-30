package main

import "fmt"

func main() {
	fmt.Print(shuffle([]int{1,4,3,2,2,3,4,1}, 4))
}

func shuffle(nums []int, n int) []int {
	index := 0
	res := make([]int, len(nums))
	for i := 0; i < n; i ++ {
		res[i * 2] = nums[index]
		res[i * 2 + 1] = nums[index + n]
		index++
	}
	return res
}