package main

import "fmt"

func main() {
	var arr = []int{
		2,3,4,
	}
	fmt.Printf("%d", twoSum(arr, 6))
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i + 1, len(numbers) - 1
		for left <= right {
			mid := left + (right - left) / 2
			if numbers[mid] == target - numbers[i] {
				return []int{
					i + 1, mid + 1,
				}
			}else if numbers[mid] > target - numbers[i]{
				right = mid - 1
			}else{
				left = mid + 1
			}
		}
	}
	return []int{-1, -1}
}