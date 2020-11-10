package main

import "fmt"

func main() {
	arr := []int{
		1,3,5,
	}
	fmt.Printf("%d \n", minArray(arr))
}


func minArray(numbers []int) int {
	left, right := 0, len(numbers) - 1
	for left < right {
		mid := left + right /2
		if numbers[mid] > numbers[right]{
			left = mid + 1
		}else if numbers[mid] < numbers[right] {
			right = mid
		}else{
			right --
		}
	}
	return numbers[left]
}