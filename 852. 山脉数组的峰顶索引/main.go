package main

import "fmt"

func main() {
	arr := []int{
		0,1,2,1,
	}
	fmt.Printf("%d \n", peakIndexInMountainArray(arr))
}


func peakIndexInMountainArray(arr []int) int {
	// arr 0 1  2 1
	left, right := 0, len(arr) - 1
	// left 0 right = 3
	for left <= right {
		mid := left + (right - left) / 2 // 1
		if arr[mid] < arr[mid + 1] { // arr[mid] = 1 arr[mid + 1] 0
			left = mid + 1 // left = 2
		}else if arr[mid] < arr[mid - 1]{
			right = mid - 1
		}else{
			return mid
		}
	}
	return -1
}