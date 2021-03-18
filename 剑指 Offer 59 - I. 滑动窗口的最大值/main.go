package main

import (
	"fmt"
)

func main() {
	arr := []int{
		7,2,4,
	}
	fmt.Print(maxSlidingWindow(arr, 2))
}
/**

给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
示例：
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


 */
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums){
		return nil
	}
	var maxNums []int
	//max用于记录窗口中的最大值的索引
	max := -1
	for i:=0;i<=len(nums)-k;i++{
		l := i
		r := i+k-1
		//如果出去的值是最大值(的索引)，那么新的窗口需要重新查找最大值并设置
		if max == -1 || max == l-1 {
			max = getMax(nums,l,r)
		}else{
			//如果出去的值不是最大值(的索引)，则判断新进来的值是否大于当前的最大值
			//不大于则最大值(的索引)不变
			//大于则将进来的值(的索引)作为最大值(的索引)
			if nums[r] > nums[max] {
				max = r
			}
		}
		maxNums = append(maxNums,nums[max])
	}
	return maxNums
}

//双指针查找最大值的索引
func getMax(nums []int,left,right int) int {
	for left<right{
		if nums[left]>nums[right]{
			right--
		}else{
			left++
		}
	}
	return left
}

