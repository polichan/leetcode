package main

func main() {
	
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minInput := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		minInput = min(price, minInput)
		maxProfit = max(maxProfit, price - minInput)
	}
	return maxProfit
}

func max(x,y int) int{
	if   x> y{
		return x
	}
	return y
}

func min(x,y int) int{
	if   x < y{
		return x
	}
	return y
}
