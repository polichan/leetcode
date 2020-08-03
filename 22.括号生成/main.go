package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	res := new([]string)
	backtracking(n, n, "", res)
	return *res
}

func backtracking(left , right int, tmp string, res *[]string)  {
	if right == 0 {
		*res = append(*res, tmp)
		return
	}

	if left > 0 {
		backtracking(left - 1, right, tmp + "(", res)
	}

	if right > left{
		backtracking(left, right - 1, tmp + ")", res)
	}
}