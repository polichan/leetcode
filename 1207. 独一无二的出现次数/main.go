package main

import "fmt"

func main() {
	arr := []int{
		1,2,2,1,1,3,
	}
	fmt.Print(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	// 每个数字出现的次数
	m := make(map[int]int, len(arr))
	for i := 0; i< len(arr); i++{
		m[arr[i]] ++
	}
	// 假设 m 中 {1: 3, 2 : 3, 3: 1} 则会导致两次 3 的次数出现，那么 times 中就会变成两个 key，然后通过比较 map 的长度，就能知道是不是出现了同一个次数
	times := map[int]struct{}{}
	for _, c := range m {
		times[c] = struct{}{}
	}
	return len(m) == len(times)
}
