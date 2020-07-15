package main

func main() {
	nums := []int{2,7,11,15}
	twoSumUseMap(nums, 9)
}

/**
双指针
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