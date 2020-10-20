package main

import "fmt"

func main() {
	res := plusOne([]int{
		1,2,3,
	})

	fmt.Print(res)
}

func plusOne(digits []int) []int {
	res := []int{1}
	temp := 0
	l := len(digits)
	count := 0
	for i := l - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1 + temp
		if digits[i] == 10 {
			// 记录进位次数
			count ++
			// 向前进位 1
			temp = digits[i] % 10
			// 10 要变成 0
			digits[i] = 0
		}else{
			break
		}
	}
	if count == l {
		res = append(res, digits...)
	}else{
		res = digits
	}
	return res
}