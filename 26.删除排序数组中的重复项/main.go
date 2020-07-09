package main

func main() {
	nums := []int{1,1,2}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 1 {
		return len(nums)
	}
	for i:=len(nums)-1; i>0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}