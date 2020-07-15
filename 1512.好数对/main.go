package main

func main() {
	
}

/**
双指针遍历数组，符合条件结果+1 ，
每次循环将2指针+1 ， 判断2指针边界，超出边界，重置2指针为0 ， 将1指针进位。
循环以1指针到达边界停止。
 */
func numIdenticalPairs(nums []int) int {
	var index1, index2 int
	var r int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1] == nums[index2] {
				r ++
			}
		}
		index2 ++
		if index2 == len(nums) {
			index1 ++
			index2 = 0
		}
	}
	return r
}