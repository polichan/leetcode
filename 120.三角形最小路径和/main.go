package main

func main() {
	test := [][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	println(minimumTotal(test))
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	for i := n ; i>= 0 ; i -- {
		for j := len(triangle[i]) - 1; j >= 0; j -- {
			triangle[i][j] = triangle[i][j] + min(triangle[i + 1][j] , triangle[i + 1][j + 1])
		}
	}
	return triangle[0][0]
}
func min(x, y int)int{
	if x < y {
		return x
	}
	return y
}