package main

func main() {
	nums := []int{3,2,2,3}
	//删除第i个元素
	//i := 2
	//nums = append(nums[:i], nums[i+1:]...)
	removeElement(nums, 3)
	//fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	for i:= len(nums) -1 ; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}