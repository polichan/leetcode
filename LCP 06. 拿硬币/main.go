package main

func main() {
	
}


func minCount(coins []int) int {
	if len(coins) == 0 {
		return 0
	}
	var count int
	for _, item := range coins{
		if item % 2 == 0{
			count += item / 2
		}else{
			count ++
			count += (item - 1) / 2
		}
	}
	return count
}