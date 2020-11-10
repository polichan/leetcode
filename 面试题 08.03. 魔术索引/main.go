package main

import (
	"fmt"
)

func main() {
 	arr := []int{
 		0,0,2,
	}
	fmt.Printf("%d \n", findMagicIndex(arr))
}


func findMagicIndex(nums []int) int {
	for i := 0; i< len(nums); i++{
		if nums[i] == i {
			return i
		}
	}
	return -1
}