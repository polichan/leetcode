package main

import "strings"

func main() {
	
}

func strStr(haystack string, needle string) (ans int) {
	if len(needle) == 0 {
		return 0
	}
	return strings.Index(haystack, needle);
}

