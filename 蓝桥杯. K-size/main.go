package main

import "fmt"

func main() {
	fmt.Print(kSize([]string{"a", "b", "b", "b"}))
}

func kSize(arr []string)int{
	// Sliding Window
	windowSize := 2 // 窗口大小
	res := 0
	for i := 0; i < len(arr); i++{
		// 当前窗口下的数组
		if arr[i] == arr[i+ windowSize - 1] {
			res++
		}
	}
	return res
}

