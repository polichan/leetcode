package main

import "math"

func main() {
	reverse(123)
}

func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans > math.MaxInt32 /10 || ans < math.MinInt32 / 10 {
			return 0
		}
		ans = ans * 10
		ans += x % 10
		x /= 10
	}
	return ans
}