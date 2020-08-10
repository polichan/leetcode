package main

import "fmt"

func main() {
	nums := []int{2,7,11,15}
	fmt.Println(twoSumUseMap(nums, 9))
}

/**
双指针
执行用时：
124 ms
, 在所有 Go 提交中击败了
5.84%
的用户
内存消耗：
2.9 MB
, 在所有 Go 提交中击败了
100.00%
的用户
 */
func twoSum(nums []int, target int) []int {
	var index1, index2 int
	var ans []int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1]+nums[index2] == target {
				ans = append(ans, index1)
				ans = append(ans, index2)
				break
			}
		}
		index2 ++
		if index2 == len(nums) {
			index1 ++
			index2 = 0
		}
	}
	return ans
}

/**
hashMap
执行用时：
4 ms
, 在所有 Go 提交中击败了
96.88%
的用户
内存消耗：
3.8 MB
, 在所有 Go 提交中击败了
51.79%
的用户
 */
func twoSumUseMap(nums []int, target int) []int {
	m := map[int]int{}
	for i,v := range nums{
		if k, ok := m[target-v]; ok {
			return []int{i,k}
		}
		m[v] = i
	}
	return nil
}