package main

import "fmt"

func main() {
	fmt.Print(reverseLeftWords("abcdefg", 2))
}

func reverseLeftWords(s string, n int) string {
	a := []rune(s)
	if n == len(a) {
		return s
	}
	var cutStr string
	var cutRightStr string
	for i := 0; i < len(a); i ++ {
		if i < n {
			cutStr = cutStr + string(a[i])
		}else{
			cutRightStr = cutRightStr + string(a[i])
		}
	}
	return cutRightStr + cutStr
}